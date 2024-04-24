package config

type Pull struct {
	// 子模块
	Submodules bool `default:"${SUBMODULES=true}" json:"submodules,omitempty"`
	// 深度
	Depth int `default:"${DEPTH}" json:"depth,omitempty"`
}
