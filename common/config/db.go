package config

type DB struct {
	Driver    string `json:"driver"`
	Host      string `json:"host"`
	Port      uint32 `json:"port"`
	User      string `json:"user"`
	Pass      string `json:"pass"`
	DBName    string `json:"dbName"`
	Charset   string `json:"charset"`
	Threshold int64  `json:"threshold"`
}
