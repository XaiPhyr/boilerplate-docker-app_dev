# Docker boilerplate App Development

uses [SQL Migrate](https://github.com/rubenv/sql-migrate)

## Docker
- to run application type command `docker compose up --watch` to keep track of changes and will sync and restart running docker

### Secrets
password is stored in `secrets` folder  
to initialize password, rename template.env  
to `.env` and store your password  
as for backend services, remove template on filename

## Services

### app_service_0
implemented redis for caching and SSE

### app_service_1
subscribing to app_service_0 via redis

## Frontend
- Framework: Vite + Vue.js
- CSS: TailwindCSS