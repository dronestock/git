#!/bin/sh

# 清空本地缓存
if [ "${PLUGIN_CLEAR}" = true ] ; then
  rm -rf .git
fi

# 写入授权文件，在使用过程中无需输入密码
mkdir -pv "${HOME}"/.ssh
echo "${PLUGIN_SSH_KEY}" > "${HOME}"/.ssh/id_rsa
chmod 400 "${HOME}"/.ssh/id_rsa


# 执行最终配置值，支持环境变量
PLUGIN_PATH=$(eval echo "${PLUGIN_PATH}")
cd "${PLUGIN_PATH}" || exit

# 初始化Git仓库
export GIT_SSH_COMMAND='ssh -o StrictHostKeyChecking=no'
git config --global user.name "${DRONE_COMMIT_AUTHOR}"
git config --global user.email "${DRONE_COMMIT_AUTHOR_EMAIL}"
git init

# 将当前目录添加到仓库中
git add .
[ -z "${PLUGIN_COMMIT_MESSAGE}" ] && PLUGIN_COMMIT_MESSAGE='push'
[ "${PLUGIN_FORCE}" = "true" ] && PLUGIN_FORCE='--force'
git commit . -m "${PLUGIN_COMMIT_MESSAGE}"

# 添加远程服务器
git remote add origin "${PLUGIN_REMOTE}"

# 打标签，如果没有设置值，使用默认环境变量的值
[ -z "${PLUGIN_TAG}" ] && PLUGIN_TAG=${TAG}
[ -z "${PLUGIN_TAG}" ] && PLUGIN_TAG=${VERSION}
[ -z "${PLUGIN_TAG}" ] && PLUGIN_TAG=${DRONE_SEMVER_SHORT}
[ -z "${PLUGIN_TAG}" ] && PLUGIN_TAG=${DRONE_TAG}

# 执行最终配置值，支持环境变量
PLUGIN_TAG=$(eval echo "${PLUGIN_TAG}")

# 判断是否需要打标签，如果需要打上并推送到服务器
if [ -n "${PLUGIN_TAG}" ]; then
  git tag -a "${PLUGIN_TAG}" -m "${PLUGIN_COMMIT_MESSAGE}"
  git push --set-upstream origin "${PLUGIN_TAG}"
fi

# 推送代码到远程仓库
git push --set-upstream origin master "${PLUGIN_FORCE}"
