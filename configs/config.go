package configs

type Modes struct {
	Name   string `json:"name"`
	Server struct {
		Port    int    `json:"port"`
		TimeOut int64  `json:"time_out"`
		Log     string `json:"log"`
	} `json:"server"`
	Database struct {
		Name     string `json:"name"`
		User     string `json:"user"`
		Password string `json:"password"`
		Port     string `json:"port"`
		Host     string `json:"host"`
	} `json:"database"`
}
