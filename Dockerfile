FROM golang:1.21.6

LABEL maintainer="Juanlubel <juanluis.belda@gmail.com>"

WORKDIR /go/src/go_server


COPY go.mod go.sum ./
#COPY vendor ./vendor
RUN go mod download

COPY . .

RUN go build -o server .

EXPOSE 8080

CMD [ "./server" ]
