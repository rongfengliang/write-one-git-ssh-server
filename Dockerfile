FROM linuxserver/openssh-server
RUN sed -i 's/dl-cdn.alpinelinux.org/mirrors.aliyun.com/g' /etc/apk/repositories
COPY app.sh /opt/app.sh
COPY --from=dalongrong/git-shell /app/git-shell /opt/git-shell
RUN chmod +x /opt/app.sh && apk add git --no-cache