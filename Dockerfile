FROM alpine

MAINTAINER storezhang "storezhang@gmail.com"
LABEL architecture="AMD64/x86_64" version="latest" build="2021-10-12"
LABEL Description="Drone持续集成Git插件，增加标签功能。"


ENV LANG="zh_CN.UTF-8"
ENV TIMEZONE=Asia/Chongqing


# 复制文件
COPY git.sh /bin


RUN set -ex \
    \
    \
    \
    && apk update \
    && apk --no-cache add tzdata git openssh-client \
    \
    \
    \
    && cp "/usr/share/zoneinfo/${TIMEZONE}" /etc/localtime \
    && echo "${TIMEZONE}" > /etc/timezone \
    && echo "export LC_ALL=${LANG}" >> /etc/profile \
    \
    \
    \
    # 增加执行权限
    && chmod +x /bin/git.sh \
    \
    \
    \
    && rm -rf /var/cache/apk/*



ENTRYPOINT /bin/git.sh
