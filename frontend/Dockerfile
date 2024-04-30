FROM golang:1.21.6-alpine3.19 AS golang

WORKDIR /app
COPY . .

RUN go mod download
RUN go mod verify

RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o /web .

FROM gcr.io/distroless/static-debian11

COPY --from=golang /web .
COPY --from=golang app/index.html .

EXPOSE 8080

CMD ["/web"]