FROM golang:latest

RUN mkdir -p /application
COPY . /application
WORKDIR /application

RUN make docker-build

EXPOSE 8080

ENTRYPOINT [ "./bin/program" ]
