package db

import (
	"database/sql/driver"
	"encoding/json"
	"fmt"
)

type NullRawMessage struct {
	Valid      bool
	RawMessage json.RawMessage
}

func (n *NullRawMessage) Scan(src any) error {
	if src == nil {
		n.Valid = false
		return nil
	}
	switch src := src.(type) {
	case []byte:
		srcCopy := make([]byte, len(src))
		copy(srcCopy, src)
		n.RawMessage, n.Valid = srcCopy, true
	default:
		return fmt.Errorf("error scanning:  type %T to type %T", src, []byte{})
	}
	return nil
}

func (n NullRawMessage) Value() (driver.Value, error) {
	if !n.Valid {
		return nil, nil
	}
	return []byte(n.RawMessage), nil
}

type JSONType map[string]any

func (j *JSONType) Scan(value any) error {
	return json.Unmarshal(value.([]byte), j)
}

func (j JSONType) Value() (driver.Value, error) {
	return json.Marshal(j)
}
