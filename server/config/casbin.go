package config

import (
	"fmt"
	"os"

	"server/infrastructure"

	sqladapter "github.com/Blank-Xu/sql-adapter"
	"github.com/casbin/casbin/v2"
	_ "github.com/lib/pq"
)

func NewEnforcer() *casbin.Enforcer {
	adapter, err := sqladapter.NewAdapter(infrastructure.SqlDB, "postgres", "permissions")
	if err != nil {
		fmt.Printf("casbin: failed to initialize adapter - %v \n", err)
		os.Exit(1)
	}
	enforce, err := casbin.NewEnforcer("./casbin.conf", adapter)
	if err != nil {
		fmt.Printf("casbin: failed to create enforcer - %v \n", err)
		os.Exit(1)
	}

	enforce.AddPolicy("ADMIN", "/example/*", "(GET)|(POST)|(PATCH)|(PUT)|(DELETE)")
	enforce.AddPolicy("MASYARAKAT", "/example/*", "(GET)")
	enforce.LoadPolicy()

	return enforce
}
