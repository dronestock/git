FROM alpine

MAINTAINER storezhang "storezhang@gmail.com"
LABEL architecture="AMD64/x86_64" version="latest" build="2021-06-04"
LABEL Description="Drone持续集成Git插件，增加标签功能。"



# 复制文件
COPY git.sh /bin
RUN chmod +x /bin/git.sh



ENTRYPOINT /bin/git.sh
