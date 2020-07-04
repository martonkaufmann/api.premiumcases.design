FROM golang:1.14

EXPOSE 1234

WORKDIR /go/src/api.premiumcases.design

COPY . .

RUN apt-get update
RUN go get -d -v ./...

CMD ["tail", "-f", "/dev/null"]