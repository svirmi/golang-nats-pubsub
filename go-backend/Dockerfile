FROM golang:1.21.6-alpine3.19 AS golang

WORKDIR /app
COPY . .

RUN go mod download
RUN go mod verify

RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o /stream .

FROM gcr.io/distroless/static-debian11

COPY --from=golang /stream .
COPY --from=golang app/config.yaml .

# EXPOSE 8080

CMD ["/stream"]