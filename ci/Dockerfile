FROM golang:1.16-alpine as builder
RUN apk add --update gcc musl-dev
ENV IMAGE_LOADER $GOPATH/image-loader
WORKDIR $IMAGE_LOADER
COPY . .
ENV GO111MODULE=on
RUN GOOS=linux go build -o image-loader main.go

FROM alpine:latest
ENV IMAGE_LOADER /home 
WORKDIR $IMAGE_LOADER
COPY --from=builder /go/image-loader/image-loader .
COPY --from=builder /go/image-loader/app.properties .
CMD ./image-loader -mongodb mongodb://localhost:27018