package infrastructure

import (
	"database/sql"
	"fmt"
	"os"
	"time"

	"server/env"

	"github.com/jmoiron/sqlx"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

var SqlDB *sql.DB
var GormDB *gorm.DB
var SqlxDB *sqlx.DB

func ConnectSqlDB() {
	dsn := fmt.Sprintf(`
		host=%v user=%v password=%v dbname=%v port=%v sslmode=disable TimeZone=%v
	`, env.POSTGRE_HOST, env.POSTGRE_USERNAME, env.POSTGRE_PASSWORD, env.POSTGRE_DATABASE, env.POSTGRE_PORT, env.POSTGRE_TIMEZONE)
	var err error
	SqlDB, err = sql.Open("postgres", dsn)
	if err != nil {
		fmt.Println("sql database: can't connect to database")
		os.Exit(1)
	}
	SqlDB.SetMaxIdleConns(env.POSTGRE_CONN_MAX_IDLE)
	SqlDB.SetMaxOpenConns(env.POSTGRE_CONN_MAX_OPEN)
	SqlDB.SetConnMaxLifetime(time.Minute * env.POSTGRE_CONN_MAX_LIFETIME)
	if err := SqlDB.Ping(); err != nil {
		fmt.Printf("sql database: can't ping to database - %v \n", err)
		os.Exit(1)
	}

	fmt.Println("sql database: connection opened to database")
}

func ConnectGormDB() {
	var err error
	GormDB, err = gorm.Open(postgres.New(postgres.Config{
		Conn: SqlDB,
	}), &gorm.Config{
		Logger: logger.Default.LogMode(logger.Silent),
	})
	if err != nil {
		fmt.Println("gorm database: can't connect to database")
		os.Exit(1)
	}
	fmt.Println("gorm database: connection opened to database")
}

func ConnectSqlxDB() {
	dsn := fmt.Sprintf(`
		host=%v user=%v password=%v dbname=%v port=%v sslmode=disable TimeZone=%v
	`, env.POSTGRE_HOST, env.POSTGRE_USERNAME, env.POSTGRE_PASSWORD, env.POSTGRE_DATABASE, env.POSTGRE_PORT, env.POSTGRE_TIMEZONE)
	var err error
	SqlxDB, err = sqlx.Connect("postgres", dsn)
	if err != nil {
		fmt.Println("sqlx database: can't connect to database")
		os.Exit(1)
	}
	SqlxDB.SetMaxIdleConns(env.POSTGRE_CONN_MAX_IDLE)
	SqlxDB.SetMaxOpenConns(env.POSTGRE_CONN_MAX_OPEN)
	SqlxDB.SetConnMaxLifetime(time.Minute * env.POSTGRE_CONN_MAX_LIFETIME)
	if err := SqlxDB.Ping(); err != nil {
		fmt.Printf("sqlx database: can't ping to database - %v \n", err)
		os.Exit(1)
	}

	fmt.Println("sqlx database: connection opened to database")
}
