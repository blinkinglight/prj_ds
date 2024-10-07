package types

type DataStore struct {
	Elf string `json:"elf"`
	Elv string `json:"elv"`
	Elr string `json:"elr"`
}

type Form struct {
	Name string `json:"name"`

	Roles []FormRole `json:"roles"`
}

type FormRole struct {
	Name     string `json:"name"`
	Category string `json:"category"`
}
