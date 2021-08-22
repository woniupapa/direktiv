FROM alpine:3.13.2

RUN echo -e http://mirrors.ustc.edu.cn/alpine/v3.13/main/ >> /etc/apk/repositories

RUN wget -O /bin/grpc_health_probe https://github.com/grpc-ecosystem/grpc-health-probe/releases/download/v0.3.6/grpc_health_probe-linux-amd64
#RUN wget -O /bin/grpc_health_probe https://minio.bjx.cloud/local/uploads/grpc_health_probe-linux-amd64?X-Amz-Algorithm=AWS4-HMAC-SHA256&X-Amz-Credential=admin%2F20210822%2F%2Fs3%2Faws4_request&X-Amz-Date=20210822T124116Z&X-Amz-Expires=432000&X-Amz-SignedHeaders=host&X-Amz-Signature=d2e39d3a8dfe74ea2c8a8c9141f334ec11d47fa67e3054959f7e60e6361978c8

RUN chmod 755 /bin/grpc_health_probe

COPY flow /bin/direktiv
RUN chmod 755 /bin/direktiv

EXPOSE 6666
EXPOSE 7777

RUN apk add shadow
RUN /usr/sbin/groupadd -g 22222 direktivg && /usr/sbin/useradd -s /bin/sh -g 22222 -u 33333 direktivu

CMD /bin/direktiv
