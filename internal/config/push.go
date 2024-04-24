package config

type Push struct {
	// 标签
	Tag string `default:"${TAG}" json:"tag,omitempty"`
	// 作者
	Author string `default:"${AUTHOR=${DRONE_COMMIT_AUTHOR_NAME}}" json:"author,omitempty"`
	// 邮箱
	Email string `default:"${EMAIL=${DRONE_COMMIT_AUTHOR_EMAIL}}" json:"email,omitempty"`
	// 提交消息
	Message string `default:"${MESSAGE=${DRONE_COMMIT_MESSAGE=drone}}" json:"message,omitempty"`
	// 是否强制提交
	Force *bool `default:"${FORCE}" json:"force,omitempty"`
}
