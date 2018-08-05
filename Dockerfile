FROM golang:latest as builder
RUN mkdir -p /application
COPY . /application
WORKDIR /application
RUN make docker-build


FROM alpine:latest
RUN apk --no-cache add ca-certificates
WORKDIR /root/
COPY --from=builder /application .

ENTRYPOINT [ "./bin/program" ]
