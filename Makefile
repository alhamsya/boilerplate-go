include .env
export

start:
	@go mod vendor
	@bash ./script/setup.sh
	@docker-compose up -d

#start-and-migrate:
#	@docker-compose up -d
#	@docker-compose exec database sh -c 'until mysqladmin ping -P 3306 -palhamsya -uroot | grep "mysqld is alive"; do { echo "MySQL is unavailable - waiting for it... ?"; sleep 1; }; done'
#	@go get -v github.com/rubenv/sql-migrate/...
#	@sql-migrate up

stop:
	@docker-compose down --rmi local -v

migrate-up:
	@go get -d -v github.com/rubenv/sql-migrate/...@f842348935589e4563be545226d465178bd439cf
	@sql-migrate up

migrate-status:
	@sql-migrate status -config=dbconfig.yml -env="development"

migrate-down:
	@sql-migrate down

migrate-redo:
	@sql-migrate redo

create:
	@read -p  "What is the name of migration : " NAME; \
	sql-migrate new $$NAME

run-rest:
	@printf "\033[0;30m\033[42m === RUNNING SERVICE REST === \033[0m\n"
	@go mod vendor
	@go build -o ./bin/rest ./app/rest/app.go && ./bin/rest

run-grpc:
	@printf "\033[0;30m\033[42m === RUNNING SERVICE GRPC === \033[0m\n"
	@go mod vendor
	@go build -o ./bin/grpc ./app/grpc/app.go && ./bin/grpc

run-cron:
	@printf "\033[0;30m\033[42m === RUNNING SERVICE CRON === \033[0m\n"
	@go mod vendor
	@go build -o ./bin/cron ./app/cron/app.go && ./bin/cron

run-consumer:
	@printf "\033[0;30m\033[42m === RUNNING SERVICE CONSUMER === \033[0m\n"
	@go mod vendor
	@go build -o ./bin/consumer ./app/consumer/app.go && ./bin/consumer

build-proto:
	@printf "\033[0;30m\033[42m === BUILDING PROTOBUF === \033[0m\n"
	@rm -rf ./protos/*.pb.go
	@protoc -I protos/ protos/*.proto --go_out=plugins=grpc:protos

mocks:
	@printf "\033[0;30m\033[42m === GENERATE MOCKS === \033[0m\n"
	@rm -rf ./domain/repositorys/mocks/
	@mockery --disable-version-string --all --dir ./domain/repositorys/ --output ./domain/repositorys/mocks/

test:
	@bash ./script/test.sh

build:
	@printf "\033[0;30m\033[42m === GENERATE BUILD === \033[0m\n"
	@GOOS=linux GOARCH=amd64
	@go build -o ./bin/REST ./app/rest
	@go build -o ./bin/GRPC ./app/grpc
	@go build -o ./bin/CRON ./app/cron
	@go build -o ./bin/CONSUMER ./app/consumer