# Decisions — Nucleus API

- Language: Go
- Router: chi (or Gin)
- DB: Postgres 16 + pgx v5 + SQLC
- Migrations: golang-migrate
- Cache/limits: Redis (Upstash)
- Auth: JWT access (15m) + rotating refresh; argon2id
- Contracts: OpenAPI 3.1
- Deploy: Render (Web Service)

Non-goals (MVP): no GraphQL, no microservices, no email/SSO.

