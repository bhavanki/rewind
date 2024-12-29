package store

import (
	"fmt"
	"strings"

	"github.com/bhavanki/rewind/pkg/model"
	"github.com/jmoiron/sqlx"
	_ "github.com/mattn/go-sqlite3"
)

type sqliteStore struct {
	db *sqlx.DB
}

var _ Store = sqliteStore{}

func NewSqliteStore(connString string) (*sqliteStore, error) {
	db, err := sqlx.Open("sqlite3", connString+"?_fk=on")
	if err != nil {
		return nil, fmt.Errorf("failed to open database at %s: %w", connString, err)
	}

	if err = runMigrations(db.DB, "sqlite3"); err != nil {
		return nil, fmt.Errorf("failed to run migrations: %w", err)
	}

	return &sqliteStore{
		db: db,
	}, nil
}

// for testing
func (s sqliteStore) readEntity(ref model.EntityRef) (model.Entity, error) {
	tx, err := s.db.Beginx()
	if err != nil {
		return model.Entity{}, fmt.Errorf("failed to begin transaction for read: %w", err)
	}
	defer tx.Commit()

	entity, err := readEntity(ref, tx)
	if err != nil {
		return model.Entity{}, err
	}

	return entity, nil
}

var (
	componentInsertStatement = `INSERT INTO component (entity_id, type, lifecycle, owner, system, subcomponent_of, provides_apis, consumes_apis, depends_on, dependency_of) VALUES (?, ?, ?, ?, ?, ?, ?, ?, ?, ?)`
	componentSelectStatement = `SELECT type, lifecycle, owner, system, subcomponent_of, provides_apis, consumes_apis, depends_on, dependency_of FROM component WHERE entity_id = ?`
	componentUpdateStatement = `UPDATE component SET (type, lifecycle, owner, system, subcomponent_of, provides_apis, consumes_apis, depends_on, dependency_of) = (?, ?, ?, ?, ?, ?, ?, ?, ?) WHERE entity_id = ?`

	componentListStatementPrefix = `SELECT namespace, name FROM entity INNER JOIN component ON entity.id = component.entity_id WHERE entity.kind = ?`

	apiInsertStatement = `INSERT INTO api (entity_id, type, lifecycle, owner, system, definition) VALUES (?, ?, ?, ?, ?, ?)`
	apiSelectStatement = `SELECT type, lifecycle, owner, system, definition FROM api WHERE entity_id = ?`
	apiUpdateStatement = `UPDATE api SET (type, lifecycle, owner, system, definition) = (?, ?, ?, ?, ?) WHERE entity_id = ?`

	userInsertStatement = `INSERT INTO user (entity_id, display_name, email, picture, member_of) VALUES (?, ?, ?, ?, ?)`
	userSelectStatement = `SELECT display_name, email, picture, member_of FROM user WHERE entity_id = ?`
	userUpdateStatement = `UPDATE user SET (display_name, email, picture, member_of) = (?, ?, ?, ?) WHERE entity_id = ?`

	groupInsertStatement = `INSERT INTO grp (entity_id, type, display_name, email, picture, parent, children, members) VALUES (?, ?, ?, ?, ?, ?, ?, ?)`
	groupSelectStatement = `SELECT type, display_name, email, picture, parent, children, members FROM grp WHERE entity_id = ?`
	groupUpdateStatement = `UPDATE grp SET (type, display_name, email, picture, parent, children, members) = (?, ?, ?, ?, ?, ?, ?) WHERE entity_id = ?`
)

