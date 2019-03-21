package models

// Store struct
type Store struct {
	Name   string  `json:"name"`
	Coords *Coords `json:"coords"`
}
