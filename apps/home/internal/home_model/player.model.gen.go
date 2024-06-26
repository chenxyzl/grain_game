// Code generated by https://github.com/chenxyzl/gsgen; DO NOT EDIT.
// gen_tools version: 1.1.5
// generate time: 2024-06-26 10:48:22
package home_model

import (
	"encoding/json"
	"fmt"
)

func (s *Player) GetUid() uint64 {
	return s.uid
}
func (s *Player) SetUid(v uint64) {
	s.uid = v
	s.UpdateDirty(1 << 0)
}
func (s *Player) GetBase() *Base {
	return s.base
}
func (s *Player) SetBase(v *Base) {
	if v != nil {
		v.SetParent(1, s.UpdateDirty)
	}
	s.base = v
	s.UpdateDirty(1 << 1)
}
func (s *Player) GetGame() *Game {
	return s.game
}
func (s *Player) SetGame(v *Game) {
	if v != nil {
		v.SetParent(2, s.UpdateDirty)
	}
	s.game = v
	s.UpdateDirty(1 << 2)
}
func (s *Player) String() string {
	doc := struct {
		Uid  uint64 `bson:"_id"`
		Base *Base  `bson:"base"`
		Game *Game  `bson:"game"`
	}{s.uid, s.base, s.game}
	return fmt.Sprintf("%v", &doc)
}
func (s *Player) MarshalJSON() ([]byte, error) {
	doc := struct {
		Uid  uint64 `bson:"_id"`
		Base *Base  `bson:"base"`
		Game *Game  `bson:"game"`
	}{s.uid, s.base, s.game}
	return json.Marshal(doc)
}
func (s *Player) UnmarshalJSON(data []byte) error {
	doc := struct {
		Uid  uint64 `bson:"_id"`
		Base *Base  `bson:"base"`
		Game *Game  `bson:"game"`
	}{}
	if err := json.Unmarshal(data, &doc); err != nil {
		return err
	}
	s.SetUid(doc.Uid)
	s.SetBase(doc.Base)
	s.SetGame(doc.Game)
	return nil
}
func (s *Player) Clone() (*Player, error) {
	data, err := json.Marshal(s)
	if err != nil {
		return nil, err
	}
	ret := Player{}
	err = json.Unmarshal(data, &ret)
	if err != nil {
		return nil, err
	}
	return &ret, nil
}
