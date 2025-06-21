package entity

import (
	"errors"
	"fmt"
	"github.com/google/uuid"
	"time"
)

type Member struct {
	ID         string
	Name       string
	MemberType string
	Created    time.Time
	Updated    time.Time
}

const (
	primaryMember   string = "Primary Member"
	secondaryMember string = "Secondary Member"
)

func NewMember(name, memberType string) (*Member, error) {
	if name == "" {
		return nil, errors.New(fmt.Sprintf("name of member is mandatory. name:[%s]", name))
	}
	if memberType == "" {
		memberType = primaryMember
	} else {
		memberType = secondaryMember
	}
	return &Member{
		ID:         uuid.New().String(),
		Name:       name,
		MemberType: memberType,
		Created:    time.Now().UTC(),
		Updated:    time.Now().UTC(),
	}, nil
}
