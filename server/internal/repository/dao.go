package repository

import (
	"database/sql"

	"github.com/casbin/casbin/v2"
	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
	"github.com/redis/go-redis/v9"
	"gorm.io/gorm"
)

type DAO interface {
	NewExampleQuery() ExampleQuery
	NewAuthenticationQuery() AuthenticationQuery
	NewObjectQuery() ObjectQuery
}

type dao struct{}

var SqlDB *sql.DB
var SqlxDB *sqlx.DB
var GormDB *gorm.DB
var Cache *redis.Client
var Permission *casbin.Enforcer

func NewDAO(sqlDB *sql.DB, sqlxDB *sqlx.DB, gormDB *gorm.DB, cache *redis.Client, permission *casbin.Enforcer) DAO {
	SqlDB = sqlDB
	SqlxDB = sqlxDB
	GormDB = gormDB
	Cache = cache
	Permission = permission
	return &dao{}
}

func (d *dao) NewExampleQuery() ExampleQuery {
	return &exampleQuery{
		sqlDB:  SqlDB,
		gormDB: GormDB,
		sqlxDB: SqlxDB,
	}
}

func (d *dao) NewAuthenticationQuery() AuthenticationQuery {
	return &authenticationQuery{
		sqlDB:  SqlDB,
		gormDB: GormDB,
		cache:  Cache,
	}
}

func (d *dao) NewObjectQuery() ObjectQuery {
	return &objectQuery{
		sqlDB:  SqlDB,
		gormDB: GormDB,
	}
}
