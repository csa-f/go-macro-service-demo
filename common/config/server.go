package config

type Server struct {
	Ip   string `json:"ip"`
	Port uint16 `json:"port"`
	Name string `json:"name"`
}
