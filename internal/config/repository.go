package config

type Repository struct {
	// 远程仓库地址
	Remote string `default:"${REMOTE=${DRONE_GIT_HTTP_URL}}" validate:"required" json:"remote,omitempty"`
	// 分支
	Branch string `default:"${BRANCH=master}" json:"branch,omitempty"`
	// 提交
	Commit string `default:"${COMMIT=${DRONE_COMMIT}}" validate:"required_without=Branch"`
}

func (r *Repository) Checkout() (checkout string) {
	if "" != r.Commit {
		checkout = r.Commit
	} else {
		checkout = r.Branch
	}

	return
}
