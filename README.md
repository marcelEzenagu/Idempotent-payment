# Idempotent-payments API - Go

Production-grade idempotent payments service built with Go, Redis, PostgreSQL.

**Why idempotency?** Prevents double charges on network retries. Key stored in Redis with Lua for atomicity.

**Stack:** Go, Gin, Redis (Lua), Postgres, Docker

**Features:**
- Idempotency-Key handling
- Atomic Redis check-and-set (Lua script)https://github.com/marcelEzenagu/Idempotent-payment/tree/main
- Postgres transactional outbox
- Docker Compose (api + redis + postgres)

Run:
docker-compose up --build
curl -X POST localhost:8080/pay -H "Idempotency-Key: 123"
