package conf

type Server struct {
	Ip       string `json:"ip"`
	Port     string `json:"port"`
	Name     string `json:"name"`
	LogLevel uint8  `json:"logLevel"`
}
