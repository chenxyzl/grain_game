package internal

import (
	"errors"
	"fmt"
	"github.com/chenxyzl/grain/actor"
	"github.com/chenxyzl/grain/al/safemap"
	"github.com/gobwas/ws"
	"github.com/gobwas/ws/wsutil"
	"github.com/golang/protobuf/proto"
	"grain_game/apps/gate/internal/constant1"
	"grain_game/apps/shared/config"
	"grain_game/apps/shared/utils"
	"grain_game/proto/gen/ret"
	"io"
	"net"
	"net/http"
	"net/url"
)

var _ actor.IActor = (*WebsocketServer)(nil)

type WebsocketServer struct {
	actor.BaseActor
	path     string
	host     string
	port     string
	sessions *safemap.SafeMap[string, *actor.ActorRef]
}

func NewWebsocketServer(path, host, port string) *WebsocketServer {
	return &WebsocketServer{path: path, host: host, port: port, sessions: safemap.NewM[string, *actor.ActorRef]()}
}

func (wss *WebsocketServer) Started() {
	wss.Logger().Info("websocket server start...")
	http.HandleFunc(wss.path, wss.helpersHighLevelHandler)
	addr := net.JoinHostPort(wss.host, wss.port)
	ln, err := net.Listen("tcp", addr)
	if err != nil {
		panic(fmt.Errorf("websocket server listen %q error: %v", addr, err))
	}

	s := new(http.Server)
	go func() {
		err = s.Serve(ln)
		if err != nil {
			if !errors.Is(err, http.ErrServerClosed) {
				panic(errors.Join(err, errors.New("websocket server closed unexpectedly")))
			} else {
				wss.Logger().Info("websocket server closed")
			}
		}
	}()
	err = wss.System().GetProvider().SetNodeExtData(config.Get().GetGate().GetWssubkey(), ln.Addr().String())
	if err != nil {
		panic(errors.Join(err, errors.New("websocket server set node ext data fail")))
	}
	wss.Logger().Info("websocket server start success", "websocket.addr", ln.Addr(), "websocket.addr", addr, "websocket.wsPath", wss.path)
}

func (wss *WebsocketServer) PreStop() {
	//TODO implement me
	panic("implement me")
}

func (wss *WebsocketServer) Receive(ctx actor.Context) {
	//TODO implement me
	panic("implement me")
}

func (wss *WebsocketServer) helpersHighLevelHandler(w http.ResponseWriter, r *http.Request) {
	//recover
	defer utils.Recover(func(e any, trace string) { wss.Logger().Error("sess closed with error", "err", e, "trace", trace) })
	//update to websocket connection
	conn, _, _, err := ws.UpgradeHTTP(r, w)
	if err != nil {
		wss.Logger().Error("session upgrade connection failed", "err", err)
		return
	}
	defer conn.Close()
	//parse uri
	uri := r.RequestURI
	rq, err := url.Parse(uri)
	if err != nil {
		wss.Logger().Error("session parse uri failed", "uri", uri, "err", err)
		return
	}
	//parse rq
	m, err := url.ParseQuery(rq.RawQuery)
	if err != nil {
		wss.Logger().Error("session parse query error", "uri", uri, "rq", rq.RawQuery, "err", err)
		return
	}
	//get token
	token := m.Get(constant1.ParamToken)
	if token == "" {
		wss.Logger().Error("session get token is nil", "uri", uri, "rq", rq.RawQuery, "m", m)
		return
	}
	//todo parse token?
	//
	wss.createSession(conn)
}

func (wss *WebsocketServer) createSession(conn net.Conn) {
	sess := wss.System().Spawn(func() actor.IActor { return newSession(wss.Self(), conn) }, actor.WithOptsKindName(common.common.SessionKind))
	wss.sessions.Set(sess.GetId(), sess)
	defer func() { wss.sessions.Delete(sess.GetId()); wss.System().Poison(sess) }()
	wss.Logger().Info("session created", "id", sess.GetId())
	//read msg
	for {
		bts, _, err := wsutil.ReadClientData(conn)
		if err == io.EOF {
			wss.Logger().Info("sess normal close1", "id", sess.GetId(), "err", err)
			return
		}
		var netError net.Error
		if errors.As(err, &netError) && netError.Timeout() {
			wss.Logger().Info("sess normal close2", "id", sess.GetId(), "err", err)
			return
		}
		var closeError wsutil.ClosedError
		if errors.As(err, &closeError) &&
			(closeError.Code == ws.StatusNoStatusRcvd ||
				closeError.Code == ws.StatusNormalClosure ||
				closeError.Code == ws.StatusGoingAway) {
			wss.Logger().Info("sess normal close3", "id", sess.GetId(), "err", err)
			return
		}
		if err != nil {
			wss.Logger().Error("sess read message error", "id", sess.GetId(), "err", err)
			return
		}
		//
		reqPack := &ret.ReqPack{}
		err = proto.Unmarshal(bts, reqPack)
		if err != nil {
			wss.Logger().Error("sess unmarshal err", "id", sess.GetId(), "err", err)
		}
		actor.NoEntrySend(wss.System(), sess, reqPack)
	}
}
