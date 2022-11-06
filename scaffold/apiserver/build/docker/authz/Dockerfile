FROM BASE_IMAGE
LABEL maintainer="<rebirthmonkey@gmail.com>"

WORKDIR /opt/apiserver

RUN ln -sf /usr/share/zoneinfo/Asia/Shanghai /etc/localtime && \
      echo "Asia/Shanghai" > /etc/timezone && \
      mkdir -p /var/log/apiserver

COPY authz /opt/apiserver/bin/

ENTRYPOINT ["/opt/apiserver/bin/authz"]
CMD ["-c", "/etc/apiserver/authz.yaml"]
