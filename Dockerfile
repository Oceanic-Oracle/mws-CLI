FROM golang:alpine AS build

WORKDIR /app

COPY go.mod go.sum ./
RUN go mod download

COPY . .

RUN go build -o /app/mws .

FROM alpine:3.21

COPY --from=build /app/mws /usr/local/bin/mws

RUN chmod +x /usr/local/bin/mws

WORKDIR /usr/local/bin/

ENTRYPOINT ["mws"]