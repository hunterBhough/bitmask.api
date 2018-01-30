FROM golang:1.9

ENV GOBIN $GOPATH/bin

ADD . /go/src/github.com/hunterBhough/go-doge
WORKDIR /go/src/github.com/hunterBhough/go-doge


RUN go get -d -v ./...
RUN go install ./src/cmd/go-doge-server/
ENTRYPOINT /go/bin/go-doge-server --port 3000 --host 0.0.0.0

EXPOSE 3000
# RUN go-wrapper install    # "go install -v ./..."

CMD ["go-doge-server"]