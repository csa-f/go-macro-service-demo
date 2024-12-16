package config

type Redis struct {
	Host         string `json:"host"`
	Port         uint16 `json:"port"`
	User         string `json:"user"`
	Pass         string `json:"pass"`
	DB           int    `json:"db"`
	DialTimeout  int64  `json:"dialTimeout"`
	ReadTimeout  int64  `json:"readTimeout"`
	WriteTimeout int64  `json:"writeTimeout"`
	PoolSize     int    `json:"poolSize"`
	PoolTimeout  int64  `json:"poolTimeout"`
}
