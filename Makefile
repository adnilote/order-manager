run:
	DOCKER_BUILDKIT=1 docker build --ssh default -t order-manager .
	DOCKER_IMAGE=order-manager docker-compose up
	
lint:
	DOCKER_BUILDKIT=1 docker build --ssh default -f Dockerfile-lint  -t order-manager-lint .
	docker run order-manager-lint

test:
	DOCKER_BUILDKIT=1 docker build --ssh default -f Dockerfile-src  -t order-manager-src .
	DOCKER_IMAGE=order-manager-src docker-compose -f docker-compose-test.yaml up

gen-proto:
	protoc -I=. -I=${GOPATH}/src \
		-I=${GOPATH}/src/github.com/gogo/protobuf/protobuf \
		--gogoslick_out=Mgoogle/protobuf/any.proto=github.com/gogo/protobuf/types,\
Mgoogle/protobuf/duration.proto=github.com/gogo/protobuf/types,\
Mgoogle/protobuf/struct.proto=github.com/gogo/protobuf/types,\
Mgoogle/protobuf/timestamp.proto=github.com/gogo/protobuf/types,\
Mgoogle/protobuf/wrappers.proto=github.com/gogo/protobuf/types,\
plugins=grpc:. \
		app/business/entities/proto/*.proto