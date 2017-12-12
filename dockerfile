FROM alpine:latest
MAINTAINER "382023823@qq.com"

RUN mkdir -p /opt/ams
ADD ams /opt/ams/ams
ADD conf /opt/ams/
ADD views /opt/ams/

EXPOSE 8080
EXPOSE 8088 
ENTRYPOINT ["/opt/ams/ams"]
