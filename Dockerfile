FROM slcnx/fastgithub AS fastgithub
FROM ccr.ccs.tencentyun.com/storezhang/alpine:3.20.0 AS builder

# 复制加速程序
COPY --from=fastgithub /fastgithub /docker/opt/fastgithub
# 复制脚本程序
COPY docker /docker
# 复制执行程序
ARG TARGETPLATFORM
COPY dist/${TARGETPLATFORM}/git /docker/usr/local/bin/



FROM ccr.ccs.tencentyun.com/storezhang/alpine:3.20.0

LABEL author="storezhang<华寅>" \
    email="storezhang@gmail.com" \
    qq="160290688" \
    wechat="storezhang" \
    description="Drone持续集成Git插件，增加标签功能以及Github加速功能。同时支持推拉模式"

# 复制文件，复合成一个步骤
COPY --from=builder /docker /

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
    && chmod +x /usr/local/bin/* \
    \
    \
    \
    && rm -rf /var/cache/apk/*

# 修改默认参数
ENV PLUGIN_TIMES 10

# 执行命令
ENTRYPOINT /usr/local/bin/bootstrap
