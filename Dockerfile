FROM golang:alpine
LABEL maintainer="sergiogc5@correo.ugr.es"
LABEL url="https://github.com/Olasergiolas/Go-AutoEQ"

RUN apk add --no-cache g++

RUN addgroup -S goautoeq && adduser -S goautoeq -G goautoeq
USER goautoeq

WORKDIR /app/test
COPY go.mod ./

RUN go install github.com/go-task/task/v3/cmd/task@latest
RUN go mod download

ENTRYPOINT ["/go/bin/task", "test"]