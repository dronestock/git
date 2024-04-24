package config

type Credential struct {
	// 用户名
	Username string `default:"${PLUGIN_USERNAME=${DRONE_NETRC_USERNAME=${USERNAME}}}" json:"username,omitempty"`
	// 密码
	Password string `default:"${PLUGIN_PASSWORD=${DRONE_NETRC_PASSWORD=${PASSWORD}}}" json:"password,omitempty"`
	// 密钥
	Key string `default:"${KEY}" json:"key,omitempty"`
}
