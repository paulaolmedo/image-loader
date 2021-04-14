# Image loader 🛰️

## status

[![Tests](https://github.com/paulaolmedo/image-loader/actions/workflows/go.yml/badge.svg)](https://github.com/paulaolmedo/image-loader/actions/workflows/go.yml)

[![Build docker image](https://github.com/paulaolmedo/image-loader/actions/workflows/docker.yml/badge.svg)](https://github.com/paulaolmedo/image-loader/actions/workflows/docker.yml)

## how it works

### with 🐳
Build the docker image:
 ```
 docker build -f ci/Dockerfile -t image-loader .
 ```
 
 And then run the compose:
  ```
 docker-compose -f ci/docker-compose.yml up
 ```


### without docker ⚒️
Then, to build the application run:
 ```
 go build -o image-loader main.go
 ```

And finally:
```
./image-loader -mongodb mongodb://localhost:27017
```
