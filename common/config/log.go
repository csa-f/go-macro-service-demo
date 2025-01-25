package config

import "github.com/sirupsen/logrus"

type Log struct {
	Level logrus.Level `json:"level"`
	Color LogColor     `json:"color"`
}

type LogColor struct {
	Debug string `json:"debug"`
	Trace string `json:"trace"`
	Warn  string `json:"warn"`
	Error string `json:"error"`
	Fatal string `json:"fatal"`
	Panic string `json:"panic"`
	Info  string `json:"info"`
}
