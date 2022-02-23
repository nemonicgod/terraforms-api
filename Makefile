DC=docker-compose
IG=terraforms-api:local

GO_SRC_DIRS := $(shell \
	find . -name "*.go" -not -path "./vendor/*" | \
	xargs -I {} dirname {}  | \
	uniq)

GO_TEST_DIRS := $(shell \
	find . -name "*_test.go" -not -path "./vendor/*" | \
	xargs -I {} dirname {}  | \
	uniq)

.PHONY: all build compose-down compose-up deps erd run test watch migrate-create tag-and-deploy-aws

b: build
c: clean
cd: compose-down
cu: compose-up
d: deps
r: run
w: watch

build: deps
	@echo "building [\033[31mapi\033[0m]..."
	@go build -o bin/api pkg/api/main.go
	@echo "building [\033[32mclient\033[0m]..."
	@go build -o bin/client pkg/client/main.go
	@echo "building [\033[35mworker\033[0m]..."
	@go build -o bin/worker pkg/worker/main.go

clean:
	@rm ./bin/* || true
	@docker system prune -f
	@docker volume rm $$(docker volume ls -qf dangling=true)

compose-down: clean

compose-up:
	@$(DC) -f $(DC).yml up --build

compose:
	@$(DC) -f $(DC).yml up postgres redis api $(SERVICE)

deps:
	@go get -u ./...
	@go mod download
	@go mod tidy
	@go mod vendor

docker-build:
	docker build --squash -t $(IG) -f Dockerfile .

docker-run:
	docker run \
		--env ENVIRONMENT=local \
		--env POSTGRES_USER=postgres \
		--env POSTGRES_PASS="" \
		--env POSTGRES_HOST=localhost \
		--env POSTGRES_PORT=5432 \
		-it $(IG)

docker-shell:
	docker run -it --entrypoint /bin/sh $(IG)

erd:
	@go-erd -path ./infra/database/models | dot -T svg > erd.svg

inspect:
	@go tool cover -html=coverage.out

lint:
	@golint ./...

migrate-create:
	@migrate create -ext sql -dir infra/database/migrations -format unix $(FILENAME)

run:
	@PORT=8000 ENVIRONMENT=development ./bin/api

test:
	@go test -v ./...

test_richgo:
	@ENVIRONMENT=test richgo test -v ./... -cover -coverprofile=coverage.out

uml-create:
	@planter "postgres://postgres:password@localhost/postgres?sslmode=disable" -o database.uml

uml-compile:
	@java -jar  /usr/local/Cellar/plantuml/1.2021.9/libexec/plantuml.jar -verbose database.uml

uml-clean:
	rm database.uml || true && rm database.uml || true

uml-exec: uml-clean uml-create uml-compile