build:
	rm "./bin" -rf
	go build -o "./bin/Program.exe" "./src/main"
	cp "./src/main/config.json" "./bin/config.json"

get:
	go get github.com/gorilla/mux
	go get github.com/mongodb/mongo-go-driver/mongo