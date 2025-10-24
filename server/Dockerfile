FROM golang:1.24-alpine AS builder

WORKDIR /app

COPY go.mod go.sum ./
RUN go mod download

COPY . .

RUN go install goa.design/goa/v3/cmd/goa@v3
RUN goa gen app/server/design

RUN go build -o /server ./main.go

FROM alpine:latest

COPY --from=builder /server /server

EXPOSE 8080

ENTRYPOINT ["/server"]
