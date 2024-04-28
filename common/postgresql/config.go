package postgresql

//postgreye hangi bilgilerle bağlanacağımız

type Config struct {
	Host                   string
	Port                   string
	UserName               string
	Password               string
	DbName                 string
	MaxConnections         string
	MaxConnectionsIdleTime string
}
