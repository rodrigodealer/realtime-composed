package models

import (
	"encoding/json"
)

type FacebookUpdate struct {
	Object string                `json:"object"`
	Entry  []FacebookUpdateEntry `json:"entry"`
}

type FacebookUpdateEntry struct {
	UID           string   `json:"uid"`
	ID            string   `json:"id"`
	Time          int64    `json:"time"`
	ChangedFields []string `json:"changed_fields"`
}

func (m *FacebookUpdate) FromJson(value string) error {
	err := json.Unmarshal([]byte(value), &m)
	return err
}
