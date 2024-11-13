package store

import (
	"maps"
	"slices"
	"strings"

	"github.com/bhavanki/rewind/pkg/model"
	"gorm.io/gorm"
)

type Entity struct {
	gorm.Model
	APIVersion  string
	Kind        string `gorm:"uniqueIndex:entity_ref,priority:10"`
	Name        string `gorm:"uniqueIndex:entity_ref,priority:30"`
	Namespace   string `gorm:"uniqueIndex:entity_ref,priority:20"`
	Title       *string
	Description *string
	Labels      []Label
	Annotations []Annotation
	Tags        *string
	Links       []Link
}

type Label struct {
	gorm.Model
	Key      string
	Value    string
	EntityID uint // foreign key
}

type Annotation struct {
	gorm.Model
	Key      string
	Value    string
	EntityID uint // foreign key
}

type Link struct {
	gorm.Model
	URL      string
	Title    *string
	Icon     *string
	Type     *string
	EntityID uint // foreign key
}

// ---

func FromDBModel(e Entity) (model.Entity, error) {
	me := model.Entity{
		APIVersion: e.APIVersion,
		Kind:       e.Kind,
		Metadata: model.Metadata{
			Name:      e.Name,
			Namespace: e.Namespace,
		},
	}

	if e.Title != nil {
		me.Metadata.Title = *e.Title
	}
	if e.Description != nil {
		me.Metadata.Description = *e.Description
	}
	if e.Tags != nil {
		me.Metadata.Tags = strings.Split(*e.Tags, ",")
	}

	if len(e.Labels) > 0 {
		me.Metadata.Labels = make(map[string]string, len(e.Labels))
		for _, l := range e.Labels {
			me.Metadata.Labels[l.Key] = l.Value
		}
	}
	if len(e.Annotations) > 0 {
		me.Metadata.Annotations = make(map[string]string, len(e.Annotations))
		for _, a := range e.Annotations {
			me.Metadata.Annotations[a.Key] = a.Value
		}
	}
	if len(e.Links) > 0 {
		me.Metadata.Links = make([]model.Link, len(e.Links))
		for i := range e.Links {
			me.Metadata.Links[i] = model.Link{
				URL: e.Links[i].URL,
			}
			if e.Links[i].Title != nil {
				me.Metadata.Links[i].Title = *e.Links[i].Title
			}
			if e.Links[i].Icon != nil {
				me.Metadata.Links[i].Icon = *e.Links[i].Icon
			}
			if e.Links[i].Type != nil {
				me.Metadata.Links[i].Type = *e.Links[i].Type
			}
		}
	}

	return me, nil
}

func ToDBModel(me model.Entity) (Entity, error) {
	e := Entity{
		APIVersion: me.APIVersion,
		Kind:       me.Kind,
		Name:       me.Metadata.Name,
		Namespace:  me.Metadata.Namespace,
	}

	if me.Metadata.Title != "" {
		e.Title = &me.Metadata.Title
	}
	if me.Metadata.Description != "" {
		e.Description = &me.Metadata.Description
	}
	if len(me.Metadata.Tags) > 0 {
		t := strings.Join(me.Metadata.Tags, ",")
		e.Tags = &t
	}

	if len(me.Metadata.Labels) > 0 {
		e.Labels = make([]Label, 0, len(me.Metadata.Labels))
		labelKeys := slices.Sorted(maps.Keys(me.Metadata.Labels))
		for _, k := range labelKeys {
			e.Labels = append(e.Labels, Label{
				Key:   k,
				Value: me.Metadata.Labels[k],
			})
		}
	}
	if len(me.Metadata.Annotations) > 0 {
		e.Annotations = make([]Annotation, 0, len(me.Metadata.Annotations))
		annotationKeys := slices.Sorted(maps.Keys(me.Metadata.Annotations))
		for _, k := range annotationKeys {
			e.Annotations = append(e.Annotations, Annotation{
				Key:   k,
				Value: me.Metadata.Annotations[k],
			})
		}
	}

	if len(me.Metadata.Links) > 0 {
		e.Links = make([]Link, len(me.Metadata.Links))
		for i := range me.Metadata.Links {
			e.Links[i] = Link{
				URL: me.Metadata.Links[i].URL,
			}
			if me.Metadata.Links[i].Title != "" {
				e.Links[i].Title = &me.Metadata.Links[i].Title
			}
			if me.Metadata.Links[i].Icon != "" {
				e.Links[i].Icon = &me.Metadata.Links[i].Icon
			}
			if me.Metadata.Links[i].Type != "" {
				e.Links[i].Type = &me.Metadata.Links[i].Type
			}
		}
	}

	return e, nil
}
