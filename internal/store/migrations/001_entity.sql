-- +migrate Up
CREATE TABLE entity (
  id INTEGER PRIMARY KEY,
  apiVersion VARCHAR(50) NOT NULL,
  kind VARCHAR(50) NOT NULL,
  name VARCHAR(255) NOT NULL,
  namespace VARCHAR(255) NOT NULL,
  title VARCHAR(255),
  description TEXT,
  tags TEXT
);
CREATE UNIQUE INDEX entity_ref_idx ON entity (kind, namespace, name);

CREATE TABLE label (
  id INTEGER PRIMARY KEY,
  entity_id INTEGER NOT NULL,
  k VARCHAR(255) NOT NULL,
  v VARCHAR(255) NOT NULL,
  CONSTRAINT fk_entity
    FOREIGN KEY (entity_id)
    REFERENCES entity(id)
    ON DELETE CASCADE
);
CREATE UNIQUE INDEX label_idx ON label (entity_id, k);

CREATE TABLE annotation (
  id INTEGER PRIMARY KEY,
  entity_id INTEGER NOT NULL,
  k VARCHAR(255) NOT NULL,
  v TEXT NOT NULL,
  CONSTRAINT fk_entity
    FOREIGN KEY (entity_id)
    REFERENCES entity(id)
    ON DELETE CASCADE
);
CREATE UNIQUE INDEX annotation_idx ON annotation (entity_id, k);

CREATE TABLE link (
  id INTEGER PRIMARY KEY,
  entity_id INTEGER NOT NULL,
  url TEXT NOT NULL,
  title TEXT,
  icon TEXT,
  type VARCHAR(255),
  CONSTRAINT fk_entity
    FOREIGN KEY (entity_id)
    REFERENCES entity(id)
    ON DELETE CASCADE
);

CREATE TABLE component (
  id INTEGER PRIMARY KEY,
  entity_id INTEGER NOT NULL,
  type VARCHAR(255) NOT NULL,
  lifecycle VARCHAR(255) NOT NULL,
  owner VARCHAR(512) NOT NULL,
  system VARCHAR(512),
  subcomponent_of VARCHAR(512),
  provides_apis TEXT,
  consumes_apis TEXT,
  depends_on TEXT,
  dependency_of TEXT,
  CONSTRAINT fk_entity
    FOREIGN KEY (entity_id)
    REFERENCES entity(id)
    ON DELETE CASCADE
);

CREATE TABLE api (
  id INTEGER PRIMARY KEY,
  entity_id INTEGER NOT NULL,
  type VARCHAR(255) NOT NULL,
  lifecycle VARCHAR(255) NOT NULL,
  owner VARCHAR(512) NOT NULL,
  system VARCHAR(512),
  definition TEXT,
  CONSTRAINT fk_entity
    FOREIGN KEY (entity_id)
    REFERENCES entity(id)
    ON DELETE CASCADE
);

CREATE TABLE user (
  id INTEGER PRIMARY KEY,
  entity_id INTEGER NOT NULL,
  display_name VARCHAR(255),
  email VARCHAR(255),
  picture TEXT,
  member_of TEXT NOT NULL,
  CONSTRAINT fk_entity
    FOREIGN KEY (entity_id)
    REFERENCES entity(id)
    ON DELETE CASCADE
);

CREATE TABLE grp (
  id INTEGER PRIMARY KEY,
  entity_id INTEGER NOT NULL,
  type VARCHAR(50) NOT NULL,
  display_name VARCHAR(255),
  email VARCHAR(255),
  picture TEXT,
  parent VARCHAR(512),
  children TEXT,
  members TEXT,
  CONSTRAINT fk_entity
    FOREIGN KEY (entity_id)
    REFERENCES entity(id)
    ON DELETE CASCADE
);

-- +migrate Down

DROP TABLE grp;

DROP TABLE user;

DROP TABLE api;

DROP TABLE component;

DROP TABLE link;

DROP INDEX annotation_idx;
DROP TABLE annotation;

DROP INDEX label_idx;
DROP TABLE label;

DROP INDEX entity_ref_idx;
DROP TABLE entity;
