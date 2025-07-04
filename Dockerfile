FROM golang:1.12.0-alpine3.9
RUN mkdir /app
ADD . /app
WORKDIR /app
RUN go build -o app .
CMD ["/app/app"]

#docker build -t golang-web-app .
#docker run -p 8080:8080 -it golang-web-app