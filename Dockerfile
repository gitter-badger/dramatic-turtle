FROM golang:latest as builder
RUN mkdir -p /application
COPY . /application
WORKDIR /application
RUN make docker-build


FROM alpine:3.8
LABEL maintainer="https://github.com/HeikoAlexanderWeber"
RUN apk --no-cache add ca-certificates
WORKDIR /root/
COPY --from=builder /application .

ENTRYPOINT [ "./bin/program" ]
