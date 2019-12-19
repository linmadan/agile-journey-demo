package pg

import (
	"fmt"
	"github.com/go-pg/pg"
	"github.com/go-pg/pg/orm"
	"github.com/linmadan/agile-journey-demo/pkg/constant"
	"github.com/linmadan/agile-journey-demo/pkg/infrastructure/pg/models"

	//_ "github.com/linmadan/agile-journey-demo/pkg/infrastructure/pg/models"
	"github.com/linmadan/egglib-go/persistent/pg/hooks"
)

var DB *pg.DB

func init() {
	DB = pg.Connect(&pg.Options{
		User:     constant.POSTGRESQL_USER,
		Password: constant.POSTGRESQL_PASSWORD,
		Database: constant.POSTGRESQL_DB_NAME,
		Addr:     fmt.Sprintf("%s:%s", constant.POSTGRESQL_HOST, constant.POSTGRESQL_PORT),
	})
	if !constant.DISABLE_SQL_GENERATE_PRINT {
		DB.AddQueryHook(hooks.SqlGeneratePrintHook{})
	}
	if !constant.DISABLE_CREATE_TABLE {
		for _, model := range []interface{}{
			(*models.Lecturer)(nil),
		} {
			err := DB.CreateTable(model, &orm.CreateTableOptions{
				Temp:          false,
				IfNotExists:   true,
				FKConstraints: true,
			})
			if err != nil {
				panic(err)
			}
		}
	}
}
