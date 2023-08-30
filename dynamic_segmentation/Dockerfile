FROM golang:1.21

WORKDIR /go/src/app

COPY . .

RUN go mod download

RUN go install -tags 'postgres' github.com/golang-migrate/migrate/v4/cmd/migrate@latest

RUN go build -o dynamic_segmentation cmd/dynamic_segmentation/main.go

EXPOSE 8080

CMD ["./dynamic_segmentation"]
