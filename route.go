package main

import "github.com/google/uuid"

type Car struct {
	Id      uuid.UUID
	Model   string
	Brand   string
	Hp      int
	License string
}
