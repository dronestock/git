#!/bin/sh

# 清空本地缓存
if [ "${PLUGIN_CLEAR}" = true ] ; then
  rm -rf .git
fi

# 写入授权文件，在使用过程中无需输入密码
mkdir -pv "${HOME}"/.ssh
echo "${PLUGIN_SSH_KEY}" > "${HOME}"/.ssh/id_rsa
chmod 400 "${HOME}"/.ssh/id_rsa


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

# 打标签
echo "-------------"
echo "${PLUGIN_TAG}"
echo "-------------"
if [ -n "${PLUGIN_TAG}" ]; then
  git tag -a "${PLUGIN_TAG}" -m ${PLUGIN_COMMIT_MESSAGE}
fi

# 添加远程服务器并推送代码
git remote add origin "${PLUGIN_REMOTE}"
git push --set-upstream origin master "${PLUGIN_FORCE}" "${PLUGIN_TAG}"
