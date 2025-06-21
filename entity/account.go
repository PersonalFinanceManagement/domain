package entity

import (
	"fmt"
	"time"
)

type Account struct {
	ID         string
	Name       string
	LinkedBank string
	Balance    float64
	Created    time.Time
	Updated    time.Time
	Owners     []Member
	Metadata   Metadata
}
