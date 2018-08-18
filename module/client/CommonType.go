package client

type ConnectOptionType struct {
	dbType   string `json:"dbType"`
	Host     string `json:"host"`
	Port     string `json:"port"`
	Database string `json:"database"`
	Username string `json:"username"`
	Password string `json:"password"`
}
