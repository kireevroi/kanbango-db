FROM golang:1.20.3-alpine3.16 as build

WORKDIR /usr/src/app
COPY go.mod go.sum ./
RUN go mod tidy
COPY . .
RUN go build -o /kanbango-auth ./cmd/main.go

FROM alpine:3.16
COPY --from=build /kanbango-auth /kanbango-auth
COPY --from=build /usr/src/app/.env /.env

CMD ["/kanbango-auth"]