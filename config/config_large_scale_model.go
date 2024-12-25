package config

type ModelSetting struct {
	Name      string `yaml:"name" json:"name"`
	Enable    bool   `yaml:"enable" json:"enable"`
	ApiKey    string `yaml:"api-key" json:"api-key"`
	ApiSecret string `yaml:"api_secret" json:"api_secret"`
	Title     string `yaml:"title" json:"title"`
	Prompt    string `yaml:"prompt" json:"prompt"`
	Slogan    string `yaml:"slogan" json:"slogan"`
}

type ModelOption struct {
	Label   string `yaml:"label" json:"label"`
	Value   string `yaml:"value" json:"value"`
	Disable bool   `yaml:"disable" json:"disable"`
}

type LargeScaleModel struct {
	ModelSetting        ModelSetting        `yaml:"model_setting"`
	ModelOption         []ModelOption       `yaml:"model_option"`
	ModelSessionSetting ModelSessionSetting `yaml:"model_session"`
}

type ModelSessionSetting struct {
	ChatScope    int `yaml:"chat_scope" json:"chat_scope"`       // 对话的积分消耗
	SessionScope int `yaml:"session_scope" json:"session_scope"` // 会话的积分消耗
	DailyScope   int `yaml:"daily_scope" json:"daily_scope"`     // 每日可以领取的积分
}
