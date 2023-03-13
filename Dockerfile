FROM dockerproxy.com/slcnx/fastgithub AS fastgithub





FROM ccr.ccs.tencentyun.com/storezhang/alpine:3.17.2


LABEL author="storezhang<华寅>" \
    email="storezhang@gmail.com" \
    qq="160290688" \
    wechat="storezhang" \
    description="Drone持续集成Git插件，增加标签功能以及Github加速功能。同时支持推拉模式"


# 复制文件
COPY --from=fastgithub /fastgithub /opt/fastgithub
COPY docker /


RUN set -ex \
    \
    \
    \
    && apk update \
    \
    # 安装FastGithub依赖库 \
    && apk --no-cache add libgcc libstdc++ gcompat icu \
    \
    # 安装Git工具
    && apk --no-cache add git openssh-client \
    \
    \
    \
    # 增加执行权限
    && chmod +x /bin/git \
    && chmod +x /bin/gw \
    \
    \
    \
    && rm -rf /var/cache/apk/*


# 修改默认参数
ENV PLUGIN_TIMES 10


# 执行命令
ENTRYPOINT /bin/gw
