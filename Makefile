include .env
export

start:
	@docker-compose up -d

stop:
	@docker-compose down --rmi local -v

migrate-up:
	go get -v github.com/rubenv/sql-migrate/...
	sql-migrate up

migrate-status:
	sql-migrate status

migrate-down:
	sql-migrate down

migrate-redo:
	sql-migrate redo

create:
	@read -p  "What is the name of migration : " NAME; \
	sql-migrate new $$NAME

run-rest:
	@echo "===== RUNNING SERVICE REST ====="
	@go mod vendor
	@go build -o ./bin/rest ./app/rest/app.go && ./bin/rest

run-grpc:
	@echo "===== RUNNING SERVICE GRPC ====="
	@go mod vendor
	@go build -o ./bin/grpc ./app/grpc/app.go && ./bin/grpc

build-proto:
	@echo " >> building debt protobuf"
	go get -u github.com/golang/protobuf/protoc-gen-go
	@protoc -I protos/ protos/*.proto --go_out=plugins=grpc:protos

mocks:
	@mockery --disable-version-string --all --dir ./domain/repository/ --output ./domain/repository/mocks/
	@mockery --disable-version-string --all --dir ./domain/definition/ --output ./domain/definition/mocks/

test:
	@bash ./script/go_test.sh