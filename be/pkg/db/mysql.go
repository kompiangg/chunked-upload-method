package db

type MySQLConfig struct {
	Username     string
	Password     string
	Host         string
	Port         string
	DatabaseName string
	ParseTime    bool
	Localization string
}
