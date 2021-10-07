package request

type NetbankingHdfc struct {
	AuthUrl    string   `json:"auth_url"`
	UserName   string   `json:"user_name"`
	Password   string   `json:"password"`
	Timestamp  int64    `json:"timestamp"`
}
