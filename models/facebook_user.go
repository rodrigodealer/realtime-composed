package models

import (
	"encoding/json"
)

type FacebookUser struct {
	UID  string `json:"uid"`
	ID   string `json:"id"`
	Name string `json:"name"`
}

func (m *FacebookUser) ToJson() string {
	userJSON, _ := json.Marshal(m)
	return string(userJSON)
}
