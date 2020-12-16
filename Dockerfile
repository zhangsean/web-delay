FROM golang:alpine AS build
COPY . src/
RUN cd src \
 && CGO_ENABLED=0 go build -ldflags="-s -w"

FROM scratch
LABEL maintainer="zhangsean <zxf2342@qq.com>"
COPY --from=build /go/src/web-delay /web-delay
EXPOSE 80
CMD [ "/web-delay" ]
