FROM golang:1.19-alpine as builder

WORKDIR /app

COPY . .

RUN apk update \
    && apk add --no-cache git curl make gcc g++ \
    && go mod tidy

RUN go get -u github.com/cosmtrek/air && \
    go build -o /go/bin/air github.com/cosmtrek/air

EXPOSE 8080

CMD ["air", "-c", ".air.toml"]

#CMD ["go", "run", "main.go"]