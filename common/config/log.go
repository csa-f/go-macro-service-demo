package config

type Log struct {
	Level uint32 `json:"level"`
	Color struct {
		Red    string `json:"red"`
		Yellow string `json:"yellow"`
		Gray   string `json:"gray"`
		Def    string `json:"def"`
		Green  string `json:"green"`
	} `json:"color"`
}
