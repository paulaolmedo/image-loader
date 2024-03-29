# Image loader 🛰️

## how it works
### with 🐳

Build the docker image:
 ```
 docker build -f ci/Dockerfile -t paulabeatrizolmedo/image-loader .
 ```
 
 And then run the compose:
  ```
 docker-compose -f ci/docker-compose.yml up
 ```


### without docker ⚒️
To build the application run:
 ```
 go build -o image-loader image-loader/cmd/image-loader-api
 ```

And finally:
```
./image-loader -mongodb mongodb://localhost:27017
```

## API documentation
- [Gitlab pages](https://satellite-forecast.gitlab.io/image-loader/)
