package model

import (
	"encoding/json"
	"fmt"
	"net/url"
	"strconv"
)

type User struct {
	Id     uint64  `json:"id, omitempty"`
	Events []Event `json:"events, omitempty"`
}

func (u *User) ToJSON() ([]byte, error) {
	return json.Marshal(u)
}

func (u *User) Parsing(data url.Values) error {
	id := data.Get("user_id")
	idInt, err := strconv.Atoi(id)
	if err != nil {
		return fmt.Errorf("no id", err)
	}
	u.Id = uint64(idInt)
	return nil
}
