clean:
	rm "./bin" -rf

build:
	$(MAKE) clean
	go build -o "./bin/Program.exe" "./src/main/program.go"
	cp "./src/main/config.json" "./bin/config.json"

run:
	$(MAKE) build
	./bin/Program.exe

get:
	go get github.com/gorilla/mux
	go get github.com/mongodb/mongo-go-driver/mongo