func (s sqliteStore) CreateComponent(c model.Component) (rc model.Component, err error) {
	var tx *sqlx.Tx
	tx, err = s.db.Beginx()
	if err != nil {
		return model.Component{}, fmt.Errorf("failed to begin transaction for create: %w", err)
	}
	defer func() {
		if err != nil {
			rerr := tx.Rollback()
			if rerr != nil {
				err = fmt.Errorf("failed to rollback transaction (%s): %w", rerr, err)
			}
		}
	}()

	entity, err := createEntity(c.Entity, tx)
	if err != nil {
		return model.Component{}, err
	}

	rc = c
	rc.Entity.ID = entity.ID

	_, err = tx.Exec(
		componentInsertStatement,
		entity.ID,
		c.Spec.Type,
		c.Spec.Lifecycle,
		c.Spec.Owner,
		c.Spec.System,
		c.Spec.SubcomponentOf,
		model.MakeEntityRefs(c.Spec.ProvidesAPIs),
		model.MakeEntityRefs(c.Spec.ConsumesAPIs),
		model.MakeEntityRefs(c.Spec.DependsOn),
		model.MakeEntityRefs(c.Spec.DependencyOf),
	)
	if err != nil {
		return model.Component{}, fmt.Errorf("failed to create component: %w", err)
	}

	if err = tx.Commit(); err != nil {
		return model.Component{}, fmt.Errorf("failed to commit transaction for create: %w", err)
	}
	return rc, nil
}

func (s sqliteStore) ReadComponent(ref model.EntityRef) (c model.Component, err error) {
	var tx *sqlx.Tx
	tx, err = s.db.Beginx()
	if err != nil {
		return model.Component{}, fmt.Errorf("failed to begin transaction for read: %w", err)
	}
	defer func() {
		if err != nil {
			rerr := tx.Rollback()
			if rerr != nil {
				err = fmt.Errorf("failed to rollback transaction (%s): %w", rerr, err)
			}
		}
	}()

	entity, err := readEntity(ref, tx)
	if err != nil {
		return model.Component{}, err
	}

	c = model.Component{
		Entity: entity,
	}

	rows, err := tx.Queryx(componentSelectStatement, entity.ID)
	if err != nil {
		return model.Component{}, fmt.Errorf("failed to query for component: %w", err)
	}
	defer rows.Close()
	if rows.Next() {
		var componentType string
		var lifecycle string
		var owner model.EntityRef
		var system model.EntityRef
		var subcomponentOf model.EntityRef
		var providesAPIs model.EntityRefs
		var consumesAPIs model.EntityRefs
		var dependsOn model.EntityRefs
		var dependencyOf model.EntityRefs

		err = rows.Scan(&componentType, &lifecycle, &owner, &system, &subcomponentOf, &providesAPIs, &consumesAPIs, &dependsOn, &dependencyOf)
		if err != nil {
			return model.Component{}, fmt.Errorf("failed to scan columns for component: %w", err)
		}
		c.Spec = model.ComponentSpec{
			Type:           componentType,
			Lifecycle:      lifecycle,
			Owner:          owner,
			System:         system,
			SubcomponentOf: subcomponentOf,
			ProvidesAPIs:   providesAPIs.Items(),
			ConsumesAPIs:   consumesAPIs.Items(),
			DependsOn:      dependsOn.Items(),
			DependencyOf:   dependencyOf.Items(),
		}
	}

	if err = tx.Commit(); err != nil {
		return model.Component{}, fmt.Errorf("failed to commit transaction for read: %w", err)
	}
	return c, nil
}

func (s sqliteStore) UpdateComponent(c model.Component) (rc model.Component, err error) {
	var tx *sqlx.Tx
	tx, err = s.db.Beginx()
	if err != nil {
		return model.Component{}, fmt.Errorf("failed to begin transaction for update: %w", err)
	}
	defer func() {
		if err != nil {
			rerr := tx.Rollback()
			if rerr != nil {
				err = fmt.Errorf("failed to rollback transaction (%s): %w", rerr, err)
			}
		}
	}()

	entity, err := updateEntity(c.Entity, tx)
	if err != nil {
		return model.Component{}, err
	}

	rc = c
	rc.Entity.ID = entity.ID

	_, err = tx.Exec(
		componentUpdateStatement,
		c.Spec.Type,
		c.Spec.Lifecycle,
		c.Spec.Owner,
		c.Spec.System,
		c.Spec.SubcomponentOf,
		model.MakeEntityRefs(c.Spec.ProvidesAPIs),
		model.MakeEntityRefs(c.Spec.ConsumesAPIs),
		model.MakeEntityRefs(c.Spec.DependsOn),
		model.MakeEntityRefs(c.Spec.DependencyOf),
		entity.ID,
	)
	if err != nil {
		return model.Component{}, fmt.Errorf("failed to update component: %w", err)
	}

	if err = tx.Commit(); err != nil {
		return model.Component{}, fmt.Errorf("failed to commit transaction for update: %w", err)
	}
	return rc, nil
}

