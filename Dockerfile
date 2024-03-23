FROM golang:1.22.1-alpine3.19

RUN apk update && apk upgrade && apk add curl \
                          bash \
                          make \
                         busybox-extras  && \
     rm -rf /var/cache/apk/*

WORKDIR /app

COPY go.mod ./
COPY go.sum ./

RUN go mod download
RUN go install github.com/swaggo/swag/cmd/swag@latest
COPY . .

RUN make build

EXPOSE 3000

ENTRYPOINT ["./main"]
