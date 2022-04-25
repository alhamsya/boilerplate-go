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
	@protoc -I protos/ protos/*.proto --go_out=plugins=grpc:protos

mocks:
	@printf "\033[0;30m\033[42m === GENERATE MOCKS === \033[0m\n"
	@mockery --disable-version-string --all --dir ./domain/repository/ --output ./domain/repository/mocks/
	@mockery --disable-version-string --all --dir ./domain/definition/ --output ./domain/definition/mocks/

test:
	@bash ./script/test.sh