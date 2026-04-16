# Peakly

A todo app that doesn't lie to you.

Most task managers let you pile on tasks without asking if you're actually capable of doing them. Peakly uses a productivity model built on circadian rhythm, sleep debt, and motivation to schedule tasks when you're genuinely able to do them — not just when they're due.

## The idea

Traditional todo apps treat every hour as equal. They don't know that you're useless at 3pm, that you slept 5 hours, or that you've been working for 6 hours straight. Peakly does. It models your energy across the day and assigns tasks to windows where your capacity matches the task's cognitive demand.

## Stack

| Layer | Technology |
|---|---|
| Frontend | Nuxt 4, Vue 3, TypeScript, Tailwind CSS v4, shadcn-vue |
| Backend | Go 1.26 |
| Database | PostgreSQL 18 |
| Infrastructure | Docker, Docker Compose |

## Project structure

```
peakly/
├── frontend/          # Nuxt 4 app
├── backend/           # Go API
├── docker-compose.yml
├── .env.example
└── README.md
```

## Getting started

### Prerequisites

- Docker and Docker Compose

That's it. You don't need to install PostgreSQL, Go, or Node locally — everything runs inside Docker.

### Development

**1. Copy and fill in the environment file:**

```bash
cp .env.example .env
```

```env
PG_USER=peakly
PG_PASS=your_password
PG_DB=peakly_db
APP_ENV=development
API_BASE=http://localhost:8080
```

**2. Start all services:**

```bash
docker compose up
```

This starts PostgreSQL, the Go backend with hot reload, and the Nuxt frontend with HMR. Edit code locally — changes reflect immediately without rebuilding.

**3. Services:**

| Service | URL |
|---|---|
| Frontend | http://localhost:3000 |
| Backend API | http://localhost:8080 |
| PostgreSQL | localhost:5432 |

> The backend waits for PostgreSQL to pass its healthcheck before starting. On first run this takes a few seconds.

### Production

```bash
APP_ENV=production docker compose up -d
```

```bash
docker compose down      # stop and remove containers
docker compose logs -f   # follow logs
```

## How the model works

Peakly scores your available energy at any given time using three factors:

- **Circadian rhythm** — your natural alertness curve across the day modeled as a cosine function
- **Sleep debt** — cumulative deficit from your sleep target that compounds over days
- **Motivation** — a dynamic multiplier that decays with consecutive completed tasks and recovers with rest

Tasks are assigned a cognitive demand score. Peakly matches high-demand tasks to your peak windows and light tasks to your low-energy periods. The result is a schedule that reflects reality, not optimism.

## Status

Early development. MVP in progress.

## License

MIT