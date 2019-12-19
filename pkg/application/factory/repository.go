package factory

import (
	"github.com/linmadan/agile-journey-demo/pkg/domain/lecturer"
	"github.com/linmadan/agile-journey-demo/pkg/infrastructure/repository"
	"github.com/linmadan/egglib-go/transaction/pg"
)

func CreateLecturerRepository(options map[string]interface{}) (lecturer.LecturerRepository, error) {
	var transactionContext *pg.TransactionContext
	if value, ok := options["transactionContext"]; ok {
		transactionContext = value.(*pg.TransactionContext)
	}
	return repository.NewLecturerRepository(transactionContext)
}
