clean:
	rm "./bin" -rf

build:
	$(MAKE) clean
	go build -o "./bin/Program.exe" "./src/main"
	cp "./src/main/config.json" "./bin/config.json"

get:
	go get github.com/gorilla/mux
	go get github.com/mongodb/mongo-go-driver/mongo