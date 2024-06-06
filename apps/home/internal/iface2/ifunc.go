package iface2

import "grain_game/proto/gen/ret"

// OnInit 开始前-激活actor时 只能初始化db和代码
type OnInit func() ret.Code

// OnStarted 开始后~首次tick执行前
type OnStarted func() ret.Code

// OnPreStop 结束前
type OnPreStop func() ret.Code

// OnTickFun 每秒
type OnTickFun func(nowUnix int64)

// OnCrossDay 跨天
// @natural: true:自然跨天 false:登录的非自然跨天
// @daysUnix: 跨了哪些天
type OnCrossDay = func(natural bool, daysUnix ...int64)

// OnCrossWeek
// @natural: true:自然跨天 false:登录的非自然跨天
// @daysUnix: 跨了哪些天
type OnCrossWeek = func(natural bool, daysUnix ...int64)

// OnlineFun 在线
type OnlineFun = func(nowUnix int64)

// OfflineFun 离线
type OfflineFun = func(nowUnix int64)
