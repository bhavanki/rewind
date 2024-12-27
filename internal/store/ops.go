package store

import (
	"database/sql"
	"fmt"
	"slices"
	"strings"

	"github.com/bhavanki/rewind/pkg/model"
	"github.com/jmoiron/sqlx"
)

var (
	entityInsertStatement = `INSERT INTO entity (apiVersion, kind, namespace, name, title, description, tags) VALUES (?, ?, ?, ?, ?, ?, ?)`
	// entityIDStatement     = `SELECT id FROM entity WHERE kind = ? AND namespace = ? AND name = ?`
	entityReadStatement   = `SELECT id, apiVersion, kind, namespace, name, title, description, tags FROM entity WHERE kind = ? AND namespace = ? AND name = ?`
	entityUpdateStatement = `UPDATE entity SET (apiVersion, kind, namespace, name, title, description, tags) = (?, ?, ?, ?, ?, ?, ?) WHERE id = ?`

	labelInsertStatement      = `INSERT INTO label (entity_id, k, v) VALUES (?, ?, ?)`
	labelSelectStatement      = `SELECT k, v FROM label WHERE entity_id = ?`
	labelUpdateStatement      = `UPDATE label SET v = ? WHERE entity_id = ? AND k = ?`
	labelDeleteStatement      = `DELETE FROM label WHERE entity_id = ? AND k = ?`
	annotationInsertStatement = `INSERT INTO annotation (entity_id, k, v) VALUES (?, ?, ?)`
	annotationSelectStatement = `SELECT k, v FROM annotation WHERE entity_id = ?`
	annotationUpdateStatement = `UPDATE annotation SET v = ? WHERE entity_id = ? AND k = ?`
	annotationDeleteStatement = `DELETE FROM annotation WHERE entity_id = ? AND k = ?`
	linkInsertStatement       = `INSERT INTO link (entity_id, idx, url, title, icon, type) VALUES (?, ?, ?, ?, ?, ?)`
	linkSelectStatement       = `SELECT url, title, icon, type FROM link WHERE entity_id = ? ORDER BY idx`
	linkUpdateStatement       = `UPDATE link SET (idx, title, icon, type) = (?, ?, ?, ?) WHERE entity_id = ? and url = ?`
	linkDeleteStatement       = `DELETE FROM link WHERE entity_id = ? and url = ?`

	entityDeleteStatement = `DELETE FROM entity WHERE id = ?`
)

