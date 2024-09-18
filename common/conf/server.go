package conf

type Server struct {
	Ip       string `json:"ip"`
	Port     string `json:"port"`
	Name     string `json:"name"`
	LogLevel string `json:"logLevel"`
}
