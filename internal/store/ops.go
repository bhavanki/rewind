package store

import (
	"database/sql"
	"fmt"
	"strings"

	"github.com/bhavanki/rewind/pkg/model"
	"github.com/jmoiron/sqlx"
)

var (
	entityInsertStatement = `INSERT INTO entity (apiVersion, kind, namespace, name, title, description, tags) VALUES (?, ?, ?, ?, ?, ?, ?)`
	entityIDStatement     = `SELECT id FROM entity WHERE kind = ? AND namespace = ? AND name = ?`
	entityReadStatement   = `SELECT id, apiVersion, kind, namespace, name, title, description, tags FROM entity WHERE kind = ? AND namespace = ? AND name = ?`

	labelInsertStatement      = `INSERT INTO label (entity_id, k, v) VALUES (?, ?, ?)`
	labelSelectStatement      = `SELECT k, v FROM label WHERE entity_id = ?`
	annotationInsertStatement = `INSERT INTO annotation (entity_id, k, v) VALUES (?, ?, ?)`
	annotationSelectStatement = `SELECT k, v FROM annotation WHERE entity_id = ?`
	linkInsertStatement       = `INSERT INTO link (entity_id, url, title, icon, type) VALUES (?, ?, ?, ?, ?)`
	linkSelectStatement       = `SELECT url, title, icon, type FROM link WHERE entity_id = ?`

	entityDeleteStatement = `DELETE FROM entity WHERE id = ?`
)

func createEntity(e model.Entity, tx *sql.Tx) (model.Entity, error) {
	tags := strings.Join(e.Metadata.Tags, ",")
	result, err := tx.Exec(
		entityInsertStatement,
		e.APIVersion,
		e.Kind,
		e.Metadata.Namespace,
		e.Metadata.Name,
		nullString(e.Metadata.Title),
		nullString(e.Metadata.Description),
		nullString(tags),
	)
	if err != nil {
		return model.Entity{}, fmt.Errorf("failed to create entity: %w", err)
	}

	re := e
	id, err := result.LastInsertId()
	if err == nil && id > 0 {
		re.ID = id
	} else {
		return model.Entity{}, fmt.Errorf("failed to get new entity ID: %w", err)
	}

	for labelKey, labelValue := range e.Metadata.Labels {
		_, err := tx.Exec(
			labelInsertStatement,
			id,
			labelKey,
			labelValue,
		)
		if err != nil {
			return model.Entity{}, fmt.Errorf("failed to create label: %w", err)
		}
	}
	for annotationKey, annotationValue := range e.Metadata.Annotations {
		_, err := tx.Exec(
			annotationInsertStatement,
			id,
			annotationKey,
			annotationValue,
		)
		if err != nil {
			return model.Entity{}, fmt.Errorf("failed to create annotation: %w", err)
		}
	}
	for _, link := range e.Metadata.Links {
		_, err := tx.Exec(
			linkInsertStatement,
			id,
			link.URL,
			nullString(link.Title),
			nullString(link.Icon),
			nullString(link.Type),
		)
		if err != nil {
			return model.Entity{}, fmt.Errorf("failed to create link: %w", err)
		}
	}

	return re, nil
}

func getEntityID(ref model.EntityRef, db *sqlx.DB) (int64, error) {
	rows, err := db.Queryx(entityIDStatement, ref.Kind, ref.Namespace, ref.Name)
	if err != nil {
		return 0, fmt.Errorf("failed to query for entity ID: %w", err)
	}
	defer rows.Close()
	if !rows.Next() {
		return 0, fmt.Errorf("entity not found")
	}
	var id int64
	err = rows.Scan(&id)
	if err != nil {
		return 0, fmt.Errorf("failed to scan ID column for entity: %w", err)
	}
	return id, nil
}

func readEntity(ref model.EntityRef, db *sqlx.DB) (model.Entity, error) {
	rows, err := db.Queryx(entityReadStatement, ref.Kind, ref.Namespace, ref.Name)
	if err != nil {
		return model.Entity{}, fmt.Errorf("failed to query for entity: %w", err)
	}
	if !rows.Next() {
		rows.Close()
		return model.Entity{}, fmt.Errorf("entity not found")
	}
	var id int64
	var apiVersion string
	var kind string
	var namespace string
	var name string
	var title sql.NullString
	var description sql.NullString
	var tags sql.NullString
	err = rows.Scan(&id, &apiVersion, &kind, &namespace, &name, &title, &description, &tags)
	rows.Close()
	if err != nil {
		return model.Entity{}, fmt.Errorf("failed to scan columns for entity: %w", err)
	}

	e := model.Entity{
		ID:         id,
		APIVersion: apiVersion,
		Kind:       kind,
		Metadata: model.Metadata{
			Name:        name,
			Namespace:   namespace,
			Title:       fromNullString(title),
			Description: fromNullString(description),
		},
	}
	if tags.Valid {
		e.Metadata.Tags = strings.Split(tags.String, ",")
	}

	lrows, err := db.Queryx(labelSelectStatement, id)
	if err != nil {
		return model.Entity{}, fmt.Errorf("failed to query for labels: %w", err)
	}
	for lrows.Next() {
		if e.Metadata.Labels == nil {
			e.Metadata.Labels = make(map[string]string)
		}
		var k string
		var v string
		err = lrows.Scan(&k, &v)
		if err != nil {
			return model.Entity{}, fmt.Errorf("failed to scan columns for label: %w", err)
		}
		e.Metadata.Labels[k] = v
	}

	arows, err := db.Queryx(annotationSelectStatement, id)
	if err != nil {
		return model.Entity{}, fmt.Errorf("failed to query for annotations: %w", err)
	}
	for arows.Next() {
		if e.Metadata.Annotations == nil {
			e.Metadata.Annotations = make(map[string]string)
		}
		var k string
		var v string
		err = arows.Scan(&k, &v)
		if err != nil {
			return model.Entity{}, fmt.Errorf("failed to scan columns for annotation: %w", err)
		}
		e.Metadata.Annotations[k] = v
	}

	krows, err := db.Queryx(linkSelectStatement, id)
	if err != nil {
		return model.Entity{}, fmt.Errorf("failed to query for links: %w", err)
	}
	for krows.Next() {
		var url string
		var title sql.NullString
		var icon sql.NullString
		var linkType sql.NullString
		err = krows.Scan(&url, &title, &icon, &linkType)
		if err != nil {
			return model.Entity{}, fmt.Errorf("failed to scan columns for link: %w", err)
		}
		e.Metadata.Links = append(e.Metadata.Links, model.Link{
			URL:   url,
			Title: fromNullString(title),
			Icon:  fromNullString(icon),
			Type:  fromNullString(linkType),
		})
	}

	return e, nil
}

func deleteEntity(id int64, db *sqlx.DB) error {
	_, err := db.Exec(entityDeleteStatement, id)
	if err != nil {
		return fmt.Errorf("failed to delete entity: %w", err)
	}

	return nil
}

func nullString(s string) sql.NullString {
	if s == "" {
		return sql.NullString{
			Valid: false,
		}
	}
	return sql.NullString{
		String: s,
		Valid:  true,
	}
}
func fromNullString(ns sql.NullString) string {
	if ns.Valid {
		return ns.String
	}
	return ""
}
