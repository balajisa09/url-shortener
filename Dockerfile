FROM golang:1.18-alpine

WORKDIR /app

COPY . ./

RUN go mod download

RUN go build .

EXPOSE 8080

CMD [ "./url-shortener" ]