func (s sqliteStore) DeleteComponent(ref model.EntityRef) (model.Component, error) {
	component, err := s.ReadComponent(ref)
	if err != nil {
		return model.Component{}, err
	}

	err = deleteEntity(component.Entity.ID, s.db)
	if err != nil {
		return model.Component{}, fmt.Errorf("failed to delete component: %w", err)
	}

	return component, nil
}

func (s sqliteStore) ListComponents(filters []Filter, ordering Ordering, pagination Pagination) ([]model.EntityRef, Pagination, error) {
	whereClauses := []string{}
	queryParameters := []any{
		model.KindComponent,
	}
	for _, filter := range filters {
		whereClauses = append(whereClauses, fmt.Sprintf("%s = ?", filter.Key))
		queryParameters = append(queryParameters, filter.Value)
	}

	listStatement := componentListStatementPrefix
	if len(whereClauses) > 0 {
		listStatement += " AND "
		listStatement += strings.Join(whereClauses, " AND ")
	}

	if ordering.OrderBy != "" {
		orderByClause := " ORDER BY " + string(ordering.OrderBy)
		if ordering.Descending {
			orderByClause += " DESC"
		} else {
			orderByClause += " ASC"
		}
		listStatement += orderByClause
	}

	if pagination.Limit > 0 {
		limitClause := fmt.Sprintf(" LIMIT %d", pagination.Limit)
		if pagination.Offset > 0 {
			limitClause += fmt.Sprintf(" OFFSET %d", pagination.Offset)
		}
		listStatement += limitClause
	}

	rows, err := s.db.Queryx(listStatement, queryParameters...)
	if err != nil {
		return nil, Pagination{}, fmt.Errorf("failed to list components: %w", err)
	}
	defer rows.Close()
	results := []model.EntityRef{}
	nextOffset := pagination.Offset
	for rows.Next() {
		var namespace string
		var name string
		err = rows.Scan(&namespace, &name)
		if err != nil {
			return nil, Pagination{}, fmt.Errorf("failed to scan columns for component: %w", err)
		}
		results = append(results, model.EntityRef{
			Kind:      model.KindComponent,
			Namespace: namespace,
			Name:      name,
		})
		nextOffset++
	}

	return results, Pagination{
		Limit:  pagination.Limit,
		Offset: nextOffset,
	}, nil
}

// ---

func (s sqliteStore) CreateAPI(a model.API) (ra model.API, err error) {
	var tx *sqlx.Tx
	tx, err = s.db.Beginx()
	if err != nil {
		return model.API{}, fmt.Errorf("failed to begin transaction for create: %w", err)
	}
	defer func() {
		if err != nil {
			rerr := tx.Rollback()
			if rerr != nil {
				err = fmt.Errorf("failed to rollback transaction (%s): %w", rerr, err)
			}
		}
	}()

	entity, err := createEntity(a.Entity, tx)
	if err != nil {
		return model.API{}, err
	}

	ra = a
	ra.Entity.ID = entity.ID

	_, err = tx.Exec(
		apiInsertStatement,
		entity.ID,
		a.Spec.Type,
		a.Spec.Lifecycle,
		a.Spec.Owner,
		a.Spec.System,
		a.Spec.Definition,
	)
	if err != nil {
		return model.API{}, fmt.Errorf("failed to create API: %w", err)
	}

	if err = tx.Commit(); err != nil {
		return model.API{}, fmt.Errorf("failed to commit transaction for create: %w", err)
	}
	return ra, nil
}

