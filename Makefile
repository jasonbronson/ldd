DIST_DIR = ./dist
TIMESTAMP ?= $(shell date -u +'%Y-%m-%dT%H:%M:%SZ')
SHA ?= $(shell git rev-parse --short HEAD)

DATABASE_URL := postgresql://admin:admin@postgres:5432/ldd_db?sslmode=disable

local:
	docker-compose up

test:
	docker exec -it -e ENVIRONMENT=test promoengine-api go test ./controller -v

build: ## Build go binary
	CGO_ENABLED=0 go build -ldflags "-s -w -X main.buildtime=$(TIMESTAMP) -X main.commitsha=$(SHA)" -o ${DIST_DIR}/api ./cmd/api/main.go

buildcron: ## Build go binary
	CGO_ENABLED=0 go build -ldflags "-s -w -X main.buildtime=$(TIMESTAMP) -X main.commitsha=$(SHA)" -o ${DIST_DIR}/cron ./cmd/cron/main.go

buildmigrate: 
	CGO_ENABLED=0 go build -ldflags "-s -w -X main.buildtime=$(TIMESTAMP) -X main.commitsha=$(SHA)" -o ${DIST_DIR}/migrations ./cmd/migrations/main.go