FROM golang:1.14 AS golang

WORKDIR /go/src/api.premiumcases.design

COPY . .

RUN apt-get update
RUN go get -d -v ./...
RUN CGO_ENABLED=0 GOOS=linux go build

FROM alpine:3.12.0

EXPOSE 80

COPY --from=golang /go/src/api.premiumcases.design/api.premiumcases.design .

CMD ["./api.premiumcases.design"]