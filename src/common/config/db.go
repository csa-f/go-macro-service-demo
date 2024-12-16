package config

type DB struct {
	Driver     string `json:"driver"`
	Host       string `json:"host"`
	Port       uint32 `json:"port"`
	User       string `json:"user"`
	Pass       string `json:"pass"`
	DBName     string `json:"dbName"`
	CharSet    string `json:"charSet"`
	Threshould int64  `json:"threshould"`
}
