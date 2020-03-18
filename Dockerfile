FROM golang:1.13.4-alpine


WORKDIR /go/src/github.com/Rule-BasedGO/

COPY . .

RUN go build main.go

CMD ["go", "run", "main.go"]