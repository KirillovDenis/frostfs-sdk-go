package container

import (
	"github.com/nspcc-dev/neofs-api-go/v2/container"
)

type (
	Attribute  container.Attribute
	Attributes []Attribute
)

// NewAttribute creates and initializes blank Attribute.
//
// Defaults:
//  - key: "";
//  - value: "".
func NewAttribute() *Attribute {
	return NewAttributeFromV2(new(container.Attribute))
}

func (a *Attribute) SetKey(v string) {
	(*container.Attribute)(a).SetKey(v)
}

func (a *Attribute) SetValue(v string) {
	(*container.Attribute)(a).SetValue(v)
}

func (a *Attribute) Key() string {
	return (*container.Attribute)(a).GetKey()
}

func (a *Attribute) Value() string {
	return (*container.Attribute)(a).GetValue()
}

// NewAttributeFromV2 wraps protocol dependent version of
// Attribute message.
//
// Nil container.Attribute converts to nil.
func NewAttributeFromV2(v *container.Attribute) *Attribute {
	return (*Attribute)(v)
}

// ToV2 converts Attribute to v2 Attribute message.
//
// Nil Attribute converts to nil.
func (a *Attribute) ToV2() *container.Attribute {
	return (*container.Attribute)(a)
}

func NewAttributesFromV2(v []container.Attribute) Attributes {
	if v == nil {
		return nil
	}

	attrs := make(Attributes, len(v))
	for i := range v {
		attrs[i] = *NewAttributeFromV2(&v[i])
	}

	return attrs
}

func (a Attributes) ToV2() []container.Attribute {
	if a == nil {
		return nil
	}

	attrs := make([]container.Attribute, len(a))
	for i := range a {
		attrs[i] = *a[i].ToV2()
	}

	return attrs
}

// sets value of the attribute by key.
func setAttribute(c *Container, key, value string) {
	attrs := c.Attributes()
	found := false

	for i := range attrs {
		if attrs[i].Key() == key {
			attrs[i].SetValue(value)
			found = true
			break
		}
	}

	if !found {
		index := len(attrs)
		attrs = append(attrs, Attribute{})
		attrs[index].SetKey(key)
		attrs[index].SetValue(value)
	}

	c.SetAttributes(attrs)
}

// iterates over container attributes. Stops at f's true return.
//
// Handler must not be nil.
func iterateAttributes(c *Container, f func(*Attribute) bool) {
	attrs := c.Attributes()
	for i := range attrs {
		if f(&attrs[i]) {
			return
		}
	}
}

// SetNativeNameWithZone sets container native name and its zone.
//
// Use SetNativeName to set default zone.
func SetNativeNameWithZone(c *Container, name, zone string) {
	setAttribute(c, container.SysAttributeName, name)
	setAttribute(c, container.SysAttributeZone, zone)
}

// SetNativeName sets container native name with default zone (container).
func SetNativeName(c *Container, name string) {
	SetNativeNameWithZone(c, name, container.SysAttributeZoneDefault)
}

// GetNativeNameWithZone returns container native name and its zone.
func GetNativeNameWithZone(c *Container) (name string, zone string) {
	iterateAttributes(c, func(a *Attribute) bool {
		if key := a.Key(); key == container.SysAttributeName {
			name = a.Value()
		} else if key == container.SysAttributeZone {
			zone = a.Value()
		}

		return name != "" && zone != ""
	})

	return
}
