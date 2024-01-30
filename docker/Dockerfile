FROM debian:bullseye-slim

LABEL org.opencontainers.image.authors="anshan"
MAINTAINER GIT:github.com/jwping/wxbot

RUN sed -i 's#http://.*debian.org#http://mirrors.aliyun.com#g' /etc/apt/sources.list && \
    apt update && \
    apt install xvfb wget -y && \
    dpkg --add-architecture i386 && \
    mkdir -pm755 /etc/apt/keyrings && \
    wget -O /etc/apt/keyrings/winehq-archive.key https://dl.winehq.org/wine-builds/winehq.key && \
    wget -NP /etc/apt/sources.list.d/ https://dl.winehq.org/wine-builds/debian/dists/bullseye/winehq-bullseye.sources && \
    sed -i 's/dl.winehq.org/mirrors.tuna.tsinghua.edu.cn/g' /etc/apt/sources.list.d/winehq-bullseye.sources && \
    apt update && \
    apt install --install-recommends winehq-stable -y && \
    useradd -ms /bin/bash wxbot

USER wxbot

WORKDIR /home/wxbot

# 这里wechat的目录需要与wxbot-sidecar支持的版本一致！
COPY --chown=wxbot:wxbot ./WeChat /home/wxbot/WeChat

COPY --chown=wxbot:wxbot ./wxbot-sidecar.exe /home/wxbot

COPY --chown=wxbot:wxbot ./start.sh /home/wxbot

RUN chmod +x /home/wxbot/start.sh

CMD ["/home/wxbot/start.sh"]