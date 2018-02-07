FROM golang:latest

RUN mkdir -p /go/src/github.com/kapitoshka438/imgapi

WORKDIR /go/src/github.com/kapitoshka438/imgapi

COPY . /go/src/github.com/kapitoshka438/imgapi

RUN go-wrapper download

RUN go-wrapper install

CMD ["go-wrapper", "run", "-web"]

EXPOSE 8000