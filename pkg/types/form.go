package types

import "github.com/ituoga/toolbox"

type DataStore struct {
	Elf string `json:"elf"`
	Elv string `json:"elv"`
	Elr string `json:"elr"`
}

type Form struct {
	Name  string                        `json:"name"`
	Count int                           `json:"count"`
	Roles *toolbox.LinkedList[FormRole] `json:"roles"`
}

type FormRole struct {
	Id       int64  `json:"id"`
	Name     string `json:"name"`
	Valid    bool   `json:"valid"`
	Category string `json:"category"`
}
