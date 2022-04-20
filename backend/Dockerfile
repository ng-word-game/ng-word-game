FROM golang:1.18.1-bullseye

RUN curl -sSfL https://raw.githubusercontent.com/cosmtrek/air/master/install.sh | sh -s -- -b $(go env GOPATH)/bin

WORKDIR /app

COPY go.mod go.sum /app/
RUN go mod download

COPY . /app

RUN chmod 755 ./entrypoint.sh
ENTRYPOINT [ "./entrypoint.sh" ]