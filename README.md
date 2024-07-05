# Docker boilerplate App Development

uses 
### Packages
- [SQL Migrate](https://github.com/rubenv/sql-migrate) - SQL Schema migration tool for Go. Based on gorp and goose.
- [go-chi](https://github.com/go-chi/chi) - lighweight router for building Go HTTP services.
- [bun](https://github.com/uptrace/bun) - Lightweight Golang ORM (using PostgreSQL).
- [jwt](https://github.com/golang-jwt/jwt) - Go implementation of JSON Web Token Authentication.
- [yaml](https://gopkg.in/yaml.v3) - enables to decode YAML files to used for configuration.
- [gomail](https://gopkg.in/gomail.v2) - Golang SMTP mailer.
- [websocket](https://github.com/gorilla/websocket) - used for notifications and chat <span style="color: red">(*in progress*)</span>.

## Docker
- to run application type command `docker compose up --watch` to keep track of changes and will sync and restart running docker

### Secrets
password is stored in `secrets` folder  
to initialize secrets, remove the word `template` on files inside secrets  

## Services
Any request made by the CLIENT will pass through GATEWAY then reaches to SERVICES  
where token will be authenticated to proceed

### folder structure
```
+---db
+---secrets
|   +---backend
|   +---db
|   \---frontend
\---src
    +---backend
    |   +---gateway
    |   |   +---conf
    |   |   +---controllers
    |   |   +---db
    |   |   |   \---sql
    |   |   +---mail
    |   |   +---middlewares
    |   |   +---models
    |   |   +---routers
    |   |   +---services
    |   |   +---templates
    |   |   |   \---emails
    |   |   \---utils
    |   \---service_1
    |       +---conf
    |       +---controllers
    |       +---middlewares
    |       +---models
    |       +---routers
    |       +---services
    |       \---utils
    \---frontend
        \---app
            +---public
            \---src
                +---api
                +---assets
                +---components
                |   +---cards
                |   +---draggables
                |   +---paginations
                |   \---tables
                +---pages
                |   +---dashboard
                |   \---users
                \---router
```

### gateway
API Gateway connecting to Services  
Authentication token stored to `REDIS` with key structure `ID:USERNAME`

### api_service_1
API Service 1

## Frontend
- Framework: Vite + Vue.js
- CSS: TailwindCSS

##
[![GoDoc Widget]][GoDoc]
[![GitHub Widget]][My GitHub]

[GoDoc]: https://pkg.go.dev/github.com/go-chi/chi/v5
[GoDoc Widget]: https://img.shields.io/badge/references-go?style=flat&logo=go&labelColor=%235C5C5C&color=%23007D9C
[My GitHub]: https://github.com/XaiPhyr
[GitHub Widget]: https://img.shields.io/badge/XaiPhyr-github?style=flat&logo=github&labelColor=%235C5C5C&color=%235C5C5C
