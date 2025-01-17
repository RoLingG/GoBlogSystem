package config

type SiteInfo struct {
	CreatedAt   string `yaml:"created_at"`
	Beian       string `yaml:"beian"`
	Title       string `yaml:"title"`
	QQImage     string `yaml:"qq_image"`
	Version     string `yaml:"version"`
	Email       string `yaml:"email"`
	WechatImage string `yaml:"wechat_image"`
	Name        string `yaml:"name"`
	Job         string `yaml:"job"`
	Addr        string `yaml:"addr"`
	Slogan      string `yaml:"slogan"`
	SloganEn    string `yaml:"slogan_en"`
	Web         string `yaml:"web"`
	GithubUrl   string `yaml:"github_url"`
}
