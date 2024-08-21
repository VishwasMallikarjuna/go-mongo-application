## build: builds the binary
SRCS:=$(find src -name *.go)
TEST:=./src/testCoverage.out

.PHONY: build
build:
	cd ./src && go build -o ../bin/go-mongo-application

test: clean $(TEST)

# '-tags tests' is used to excluded the multiple main declarations from test builds
$(TEST): $(SRCS) src/go.mod src/go.sum
	cd src; go test -coverprofile testCoverage.out ./... -v -tags tests

clean:
	-rm ./bin/go-mongo-application
	-rm ./src/testCoverage.out

.PHONY: docker_up
docker_up:
	@docker compose up --build