func (s sqliteStore) ReadAPI(ref model.EntityRef) (a model.API, err error) {
	var tx *sqlx.Tx
	tx, err = s.db.Beginx()
	if err != nil {
		return model.API{}, fmt.Errorf("failed to begin transaction for read: %w", err)
	}
	defer func() {
		if err != nil {
			rerr := tx.Rollback()
			if rerr != nil {
				err = fmt.Errorf("failed to rollback transaction (%s): %w", rerr, err)
			}
		}
	}()

	entity, err := readEntity(ref, tx)
	if err != nil {
		return model.API{}, err
	}

	a = model.API{
		Entity: entity,
	}

	rows, err := tx.Queryx(apiSelectStatement, entity.ID)
	if err != nil {
		return model.API{}, fmt.Errorf("failed to query for API: %w", err)
	}
	defer rows.Close()
	if rows.Next() {
		var apiType string
		var lifecycle string
		var owner model.EntityRef
		var system model.EntityRef
		var definition string
		err = rows.Scan(&apiType, &lifecycle, &owner, &system, &definition)
		if err != nil {
			return model.API{}, fmt.Errorf("failed to scan columns for API: %w", err)
		}
		a.Spec = model.APISpec{
			Type:       apiType,
			Lifecycle:  lifecycle,
			Owner:      owner,
			System:     system,
			Definition: definition,
		}
	}

	if err = tx.Commit(); err != nil {
		return model.API{}, fmt.Errorf("failed to commit transaction for read: %w", err)
	}
	return a, nil
}

func (s sqliteStore) UpdateAPI(a model.API) (ra model.API, err error) {
	var tx *sqlx.Tx
	tx, err = s.db.Beginx()
	if err != nil {
		return model.API{}, fmt.Errorf("failed to begin transaction for update: %w", err)
	}
	defer func() {
		if err != nil {
			rerr := tx.Rollback()
			if rerr != nil {
				err = fmt.Errorf("failed to rollback transaction (%s): %w", rerr, err)
			}
		}
	}()

	entity, err := updateEntity(a.Entity, tx)
	if err != nil {
		return model.API{}, err
	}

	ra = a
	ra.Entity.ID = entity.ID

	_, err = tx.Exec(
		apiUpdateStatement,
		a.Spec.Type,
		a.Spec.Lifecycle,
		a.Spec.Owner,
		a.Spec.System,
		a.Spec.Definition,
		entity.ID,
	)
	if err != nil {
		return model.API{}, fmt.Errorf("failed to update API: %w", err)
	}

	if err = tx.Commit(); err != nil {
		return model.API{}, fmt.Errorf("failed to commit transaction for update: %w", err)
	}
	return ra, nil
}

func (s sqliteStore) DeleteAPI(ref model.EntityRef) (model.API, error) {
	api, err := s.ReadAPI(ref)
	if err != nil {
		return model.API{}, err
	}

	err = deleteEntity(api.Entity.ID, s.db)
	if err != nil {
		return model.API{}, fmt.Errorf("failed to delete API: %w", err)
	}

	return api, nil
}

// ---

func (s sqliteStore) CreateUser(u model.User) (ru model.User, err error) {
	var tx *sqlx.Tx
	tx, err = s.db.Beginx()
	if err != nil {
		return model.User{}, fmt.Errorf("failed to begin transaction for create: %w", err)
	}
	defer func() {
		if err != nil {
			rerr := tx.Rollback()
			if rerr != nil {
				err = fmt.Errorf("failed to rollback transaction (%s): %w", rerr, err)
			}
		}
	}()

	entity, err := createEntity(u.Entity, tx)
	if err != nil {
		return model.User{}, err
	}

	ru = u
	ru.Entity.ID = entity.ID

	_, err = tx.Exec(
		userInsertStatement,
		entity.ID,
		u.Spec.Profile.DisplayName,
		u.Spec.Profile.Email,
		u.Spec.Profile.Picture,
		model.MakeEntityRefs(u.Spec.MemberOf),
	)
	if err != nil {
		return model.User{}, fmt.Errorf("failed to create user: %w", err)
	}

	if err = tx.Commit(); err != nil {
		return model.User{}, fmt.Errorf("failed to commit transaction for create: %w", err)
	}
	return ru, nil
}

