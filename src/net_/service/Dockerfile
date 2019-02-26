FROM centos:7
MAINTAINER wp
RUN mkdir -p /root/websocket-server
COPY ./bin/websocker-server /root/websocket-server/
WORKDIR /root/websocket-server
EXPOSE 9999
ENTRYPOINT ./websocker-server
