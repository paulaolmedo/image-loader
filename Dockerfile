FROM golang:1.13

WORKDIR $GOPATH/image-loader
COPY . .

ENV GO111MODULE=on

RUN go get -d -v ./...
RUN go get -u goa.design/goa/v3
RUN go get -u goa.design/goa/v3/...

RUN cd $GOPATH/image-loader
RUN go build -o image-loader image-loader/cmd/image-loader-api

EXPOSE 8080
CMD ./image-loader -mongodb mongodb://localhost:27017

#with this first version:
# 1) docker build -t image-loader .
# 2) docker run image-loader