func (s sqliteStore) ReadUser(ref model.EntityRef) (u model.User, err error) {
	var tx *sqlx.Tx
	tx, err = s.db.Beginx()
	if err != nil {
		return model.User{}, fmt.Errorf("failed to begin transaction for read: %w", err)
	}
	defer func() {
		if err != nil {
			rerr := tx.Rollback()
			if rerr != nil {
				err = fmt.Errorf("failed to rollback transaction (%s): %w", rerr, err)
			}
		}
	}()

	entity, err := readEntity(ref, tx)
	if err != nil {
		return model.User{}, err
	}

	u = model.User{
		Entity: entity,
	}

	rows, err := tx.Queryx(userSelectStatement, entity.ID)
	if err != nil {
		return model.User{}, fmt.Errorf("failed to query for user: %w", err)
	}
	defer rows.Close()
	if rows.Next() {
		var displayName string
		var email string
		var picture string
		var memberOf model.EntityRefs
		err = rows.Scan(&displayName, &email, &picture, &memberOf)
		if err != nil {
			return model.User{}, fmt.Errorf("failed to scan columns for user: %w", err)
		}
		u.Spec = model.UserSpec{
			Profile: model.UserProfile{
				DisplayName: displayName,
				Email:       email,
				Picture:     picture,
			},
			MemberOf: memberOf.Items(),
		}
	}

	if err = tx.Commit(); err != nil {
		return model.User{}, fmt.Errorf("failed to commit transaction for read: %w", err)
	}
	return u, nil
}

func (s sqliteStore) UpdateUser(u model.User) (ru model.User, err error) {
	var tx *sqlx.Tx
	tx, err = s.db.Beginx()
	if err != nil {
		return model.User{}, fmt.Errorf("failed to begin transaction for update: %w", err)
	}
	defer func() {
		if err != nil {
			rerr := tx.Rollback()
			if rerr != nil {
				err = fmt.Errorf("failed to rollback transaction (%s): %w", rerr, err)
			}
		}
	}()

	entity, err := updateEntity(u.Entity, tx)
	if err != nil {
		return model.User{}, err
	}

	ru = u
	ru.Entity.ID = entity.ID

	_, err = tx.Exec(
		userUpdateStatement,
		u.Spec.Profile.DisplayName,
		u.Spec.Profile.Email,
		u.Spec.Profile.Picture,
		model.MakeEntityRefs(u.Spec.MemberOf),
		entity.ID,
	)
	if err != nil {
		return model.User{}, fmt.Errorf("failed to update user: %w", err)
	}

	if err = tx.Commit(); err != nil {
		return model.User{}, fmt.Errorf("failed to commit transaction for update: %w", err)
	}
	return ru, nil
}

func (s sqliteStore) DeleteUser(ref model.EntityRef) (model.User, error) {
	user, err := s.ReadUser(ref)
	if err != nil {
		return model.User{}, err
	}

	err = deleteEntity(user.Entity.ID, s.db)
	if err != nil {
		return model.User{}, fmt.Errorf("failed to delete user: %w", err)
	}

	return user, nil
}

// ---

func (s sqliteStore) CreateGroup(g model.Group) (rg model.Group, err error) {
	var tx *sqlx.Tx
	tx, err = s.db.Beginx()
	if err != nil {
		return model.Group{}, fmt.Errorf("failed to begin transaction for create: %w", err)
	}
	defer func() {
		if err != nil {
			rerr := tx.Rollback()
			if rerr != nil {
				err = fmt.Errorf("failed to rollback transaction (%s): %w", rerr, err)
			}
		}
	}()

	entity, err := createEntity(g.Entity, tx)
	if err != nil {
		return model.Group{}, err
	}

	rg = g
	rg.Entity.ID = entity.ID

	_, err = tx.Exec(
		groupInsertStatement,
		entity.ID,
		g.Spec.Type,
		g.Spec.Profile.DisplayName,
		g.Spec.Profile.Email,
		g.Spec.Profile.Picture,
		g.Spec.Parent,
		model.MakeEntityRefs(g.Spec.Children),
		model.MakeEntityRefs(g.Spec.Members),
	)
	if err != nil {
		return model.Group{}, fmt.Errorf("failed to create group: %w", err)
	}

	if err = tx.Commit(); err != nil {
		return model.Group{}, fmt.Errorf("failed to commit transaction for create: %w", err)
	}
	return rg, nil
}

