package config

type SiteInfo struct {
	CreatedAt   string `json:"created_at" yaml:"created_at"`
	Beian       string `json:"beian" yaml:"beian"`
	Title       string `json:"title" yaml:"title"`
	QQImage     string `json:"qq_image" yaml:"qq_image"`
	Version     string `json:"version" yaml:"version"`
	Email       string `json:"email" yaml:"email"`
	WechatImage string `json:"wechat_image" yaml:"wechat_image"`
	Name        string `json:"name" yaml:"name"`
	Job         string `json:"job" yaml:"job"`
	Addr        string `json:"addr" yaml:"addr"`
	Slogan      string `json:"slogan" yaml:"slogan"`
	SloganEn    string `json:"slogan_en" yaml:"slogan_en"`
	Web         string `json:"web" yaml:"web"`
	GithubUrl   string `json:"github_url" yaml:"github_url"`
	GiteeUrl    string `json:"gitee_url" yaml:"gitee_url"`
}
