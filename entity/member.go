package entity

import (
	"errors"
	"fmt"
	"time"

	"github.com/google/uuid"
)

type MemberType string

type Member struct {
	ID         string
	Name       string
	MemberType MemberType
	Created    time.Time
	Updated    time.Time
}

const (
	primaryMember   MemberType = "PRIMARY"
	secondaryMember MemberType = "SECONDARY"
)

func NewMember(name string, memberType MemberType) (*Member, error) {
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