func (s sqliteStore) ReadGroup(ref model.EntityRef) (g model.Group, err error) {
	var tx *sqlx.Tx
	tx, err = s.db.Beginx()
	if err != nil {
		return model.Group{}, fmt.Errorf("failed to begin transaction for read: %w", err)
	}
	defer func() {
		if err != nil {
			rerr := tx.Rollback()
			if rerr != nil {
				err = fmt.Errorf("failed to rollback transaction (%s): %w", rerr, err)
			}
		}
	}()

	entity, err := readEntity(ref, tx)
	if err != nil {
		return model.Group{}, err
	}

	g = model.Group{
		Entity: entity,
	}

	rows, err := tx.Queryx(groupSelectStatement, entity.ID)
	if err != nil {
		return model.Group{}, fmt.Errorf("failed to query for group: %w", err)
	}
	defer rows.Close()
	if rows.Next() {
		var groupType string
		var displayName string
		var email string
		var picture string
		var parent model.EntityRef
		var children model.EntityRefs
		var members model.EntityRefs
		err = rows.Scan(&groupType, &displayName, &email, &picture, &parent, &children, &members)
		if err != nil {
			return model.Group{}, fmt.Errorf("failed to scan columns for group: %w", err)
		}
		g.Spec = model.GroupSpec{
			Type: groupType,
			Profile: model.GroupProfile{
				DisplayName: displayName,
				Email:       email,
				Picture:     picture,
			},
			Parent:   parent,
			Children: children.Items(),
			Members:  members.Items(),
		}
	}

	if err = tx.Commit(); err != nil {
		return model.Group{}, fmt.Errorf("failed to commit transaction for read: %w", err)
	}
	return g, nil
}

func (s sqliteStore) UpdateGroup(g model.Group) (rg model.Group, err error) {
	var tx *sqlx.Tx
	tx, err = s.db.Beginx()
	if err != nil {
		return model.Group{}, fmt.Errorf("failed to begin transaction for update: %w", err)
	}
	defer func() {
		if err != nil {
			rerr := tx.Rollback()
			if rerr != nil {
				err = fmt.Errorf("failed to rollback transaction (%s): %w", rerr, err)
			}
		}
	}()

	entity, err := updateEntity(g.Entity, tx)
	if err != nil {
		return model.Group{}, err
	}

	rg = g
	rg.Entity.ID = entity.ID

	_, err = tx.Exec(
		groupUpdateStatement,
		g.Spec.Type,
		g.Spec.Profile.DisplayName,
		g.Spec.Profile.Email,
		g.Spec.Profile.Picture,
		g.Spec.Parent,
		model.MakeEntityRefs(g.Spec.Children),
		model.MakeEntityRefs(g.Spec.Members),
		entity.ID,
	)
	if err != nil {
		return model.Group{}, fmt.Errorf("failed to update group: %w", err)
	}

	if err = tx.Commit(); err != nil {
		return model.Group{}, fmt.Errorf("failed to commit transaction for update: %w", err)
	}
	return rg, nil
}

func (s sqliteStore) DeleteGroup(ref model.EntityRef) (model.Group, error) {
	group, err := s.ReadGroup(ref)
	if err != nil {
		return model.Group{}, err
	}

	err = deleteEntity(group.Entity.ID, s.db)
	if err != nil {
		return model.Group{}, fmt.Errorf("failed to delete group: %w", err)
	}

	return group, nil
}
