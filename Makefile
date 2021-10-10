build-proto:
	@echo " >> building debt protobuf"
	go get -u github.com/golang/protobuf/protoc-gen-go
	@protoc -I protos/ protos/*.proto --go_out=plugins=grpc:protos