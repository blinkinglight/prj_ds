package types

import (
	"encoding/json"
	"strconv"
)

type DataStore struct {
	Elf string `json:"elf"`
	Elv string `json:"elv"`
	Elr string `json:"elr"`
}

type Form struct {
	Name string `json:"name"`

	Roles *FormRole `json:"roles"`
}

type FormRole struct {
	Name     string `json:"name"`
	Valid    bool   `json:"valid"`
	Category string `json:"category"`

	FormRole *FormRole `json:"roles"`
}

type CustomID int

func (c CustomID) UnmarshalJSON(in []byte) error {
	var s string
	if err := json.Unmarshal(in, &s); err != nil {
		return err
	}
	i, err := strconv.Atoi(s)
	if err != nil {
		return err
	}
	c = CustomID(i)

	return nil
}

func (c CustomID) MarshalJSON() ([]byte, error) {
	return json.Marshal(int(c))
}
