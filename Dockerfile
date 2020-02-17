FROM golang
RUN go get github.com/gorilla/mux
RUN go get gopkg.in/mgo.v2

WORKDIR /go/src/internal-account-bank

ADD . .

RUN go build

EXPOSE 3000

ENTRYPOINT go run main.go