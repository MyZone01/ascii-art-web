FROM golang:alpine

LABEL MAINTAINER: "ndiayetapha7@gmail.com, niangpmaxgatte@gmail.com, ssmm0498@gail.com" 

WORKDIR /app
COPY . /app

RUN go build -o serveur .

RUN apk update && apk add bash && apk add tree \
	rm -rf /var/cache/apk/* /tmp/*
EXPOSE 8080

CMD ["./serveur"]
