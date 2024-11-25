package model

import (
	"database/sql"
	"database/sql/driver"
	"errors"
	"fmt"
	"strings"

	"gopkg.in/yaml.v3"
)

type EntityRef struct {
	Kind      string
	Namespace string
	Name      string
}

func (e EntityRef) String() string {
	var b strings.Builder

	if e.Kind != "" {
		b.WriteString(e.Kind + ":")
	}
	if e.Namespace != "" {
		b.WriteString(e.Namespace + "/")
	}
	b.WriteString(e.Name)
	return b.String()
}

func MakeEntityRef(s string) (EntityRef, error) {
	var e EntityRef
	colidx := strings.Index(s, ":")
	if colidx != -1 {
		e.Kind = s[0:colidx]
		s = s[colidx+1:]
		if e.Kind == "" {
			return e, fmt.Errorf("kind empty in %s", s)
		}
	}
	slashidx := strings.Index(s, "/")
	if slashidx != -1 {
		e.Namespace = s[0:slashidx]
		s = s[slashidx+1:]
		if e.Namespace == "" {
			return e, fmt.Errorf("namespace empty in %s", s)
		}
	}
	e.Name = s
	if e.Name == "" {
		return e, fmt.Errorf("name empty in %s", s)
	}
	return e, nil
}

func (e EntityRef) Empty() bool {
	return e.Kind == "" && e.Namespace == "" && e.Name == ""
}

func (e EntityRef) MarshalYAML() (any, error) {
	return e.String(), nil
}

func (e *EntityRef) UnmarshalYAML(value *yaml.Node) error {
	if value.Tag != "!!str" {
		return fmt.Errorf("cannot unmarshal entity reference from type %s", value.Tag)
	}
	ref, err := MakeEntityRef(value.Value)
	if err != nil {
		return fmt.Errorf("failed to unmarshal entity reference from %s: %w", value.Value, err)
	}
	e.Kind = ref.Kind
	e.Namespace = ref.Namespace
	e.Name = ref.Name
	return nil
}

var _ yaml.Marshaler = EntityRef{}
var _ yaml.Unmarshaler = &EntityRef{}

func (e *EntityRef) Scan(src any) error {
	if src == nil {
		return nil
	}
	s, ok := src.(string)
	if !ok {
		return errors.New("found non-string value for entity ref")
	}

	entityRef, err := MakeEntityRef(s)
	if err != nil {
		return fmt.Errorf("failed to make entity ref from %s: %w", s, err)
	}
	e.Kind = entityRef.Kind
	e.Namespace = entityRef.Namespace
	e.Name = entityRef.Name
	return nil
}

func (e EntityRef) Value() (driver.Value, error) {
	if e.Empty() {
		return nil, nil
	}
	return e.String(), nil
}

var _ sql.Scanner = &EntityRef{}
var _ driver.Valuer = EntityRef{}

type EntityRefs struct {
	items []EntityRef
}

func MakeEntityRefs(items []EntityRef) EntityRefs {
	return EntityRefs{
		items: items,
	}
}

func (es EntityRefs) Items() []EntityRef {
	return es.items
}

func (es *EntityRefs) Scan(src any) error {
	if src == nil {
		return nil
	}
	ss, ok := src.(string)
	if !ok {
		return errors.New("found non-string value for entity ref slice")
	}

	srcStrings := strings.Split(ss, " ")
	scanned := make([]EntityRef, len(srcStrings))
	for i, s := range srcStrings {
		e := EntityRef{}
		if err := e.Scan(s); err != nil {
			return err
		}
		scanned[i] = e
	}

	es.items = scanned
	return nil
}

func (e EntityRefs) Value() (driver.Value, error) {
	if len(e.items) == 0 {
		return nil, nil
	}
	valStrings := make([]string, 0, len(e.items))
	for _, e := range e.items {
		if !e.Empty() {
			valStrings = append(valStrings, e.String())
		}
	}
	return strings.Join(valStrings, " "), nil
}

var _ sql.Scanner = &EntityRefs{}
var _ driver.Valuer = EntityRefs{}
