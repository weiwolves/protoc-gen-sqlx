GOPATH ?= $(HOME)/go
SRCPATH := $(patsubst %/,%,$(GOPATH))/src
MODIFY=Mgoogle/protobuf/timestamp.proto=github.com/gogo/protobuf/types

default: build install

GO_MODULE=on

get:
	@GO111MODULE=${GO_MODULE} go get ./...

build:
	@protoc -I. -I$(SRCPATH) \
		--gogo_out="Mgoogle/protobuf/descriptor.proto=github.com/gogo/protobuf/protoc-gen-gogo/descriptor:${SRCPATH}" \
		pb/**/*.proto
	@#protoc  -I. -I$(SRCPATH) \
		--gogo_out="Mgoogle/protobuf/descriptor.proto=github.com/gogo/protobuf/protoc-gen-gogo/descriptor:${SRCPATH}" \
		github.com/gogo/protobuf/protoc-gen-gogo/*.proto

#	protoc -I. -I$(SRCPATH) -I./vendor \
#.PHONY: types
#types:
#	protoc --go_out=. types/types.proto
#

publish:
	@#./bumper.sh
	@git add -A
	@git commit -am "Bump version to v$(shell cat RELEASE)"
	@git tag v$(shell cat RELEASE)
	@git push
	@git push --tags

install:
	@GO111MODULE=${GO_MODULE} go install

clean:
	rm -f ./example/*.go

example: clean default
	@echo "Generate example"
	@protoc -I. -I$(SRCPATH) \
		--gofast_out=${MODIFY},plugins=grpc:. \
		--sqlx_out=. \
		--proto_path=${GOPATH}/src/github.com/gogo/protobuf/protobuf \
		example/*.proto
#	protoc -I. -I$(SRCPATH) -I./vendor \
#		--go_out="plugins=grpc:$(SRCPATH)" --gorm_out="$(SRCPATH)" \
#		example/feature_demo/test.proto example/feature_demo/test2.proto
#
#	protoc -I. -I$(SRCPATH) -I./vendor \
#		-I$(SRCPATH)/github.com/google/protobuf/src/ \
#		-I$(SRCPATH)/github.com/grpc-ecosystem/grpc-gateway \
#		-I$(SRCPATH)/github.com/grpc-ecosystem/grpc-gateway/third_party/googleapis \
#		--go_out="plugins=grpc:$(SRCPATH)" --gorm_out="$(SRCPATH)" \
#		example/contacts/contacts.proto
#
#test: example
#	go test ./...
#	go build ./example/contacts
#	go build ./example/feature_demo
