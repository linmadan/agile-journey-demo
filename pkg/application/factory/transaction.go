package factory

import (
	"github.com/linmadan/agile-journey-demo/pkg/infrastructure/pg"
	"github.com/linmadan/egglib-go/core/application"
	pG "github.com/linmadan/egglib-go/transaction/pg"
)

func CreateTransactionContext(options map[string]interface{}) (application.TransactionContext, error) {
	return pG.NewPGTransactionContext(pg.DB), nil
}
