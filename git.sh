#!/bin/sh

# 匹配不是Drone插件时的使用配置
[ -z "${PLUGIN_SSH_KEY}" ] && PLUGIN_SSH_KEY=${SSH_KEY}
[ -z "${PLUGIN_TAG}" ] && PLUGIN_TAG=${TAG}
[ -z "${PLUGIN_PATH}" ] && PLUGIN_PATH=${PATH}
[ -z "${PLUGIN_MODE}" ] && PLUGIN_MODE=${MODE}
[ -z "${PLUGIN_BRANCH}" ] && PLUGIN_BRANCH=${BRANCH}
[ -z "${DRONE_COMMIT_AUTHOR}" ] && DRONE_COMMIT_AUTHOR=${COMMIT_AUTHOR}
[ -z "${DRONE_COMMIT_AUTHOR_EMAIL}" ] && DRONE_COMMIT_AUTHOR_EMAIL=${COMMIT_AUTHOR_EMAIL}
[ -z "${PLUGIN_COMMIT_MESSAGE}" ] && PLUGIN_COMMIT_MESSAGE=${COMMIT_MESSAGE}
[ -z "${PLUGIN_REMOTE}" ] && PLUGIN_REMOTE=${REMOTE}
[ -z "${PLUGIN_FORCE}" ] && PLUGIN_FORCE=${FORCE}
[ -z "${PLUGIN_CLEAR}" ] && PLUGIN_CLEAR=${CLEAR}

[ -z "${PLUGIN_COMMIT_MESSAGE}" ] && PLUGIN_COMMIT_MESSAGE='push'
[ -z "${PLUGIN_FORCE}" ] && PLUGIN_FORCE="--force"
[ -z "${PLUGIN_CLEAR}" ] && PLUGIN_CLEAR=true
[ -z "${PLUGIN_BRANCH}" ] && PLUGIN_CLEAR=${DRONE_BRANCH}

# 判断是推还是拉模式，判断方法
# 如果环境变量DRONE_STEP_NUMBER为1，在Drone里面表示这是进行第一个步骤，为推模式
# 其它情况则取环境变量TYPE的值来做判断

PULL=false
if [ "${DRONE_STEP_NUMBER}" = 1 ] ; then
  PULL=true
fi
if [ "${PLUGIN_MODE}" = "push" ] ; then
  PULL=false
else
  PULL=true
fi
if [ -z "${PLUGIN_REMOTE}" ] && [ "${PULL}" = false ] ; then
  PLUGIN_REMOTE=${DRONE_GIT_SSH_URL}
  [ -z "${PLUGIN_REMOTE}" ] && PLUGIN_REMOTE=${DRONE_GIT_HTTP_URL}
fi

# 处理配置带环境变量的情况
PLUGIN_SSH_KEY=$(eval echo "${PLUGIN_SSH_KEY}")
PLUGIN_TAG=$(eval echo "${PLUGIN_TAG}")
PLUGIN_PATH=$(eval echo "${PLUGIN_PATH}")
[ -z "${PLUGIN_PATH}" ] && PLUGIN_PATH="."

# 确认目录存在
cd "${PLUGIN_PATH}" || exit

# 清空本地缓存
if [ "${PLUGIN_CLEAR}" = true ] ; then
  echo "清空本地缓存"
  rm -rf .git
fi

# 写入授权文件，在使用过程中无需输入密码
if [ -n "${PLUGIN_CLEAR}" ] ; then
  echo "写入授权文件"
  mkdir -pv "${HOME}"/.ssh
  echo "${PLUGIN_SSH_KEY}" > "${HOME}"/.ssh/id_rsa
  chmod 400 "${HOME}"/.ssh/id_rsa
  export GIT_SSH_COMMAND='ssh -o StrictHostKeyChecking=no'
fi


# 执行真正的命令
if [ "${PULL}" = true ] ; then
  git clone  --registry "${PLUGIN_REMOTE}" --branch "${PLUGIN_BRANCH}" --remote-submodules --recurse-submodules --depth 50
else
  git config --global user.name "${DRONE_COMMIT_AUTHOR}"
  git config --global user.email "${DRONE_COMMIT_AUTHOR_EMAIL}"
  git init

  # 将当前目录添加到仓库中
  git add .
  git commit . -m "${PLUGIN_COMMIT_MESSAGE}"

  # 添加远程服务器
  git remote add origin "${PLUGIN_REMOTE}"

  # 判断是否需要打标签，如果需要打上并推送到服务器
  if [ -n "${PLUGIN_TAG}" ]; then
    git tag -a "${PLUGIN_TAG}" -m "${PLUGIN_COMMIT_MESSAGE}"
    git push --set-upstream origin "${PLUGIN_TAG}"
  fi

  # 推送代码到远程仓库
  git push --set-upstream origin master "${PLUGIN_FORCE}"
fi