func createEntity(e model.Entity, tx *sqlx.Tx) (model.Entity, error) {
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
	for i, link := range e.Metadata.Links {
		_, err := tx.Exec(
			linkInsertStatement,
			id,
			i,
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

// func getEntityID(ref model.EntityRef, db *sqlx.DB) (int64, error) {
// 	rows, err := db.Queryx(entityIDStatement, ref.Kind, ref.Namespace, ref.Name)
// 	if err != nil {
// 		return 0, fmt.Errorf("failed to query for entity ID: %w", err)
// 	}
// 	defer rows.Close()
// 	if !rows.Next() {
// 		return 0, fmt.Errorf("entity not found")
// 	}
// 	var id int64
// 	err = rows.Scan(&id)
// 	if err != nil {
// 		return 0, fmt.Errorf("failed to scan ID column for entity: %w", err)
// 	}
// 	return id, nil
// }

func readEntity(ref model.EntityRef, tx *sqlx.Tx) (model.Entity, error) {
	rows, err := tx.Queryx(entityReadStatement, ref.Kind, ref.Namespace, ref.Name)
	if err != nil {
		return model.Entity{}, fmt.Errorf("failed to query for entity: %w", err)
	}
	defer rows.Close()
	if !rows.Next() {
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
	if err != nil {
		return model.Entity{}, fmt.Errorf("failed to scan columns for entity: %w", err)
	}
	rows.Close()

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

	e.Metadata.Labels, err = readLabels(id, tx)
	if err != nil {
		return model.Entity{}, err
	}
	e.Metadata.Annotations, err = readAnnotations(id, tx)
	if err != nil {
		return model.Entity{}, err
	}
	e.Metadata.Links, err = readLinks(id, tx)
	if err != nil {
		return model.Entity{}, err
	}

	return e, nil
}

func readLabels(id int64, tx *sqlx.Tx) (map[string]string, error) {
	lrows, err := tx.Queryx(labelSelectStatement, id)
	if err != nil {
		return nil, fmt.Errorf("failed to query for labels: %w", err)
	}
	defer lrows.Close()

	labels := make(map[string]string)
	for lrows.Next() {
		var k string
		var v string
		err = lrows.Scan(&k, &v)
		if err != nil {
			return nil, fmt.Errorf("failed to scan columns for label: %w", err)
		}
		labels[k] = v
	}

	return labels, nil
}

func readAnnotations(id int64, tx *sqlx.Tx) (map[string]string, error) {
	arows, err := tx.Queryx(annotationSelectStatement, id)
	if err != nil {
		return nil, fmt.Errorf("failed to query for annotations: %w", err)
	}
	defer arows.Close()

	annotations := make(map[string]string)
	for arows.Next() {
		var k string
		var v string
		err = arows.Scan(&k, &v)
		if err != nil {
			return nil, fmt.Errorf("failed to scan columns for annotation: %w", err)
		}
		annotations[k] = v
	}

	return annotations, nil
}

func readLinks(id int64, tx *sqlx.Tx) ([]model.Link, error) {
	krows, err := tx.Queryx(linkSelectStatement, id)
	if err != nil {
		return nil, fmt.Errorf("failed to query for links: %w", err)
	}
	defer krows.Close()

	links := make([]model.Link, 0)
	for krows.Next() {
		var url string
		var title sql.NullString
		var icon sql.NullString
		var linkType sql.NullString
		err = krows.Scan(&url, &title, &icon, &linkType)
		if err != nil {
			return nil, fmt.Errorf("failed to scan columns for link: %w", err)
		}
		links = append(links, model.Link{
			URL:   url,
			Title: fromNullString(title),
			Icon:  fromNullString(icon),
			Type:  fromNullString(linkType),
		})
	}

	return links, nil
}

func updateEntity(e model.Entity, tx *sqlx.Tx) (model.Entity, error) {
	tags := strings.Join(e.Metadata.Tags, ",")
	_, err := tx.Exec(
		entityUpdateStatement,
		e.APIVersion,
		e.Kind,
		e.Metadata.Namespace,
		e.Metadata.Name,
		nullString(e.Metadata.Title),
		nullString(e.Metadata.Description),
		nullString(tags),
		e.ID,
	)
	if err != nil {
		return model.Entity{}, fmt.Errorf("failed to update entity: %w", err)
	}

	re := e

	currentLabels, err := readLabels(e.ID, tx)
	if err != nil {
		return model.Entity{}, fmt.Errorf("failed to read labels for updating entity: %w", err)
	}
	for labelKey, labelValue := range e.Metadata.Labels {
		_, exists := currentLabels[labelKey]
		if !exists {
			_, err := tx.Exec(
				labelInsertStatement,
				e.ID,
				labelKey,
				labelValue,
			)
			if err != nil {
				return model.Entity{}, fmt.Errorf("failed to create label: %w", err)
			}
		} else {
			_, err := tx.Exec(
				labelUpdateStatement,
				labelValue,
				e.ID,
				labelKey,
			)
			if err != nil {
				return model.Entity{}, fmt.Errorf("failed to update label: %w", err)
			}
		}
	}
	for labelKey := range currentLabels {
		_, remains := e.Metadata.Labels[labelKey]
		if !remains {
			_, err := tx.Exec(
				labelDeleteStatement,
				e.ID,
				labelKey,
			)
			if err != nil {
				return model.Entity{}, fmt.Errorf("failed to delete label: %w", err)
			}
		}
	}

	currentAnnotations, err := readAnnotations(e.ID, tx)
	if err != nil {
		return model.Entity{}, fmt.Errorf("failed to read annotations for updating entity: %w", err)
	}
	for annotationKey, annotationValue := range e.Metadata.Annotations {
		_, exists := currentAnnotations[annotationKey]
		if !exists {
			_, err := tx.Exec(
				annotationInsertStatement,
				e.ID,
				annotationKey,
				annotationValue,
			)
			if err != nil {
				return model.Entity{}, fmt.Errorf("failed to create annotation: %w", err)
			}
		} else {
			_, err := tx.Exec(
				annotationUpdateStatement,
				annotationValue,
				e.ID,
				annotationKey,
			)
			if err != nil {
				return model.Entity{}, fmt.Errorf("failed to update annotation: %w", err)
			}
		}
	}
	for annotationKey := range currentAnnotations {
		_, remains := e.Metadata.Annotations[annotationKey]
		if !remains {
			_, err := tx.Exec(
				annotationDeleteStatement,
				e.ID,
				annotationKey,
			)
			if err != nil {
				return model.Entity{}, fmt.Errorf("failed to delete annotation: %w", err)
			}
		}
	}

	currentLinks, err := readLinks(e.ID, tx)
	if err != nil {
		return model.Entity{}, fmt.Errorf("failed to read links for updating entity: %w", err)
	}
	currentLinkMap := make(map[string]model.Link, len(currentLinks))
	for _, link := range currentLinks {
		currentLinkMap[link.URL] = link
	}

	entityLinkURLs := make([]string, 0, len(e.Metadata.Links))
	for i, link := range e.Metadata.Links {
		_, exists := currentLinkMap[link.URL]
		if !exists {
			_, err := tx.Exec(
				linkInsertStatement,
				e.ID,
				i,
				link.URL,
				nullString(link.Title),
				nullString(link.Icon),
				nullString(link.Type),
			)
			if err != nil {
				return model.Entity{}, fmt.Errorf("failed to create link: %w", err)
			}
		} else {
			_, err := tx.Exec(
				linkUpdateStatement,
				i,
				nullString(link.Title),
				nullString(link.Icon),
				nullString(link.Type),
				e.ID,
				link.URL,
			)
			if err != nil {
				return model.Entity{}, fmt.Errorf("failed to update link: %w", err)
			}
		}
		entityLinkURLs = append(entityLinkURLs, link.URL)
	}
	for linkURL := range currentLinkMap {
		if !slices.Contains(entityLinkURLs, linkURL) {
			_, err := tx.Exec(
				linkDeleteStatement,
				e.ID,
				linkURL,
			)
			if err != nil {
				return model.Entity{}, fmt.Errorf("failed to delete link: %w", err)
			}
		}
	}

	return re, nil
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
