# web-delay

Web delay simulation powered by go.

[![DockerHub Badge](http://dockeri.co/image/zhangsean/web-delay)](https://hub.docker.com/r/zhangsean/web-delay/)

## Usage

```sh
docker run -itd --name web-delay -p 8080:80 zhangsean/web-delay
```

Visit `http://localhost:8080/?ms=10` to simulate a delay of 10 ms in every request.

## Go build

```sh
go build
./delay
```

Visit `http://localhost/?ms=100` to simulate a delay of 100 ms in every request.

## Docker build

```sh
docker build -t image:tag .
docker run -itd --name web-delay -p 888:80 image:tag
```

Visit `http://localhost:888/?ms=1000` to simulate a delay of 1000 ms in every request.
