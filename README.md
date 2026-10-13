# Idempotent Payment API — Prevents Double Charges

Production-grade payments service built with Go, Redis, PostgreSQL.

**Why idempotency?** Prevents double charges on network retries. Key stored in Redis with Lua for atomicity.

**Stack:** Go, Gin, Redis (Lua), Postgres, Docker

### How we get the Idempotency Key

We don't generate it on server — client must send it. This is how Stripe does it.

1. Client creates UUID v4 before pay: `const key = crypto.randomUUID()`
2. Sends it: `Idempotency-Key: a1b2c3d4-...`
3. Server Lua: if key exists return cached, else SET and process. Cached 24h with `X-Idempotent-Replayed: true`

### Features
- Idempotency-Key handling
- Atomic Redis check-and-set (Lua script)
- Postgres transactional outbox
- Docker Compose (api + redis + postgres)

### Run

Prereq: Docker

```bash
git clone https://github.com/marcelEzenagu/Idempotent-payment.git
cd Idempotent-payment
docker-compose up --build
