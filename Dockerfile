FROM golang:1.22 AS build-stage

WORKDIR /app
COPY ./main.go /app/
RUN CGO_ENABLED=0 go build main.go

FROM alpine:latest

WORKDIR /root/

COPY --from=build-stage /app/ .

ENTRYPOINT [ "./main" ]