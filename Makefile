ifeq ($(OS),Windows_NT)
	HOST_OS := windows
else
    UNAME_S := $(shell uname -s)
    ifeq ($(UNAME_S),Linux)
        HOST_OS := linux
    endif
    ifeq ($(UNAME_S),Darwin)
        HOST_OS := darwin
    endif
endif

ifeq ($(HOST_OS),windows)
	PROGRAM := Program.exe
else
	PROGRAM := program
endif


init:
	$(MAKE) get
	$(MAKE) build-dev

build-dev:
	GOOS=$(HOST_OS) go build -a -installsuffix cgo -o "./bin/$(PROGRAM)" "./src/main/program.go"
	cp "./src/main/config.json" "./bin/config.json"

build-prod:
	CGO_ENABLED=0 GOOS=$(HOST_OS) go build -a -installsuffix cgo -o "./bin/$(PROGRAM)" "./src/main/program.go"
	cp "./src/main/config.json" "./bin/config.json"

run-prod:
	$(MAKE) build-dev
	./bin/$(PROGRAM)

clean:
	rm "./bin" -rf

get:
	go get github.com/gorilla/mux
	go get github.com/mongodb/mongo-go-driver/mongo


docker-build-image:
	$(MAKE) clean
	docker build -t dramatic-turtle .

docker-run-image:
	docker run -d -p 8080:8080 -t dramatic-turtle

docker-build-project:
	$(MAKE) get
	$(MAKE) build-prod