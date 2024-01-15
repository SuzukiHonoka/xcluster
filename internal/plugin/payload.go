package plugin

import (
	"fmt"
	"reflect"
)

type MappedPayloadDescription map[PayloadName]PayloadDescription

// MappedPayload is used for data input
type MappedPayload map[PayloadName]Payload

type PayloadName string
type PayloadNames []PayloadName

type PayloadType string
type PayloadValue interface{}

var (
	PayloadTypeString PayloadType = PayloadType(reflect.String.String()) // string
	PayloadTypeInt    PayloadType = PayloadType(reflect.Int.String())    // int
	PayloadTypeBool   PayloadType = PayloadType(reflect.Bool.String())   // bool
)

type PayloadDescription string

type Payloads []*Payload

type Payload struct {
	Name        PayloadName
	Type        PayloadType
	Description PayloadDescription
}

// Validate validates if PayloadValue match the PayloadDescription.
func (p *Payload) Validate(value PayloadValue) error {
	if value == nil {
		return fmt.Errorf("%w: name=%s nil", ErrPayloadValueNotFound, p.Name)
	}
	kind := reflect.TypeOf(value).Kind().String()
	if kind != string(p.Type) {
		return fmt.Errorf("%w: name=%s expect type=%s, got=%s", ErrPayloadValueTypeMismatch, p.Name, p.Type, kind)
	}
	return nil
}
