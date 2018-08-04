dev-build:
	go build -o "./bin/Program.exe" "./src/main/program.go"
	cp "./src/main/config.json" "./bin/config.json"

dev-run:
	$(MAKE) dev-build
	./bin/Program.exe

clean:
	rm "./bin" -rf

get:
	go get github.com/gorilla/mux
	go get github.com/mongodb/mongo-go-driver/mongo

build-docker:
	$(MAKE) clean
	docker build -t dramatic-turtle .

run-docker:
	docker run -d -p 8080:8080 -t dramatic-turtle

docker-build:
	$(MAKE) get
	go build -o "./bin/program" "./src/main/program.go"
	cp "./src/main/config.json" "./bin/config.json"
