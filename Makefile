include .env
LOCAL_BIN:=$(CURDIR)/bin

install-deps:
	GOBIN=$(LOCAL_BIN) go install github.com/pressly/goose/v3/cmd/goose@v3.19.2
	GOBIN=$(LOCAL_BIN) go install github.com/cosmtrek/air@latest

install-golangci-lint:
	GOBIN=$(LOCAL_BIN) go install github.com/golangci/golangci-lint/cmd/golangci-lint@latest

update-packages:
	go get -u ./...

run-air:
	air --build.cmd "go build -o .bin/air cmd/telegram_bot/main.go" --build.bin "./.bin/air"

lint:
	GOBIN=$(LOCAL_BIN) golangci-lint run ./... --config .golangci.yaml

local-migration-status:
	$(LOCAL_BIN)/goose -dir ${MIGRATION_DIR} postgres ${PG_DSN} status -v

local-migration-up:
	$(LOCAL_BIN)/goose -dir ${MIGRATION_DIR} postgres ${PG_DSN} up -v

local-migration-down:
	$(LOCAL_BIN)/goose -dir ${MIGRATION_DIR} postgres ${PG_DSN} down -v

docker-migration-status:
	$(LOCAL_BIN)/goose -dir ${MIGRATION_DIR} postgres ${PG_DOCKER_DSN} status -v

docker-migration-up:
	$(LOCAL_BIN)/goose -dir ${MIGRATION_DIR} postgres ${PG_DOCKER_DSN} up -v

docker-migration-down:
	$(LOCAL_BIN)/goose -dir ${MIGRATION_DIR} postgres ${PG_DOCKER_DSN} down -v

gen-cert:
	openssl genrsa -out cert/ca.key 4096
	openssl req -new -x509 -key cert/ca.key -sha256 -subj "/C=US/ST=NJ/O=CA, Inc." -days 365 -out cert/ca.cert
	openssl genrsa -out cert/service.key 4096
	openssl req -new -key cert/service.key -out cert/service.csr -config cert/certificate.conf
	openssl x509 -req -in cert/service.csr -CA cert/ca.cert -CAkey cert/ca.key -CAcreateserial \
    		-out cert/service.pem -days 365 -sha256 -extfile cert/certificate.conf -extensions req_ext

.PHONY: mock-generate
mock-generate:
	rm -rf internal/service/mocks
	mockgen -source=internal/service/service.go \
	-destination=internal/service/mocks/mock_service.go

	rm -rf internal/repository/mocks
	mockgen -source=internal/repository/repository.go \
    	-destination=internal/repository/mocks/mock_repository.go

.PHONY: test-coverage
test-coverage:
	go test -short -count=1 -coverprofile=coverage.out ./...
	go tool cover -html=coverage.out
	rm coverage.out