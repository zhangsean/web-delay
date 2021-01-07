FROM golang:alpine AS build
COPY . src/
RUN cd src \
 && CGO_ENABLED=0 go build -ldflags="-s -w"

FROM nginx:alpine
LABEL maintainer="zhangsean <zxf2342@qq.com>"
COPY --from=build /go/src/web-delay /web-delay
COPY default.conf /etc/nginx/conf.d/
COPY start.sh /
EXPOSE 80
CMD [ "/bin/sh", "/start.sh" ]
