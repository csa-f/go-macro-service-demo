package config

type Redis struct {
	Host string `json:"host"`
	Port uint16 `json:"port"`
	User string `json:"user"`
	Pass string `json:"pass"`
}
