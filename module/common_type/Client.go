package common_type

type ConnectOptionType struct {
	DbType   string `json:"dbType"`
	Host     string `json:"host"`
	Port     string `json:"port"`
	Database string `json:"database"`
	Username string `json:"username"`
	Password string `json:"password"`
}