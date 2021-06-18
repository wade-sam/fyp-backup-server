FROM golang:1.16-alpine
WORKDIR  /backup-server
COPY . .
RUN go get -d -v ./...
RUN go install -v ./...

EXPOSE 8080
WORKDIR /backup-server/api/
CMD ["raspserver"]