FROM storezhang/alpine AS fastgithub


# Github加速版本
ENV FAST_GITHUB_VERSION 2.1.2
WORKDIR /opt


RUN apk add unzip
RUN wget https://ghproxy.com/https://github.com/dotnetcore/FastGithub/releases/download/${FAST_GITHUB_VERSION}/fastgithub_linux-x64.zip --output-document fastgithub_linux-x64.zip
RUN unzip fastgithub_linux-x64.zip
RUN mv fastgithub_linux-x64 /opt/fastgithub
RUN chmod +x /opt/fastgithub/fastgithub



# 打包真正的镜像
FROM storezhang/alpine


LABEL author="storezhang<华寅>"
LABEL email="storezhang@gmail.com"
LABEL qq="160290688"
LABEL wechat="storezhang"
LABEL description="Drone持续集成Git插件，增加标签功能以及Github加速功能。同时支持推拉模式"


# 复制文件
COPY --from=fastgithub /opt/fastgithub /opt/fastgithub
COPY docker /
COPY git /bin


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
    \
    \
    \
    # 加速Github
    && su-exec ${USERNAME} echo "199.232.69.194 github.global.ssl.fastly.net" > /etc/hosts \
    && su-exec ${USERNAME} echo "140.82.113.4 github.com" > /etc/hosts \
    && su-exec ${USERNAME} echo "140.82.114.4 github.com" > /etc/hosts \
    \
    \
    \
    && rm -rf /var/cache/apk/*



# 执行命令
ENTRYPOINT /bin/git


# 配置环境变量
ENV GOPROXY https://goproxy.io,https://goproxy.cn,https://mirrors.aliyun.com/goproxy,direct
