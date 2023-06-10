package sqlx

import (
	"errors"
	"fmt"

	_ "github.com/go-sql-driver/mysql"
	"github.com/jmoiron/sqlx"
)

type gormConfig struct {
	Username     string
	Password     string
	Host         string
	Port         string
	DatabaseName string
	ParseTime    bool
	Localization string
}

func Init(param *SqlxConfig) (db *sqlx.DB, driver string, err error) {
	if param == nil {
		return db, driver, errors.New("[ERROR] config must not be nil")
	}

	var dsn string

	if param.MySQLConfig != nil {
		driver = "mysql"
		dsn = initMySQL(gormConfig{
			Username:     param.MySQLConfig.Username,
			Password:     param.MySQLConfig.Password,
			Host:         param.MySQLConfig.Host,
			Port:         param.MySQLConfig.Port,
			DatabaseName: param.MySQLConfig.DatabaseName,
			ParseTime:    param.MySQLConfig.ParseTime,
			Localization: param.MySQLConfig.Localization,
		})
	} else {
		return db, driver, errors.New("[ERROR] need to defined at least one driver that available")
	}

	db, err = sqlx.Connect(driver, dsn)
	if err != nil {
		return db, driver, err
	}

	return db, driver, nil
}

func initMySQL(config gormConfig) string {
	dsn := fmt.Sprintf(
		"%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=%t&loc=%s",
		config.Username, config.Password,
		config.Host, config.Port, config.DatabaseName,
		config.ParseTime, config.Localization,
	)

	return dsn
}
