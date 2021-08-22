FROM alpine:3.13.2

RUN echo -e http://mirrors.ustc.edu.cn/alpine/v3.13/main/ >> /etc/apk/repositories

COPY direkcli-linux /bin/direkcli
RUN chmod 755 /bin/direkcli

CMD /bin/direkcli
