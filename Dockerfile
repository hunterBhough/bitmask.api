FROM golang:1.9

WORKDIR /go/src/app
COPY . .

RUN go get -d -v ./...
WORKDIR /go/src/github.com/hunterBhough/go-doge
RUN go install ./src/cmd/go-doge-server/
ENTRYPOINT /go/bin/go-doge-server --port 3000 --host 0.0.0.0

EXPOSE 3000
# RUN go-wrapper install    # "go install -v ./..."

CMD ["go-doge-server"]