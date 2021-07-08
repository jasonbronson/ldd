DIST_DIR = ./dist
TIMESTAMP ?= $(shell date -u +'%Y-%m-%dT%H:%M:%SZ')
SHA ?= $(shell git rev-parse --short HEAD)

DATABASE_URL := postgresql://postgres:admin@192.168.1.99:5432/ldd?sslmode=disable

local:
	docker-compose up

test:
	docker exec -it -e ENVIRONMENT=test promoengine-api go test ./controller -v

build: ## Build go binary
	go build -ldflags "-s -w -X main.buildtime=$(TIMESTAMP) -X main.commitsha=$(SHA)" -o ${DIST_DIR}/api ./cmd/api/main.go

buildcron: ## Build go binary
	go build -ldflags "-s -w -X main.buildtime=$(TIMESTAMP) -X main.commitsha=$(SHA)" -o ${DIST_DIR}/cron ./cmd/cron/main.go

buildmigrate: 
	go build -ldflags "-s -w -X main.buildtime=$(TIMESTAMP) -X main.commitsha=$(SHA)" -o ${DIST_DIR}/migrations ./cmd/migrations/main.go