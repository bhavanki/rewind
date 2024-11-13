package model

import (
	"fmt"
	"strings"
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
