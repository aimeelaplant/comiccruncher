# Docker app container name for the Go backend.
DOCKER_APP_CONTAINER_NAME = comiccruncher
# Docker container name for Redis.
DOCKER_REDIS_CONTAINER_NAME = comiccruncher_redis
# Docker container name for Postgres.
DOCKER_PG_CONTAINER_NAME = comiccruncher_postgres
# Command to run Docker.
DOCKER_RUN = docker-compose run --rm ${DOCKER_APP_CONTAINER_NAME}
DOCKER_EXEC = docker-compose exec ${DOCKER_APP_CONTAINER_NAME}
# Command to run docker with the exposed port.
DOCKER_RUN_WITH_PORTS = docker-compose run --service-ports --rm  ${DOCKER_APP_CONTAINER_NAME}
# Settings to cross-compile go binary so that it works on Linux amd64 systems.
DOCKER_RUN_XCOMPILE = docker-compose run -e GOOS=linux -e GOARCH=amd64 -e CGO_ENABLED=0 --rm ${DOCKER_APP_CONTAINER_NAME}
# The container for tests.
DOCKER_RUN_TEST = docker-compose -f docker-compose.yml -f docker-compose.test.yml run --rm ${DOCKER_APP_CONTAINER_NAME}

# Command to go run locally.
GO_RUN_LOCAL = GORACE="log_path=./" go run -race

# The path to the migrations bin.
MIGRATIONS_BIN = bin/migrations
# The path to the webapp bin.
WEBAPP_BIN = bin/webapp
# The path to the temp webapp bin.
WEBAPP_TMP_BIN = bin/webapp1
# The path to the cerebro bin.
CEREBRO_BIN = bin/cerebro
# The path to the comic bin.
COMIC_BIN = bin/comic

# Locaton of the migrations cmd.
MIGRATIONS_CMD = ./cmd/migrations/migrations.go
# Location of the cerebro cmd.
CEREBRO_CMD = ./cmd/cerebro/cerebro.go
# Location of the web cmd.
WEB_CMD = ./cmd/web/web.go
# Location of comic cmd
COMIC_CMD = ./cmd/comic/comic.go

# The username and location to the api server (that's also the tasks server for now).
LB_SERVER = aimee@206.189.188.214
API_SERVER1 = aimee@68.183.132.127

# Creates a .netrc file for access to private Github repository for cerebro.
.PHONY: netrc
netrc:
	rm -rf .netrc && echo "machine github.com\nlogin $(GITHUB_ACCESS_TOKEN)" > .netrc && chmod 600 .netrc

# Build the docker container.
.PHONY: up
docker-up:
	docker-compose up -d --build --force-recreate --remove-orphans

# stop the docker containers.
.PHONY: docker-stop
docker-stop:
	docker-compose stop

# Show the docker logs from the services.
.PHONY: docker-logs
docker-logs:
	docker-compose -f docker-compose.yml logs --tail="100" -f postgres redis comiccruncher

# Run the migrations for the test db.
.PHONY: docker-migrations-test
docker-migrations-test:
	${DOCKER_RUN_TEST} go run ${MIGRATIONS_CMD}

# Create the containers for testing.
.PHONY: docker-up-test
docker-up-test:
	docker-compose -f docker-compose.yml -f docker-compose.test.yml up -d --build

# Remove the test containers.
.PHONY: docker-rm-test
docker-rm-test:
	docker-compose -f docker-compose.test.yml rm

# Stop the test containers.
.PHONY: docker-stop-test
docker-stop-test:
	docker-compose -f docker-compose.test.yml stop

# Run the go tests in the docker container.
.PHONY: docker-test
docker-test:
	${DOCKER_RUN_TEST} go test -v $(shell ${DOCKER_RUN_TEST} go list ./... | grep -v ./cmd) -coverprofile=coverage.txt

# Install the docker images and Go dependencies.
.PHONY: docker-install
docker-install: docker-up docker-mod-download

.PHONY:
mod-download:
	go mod download

.PHONY: docker-mod
docker-mod-download:
	${DOCKER_RUN} make mod-download

# Format the files with `go fmt`.
.PHONY: docker-format
docker-format:
	${DOCKER_RUN} make format

# Format the files
.PHONY: format
format:
	go fmt $(shell go list ./...)

# Vet the files in the Docker container.
.PHONY: docker-vet
docker-vet:
	${DOCKER_RUN} make vet

# Test the files with any race conditions (unfortunately Alpine-based images don't work w/ race command...so
# use this command locally :(
.PHONY: test
test:
	go test -race -v $(shell go list ./... | grep -v ./cmd) -coverprofile=coverage.txt

# Vet the files.
.PHONY: vet
vet:
	go vet $(shell go list ./...)

# Lint the go files.
.PHONY: lint
lint:
	golint $(shell go list ./...)

# Lint the files in the Docker container.
# Not sure why I have to specify `/gocode/bin/golint` and not just `golint`?!?!
.PHONY: docker-lint
docker-lint:
	${DOCKER_RUN} /gocode/bin/golint $(shell go list ./...)

# Reports any cyclomatic complexilities over 15. For goreportcard.
.PHONY: cyclo
cyclo:
	gocyclo -over 15 $(shell ls -d */ | grep -v vendor | awk '{print $$$11}')

# Reports any ineffectual if assignments. For goreportcard.
.PHONY: ineffassign
ineffassign:
	ineffassign .

# Reports any misspellings. For goreportcard.
.PHONY: misspell
misspell:
	misspell $(shell go list ./...)

