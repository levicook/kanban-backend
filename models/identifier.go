package models

import (
	"database/sql"
	"database/sql/driver"
	"fmt"
	"strings"

	"code.google.com/p/go-uuid/uuid"
)

type identifier string

func newIdentifier() identifier {
	return identifier(strings.ToLower(uuid.New()))
}

func (id identifier) Blank() bool   { return id == "" }
func (id identifier) Present() bool { return id != "" }

func (id identifier) Invalid() bool { return !id.Valid() }
func (id identifier) Valid() bool {
	return len(id) == 36 && id[8] == '-' && id[13] == '-' && id[18] == '-' && id[23] == '-'
}

func (id identifier) MarshalJSON() ([]byte, error) {
	return []byte(fmt.Sprintf("%q", id)), nil
}

func (id *identifier) UnmarshalJSON(data []byte) error {
	if len(data) == 38 {
		*id = identifier(data[1:37])
	}

	return nil
}

func (id identifier) Value() (driver.Value, error) {
	if id.Blank() {
		return nil, nil
	}

	return string(id), nil
}

func (id *identifier) Scan(src interface{}) error {
	var ns sql.NullString

	if err := ns.Scan(src); err != nil {
		return err
	}

	*id = identifier(ns.String)

	return nil
}
