## build: builds the binary
.PHONY: build
build:
	cd ./src && go build -o ../bin/go-mongo-application

clean:
	-rm ./bin/go-mongo-application