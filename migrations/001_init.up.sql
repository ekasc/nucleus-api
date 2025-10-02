CREATE EXTENSION IF NOT EXISTS pgcrypto;

BEGIN;

-- USERS

CREATE TABLE users (
	id UUID PRIMARY KEY DEFAULT gen_random_uuid(),
	email TEXT NOT NULL,
	password_hash TEXT NOT NULL,
	name TEXT NOT NULL,
	email_verified_at TIMESTAMPTZ,
	created_at TIMESTAMPTZ NOT NULL DEFAULT now(),
	updated_at TIMESTAMPTZ NOT NULL DEFAULT now()
);

CREATE UNIQUE INDEX users_email_lower_ux ON users (lower(email));

-- ORGS
CREATE TABLE orgs (
	id	UUID PRIMARY KEY DEFAULT gen_random_uuid(),
	name TEXT NOT NULL,
	slug TEXT NOT NULL UNIQUE,
	owner_user_id UUID NOT NULL REFERENCES users(id),
	created_at TIMESTAMPTZ NOT NULL DEFAULT now()
);


-- ORG MEMBERSHIPS
CREATE TABLE org_memberships (
	id UUID PRIMARY KEY DEFAULT gen_random_uuid(),
	org_id UUID NOT NULL REFERENCES orgs(id) ON DELETE CASCADE,
	user_id UUID NOT NULL REFERENCES users(id) ON DELETE CASCADE,
	role TEXT NOT NULL CHECK (role IN ('owner','admin','member','viewer')),
	created_at TIMESTAMPTZ NOT NULL DEFAULT now(),
	UNIQUE (org_id, user_id)

);

-- PROJECTS (tenant-scoped)
CREATE TABLE projects (
	id UUID PRIMARY KEY DEFAULT gen_random_uuid(),
	org_id UUID NOT NULL REFERENCES orgs(id) ON DELETE CASCADE,
	name TEXT NOT NULL,
	created_at TIMESTAMPTZ NOT NULL DEFAULT now()
);


CREATE UNIQUE INDEX projects_org_created_idx ON projects (org_id, created_at DESC);

-- TASKS (project-scoped)
CREATE TABLE tasks (
    id UUID PRIMARY KEY DEFAULT gen_random_uuid(),
    project_id UUID NOT NULL REFERENCES projects(id) ON DELETE CASCADE,
	title TEXT NOT NULL,
	body TEXT NOT NULL,
	status TEXT NOT NULL CHECK (status IN ('open', 'doing','done')),
    created_at TIMESTAMPTZ NOT NULL DEFAULT now()
);

CREATE INDEX tasks_project_created_idx ON tasks (project_id, created_at DESC);

COMMIT;

