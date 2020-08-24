@wip

builded with goa design (v3) for go 1.13 (i must update to 1.14!) 

how to build the endpoints:
goa gen image-loader/design

how to build the application
go build -o image-loader image-loader/cmd/image-loader-api

how to run the application (please note that you must have a running mongo)
./image-loader -mongodb mongodb://localhost:27017