# Generate any errors for go report card.
.PHONY: reportcard
reportcard: ineffassign misspell lint vet cyclo

# Run the Docker redis-cli.
.PHONY: redis-cli
docker-redis-cli:
	docker exec -it ${DOCKER_REDIS_CONTAINER_NAME} redis-cli -p 6380 -a foo

# Flush the redis cache.
.PHONY: redis-flush
docker-redis-flush:
	docker exec -it ${DOCKER_REDIS_CONTAINER_NAME} redis-cli -p 6380 -a foo flushall

# Build the web application in the Docker container with cross compilation settings so it works on linux amd64 systems.
.PHONY: docker-build-webapp-xcompile
docker-build-webapp-xcompile:
	${DOCKER_RUN_XCOMPILE} make build-webapp

# Builds the webapp binary.
.PHONY: build-webapp
build-webapp:
	 go build -o ./build/deploy/api/bin/webapp ${WEB_CMD}

docker-build-webapp:
	docker build -f ./build/deploy/api/Dockerfile . -t comiccruncher/api:latest

docker-docker-push-webapp:
	docker push comiccruncher/api:latest

docker-build-tasks:
	docker build -f ./build/deploy/tasks/Dockerfile . -t comiccruncher/tasks:latest

docker-push-tasks:
	docker push ${DOCKER_REPO}comiccruncher/tasks:latest

docker-run-migrations:
	docker run --env-file=.env comiccruncher/tasks:latest migrations

docker-run-cerebro:
	docker run --env-file=.env comiccruncher/tasks:latest cerebro ${F}

docker-run-comic:
	docker run --env-file=.env comiccruncher/tasks:latest comic ${F}

# Run the web application.
.PHONY: web
web:
	go run -race ${WEB_CMD} start -p 8001

# Run the web application in Docker container.
.PHONY: docker-web
docker-web:
	${DOCKER_RUN_WITH_PORTS} make web

# Docker run the migrations for the development database.
.PHONY: docker-migrations
docker-migrations:
	${DOCKER_RUN} go run ${MIGRATIONS_CMD}

# Run the migrations for the development database.
.PHONY: migrations
migrations:
	${GO_RUN_LOCAL} ${MIGRATIONS_CMD}

.PHONY: import-characterissues
import-characterissues:
	${GO_RUN_LOCAL} ${CEREBRO_CMD} import characterissues ${EXTRA_FLAGS}

.PHONY: import-charactersources
import-charactersources:
	${GO_RUN_LOCAL} ${CEREBRO_CMD} import charactersources ${EXTRA_FLAGS}

.PHONY: import-charactersources
start-characterissues:
	${GO_RUN_LOCAL} ${CEREBRO_CMD} start characterissues ${EXTRA_FLAGS}

# Runs the program for creating characters from the Marvel API.
.PHONY: import-characters
import-characters:
	 ${GO_RUN_LOCAL} ${CEREBRO_CMD} import characters ${EXTRA_FLAGS}

# Runs the program for generating thumbnails for characters.
.PHONY: import-characters
generate-thumbs:
	 ${GO_RUN_LOCAL} ${COMIC_CMD} generate thumbs ${EXTRA_FLAGS}

.PHONY: docker-import-characters
docker-import-characters:
	${DOCKER_RUN} go run ${CEREBRO_CMD} import characters

.PHONY: docker-generate-thumbs
docker-generate-thumbs:
	${DOCKER_RUN} go run ${COMIC_CMD} generate thumbs ${EXTRA_FLAGS}

# Generate mocks for testing.
.PHONY: mockgen
mockgen:
	mockgen -destination=internal/mocks/comic/sync.go -source=comic/sync.go
	mockgen -destination=internal/mocks/comic/repositories.go -source=comic/repositories.go
	mockgen -destination=internal/mocks/comic/services.go -source=comic/services.go
	mockgen -destination=internal/mocks/comic/cache.go -source=comic/cache.go
	mockgen -destination=internal/mocks/cerebro/characterissue.go -source=cerebro/characterissue.go
	mockgen -destination=internal/mocks/search/service.go -source=search/service.go
	mockgen -destination=internal/mocks/storage/s3.go -source=storage/s3.go
	mockgen -destination=internal/mocks/cerebro/utils.go -source=cerebro/utils.go
	mockgen -destination=internal/mocks/imaging/thumbnail.go -source=imaging/thumbnail.go
	mockgen -destination=internal/mocks/auth/auth.go -source=auth/auth.go

# Generate mocks for testing.
docker-mockgen:
	${DOCKER_RUN} make mockgen

# Builds the migrations binary.
.PHONY: build-migrations
build-migrations:
	go build -o ./bin/migrations -v ${MIGRATIONS_CMD}

# Builds the cerebro binary.
.PHONY: build-cerebro
build-cerebro:
	go build -o ./bin/cerebro -v ${CEREBRO_CMD}

# Builds the comic commands.
.PHONY: build-comic
build-comic:
	go build -o ./bin/comic -v ${COMIC_CMD}

.PHONY: remote-upload-deployfiles
remote-upload-deployfiles:
	scp -r .env ${LB_SERVER}:~/.
	scp -r ./build/deploy/nginx/* ${LB_SERVER}:~/.
	scp -r .env ${API_SERVER1}:~/.
	scp -r ./build/deploy/api/* ${API_SERVER1}:~/.

.PHONY: remote-deploy
remote-deploy: remote-upload-deployfiles
	ssh ${API_SERVER1} "sh deploy.sh"
	ssh ${LB_SERVER} "sh deploy.sh"
