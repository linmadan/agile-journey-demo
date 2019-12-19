package service

import (
	"fmt"
	"github.com/linmadan/agile-journey-demo/pkg/application/factory"
	"github.com/linmadan/agile-journey-demo/pkg/application/lecturer/command"
	"github.com/linmadan/agile-journey-demo/pkg/application/lecturer/query"
	"github.com/linmadan/agile-journey-demo/pkg/domain/lecturer"
	"github.com/linmadan/egglib-go/core/application"
	"github.com/linmadan/egglib-go/utils/tool_funs"
)

// 讲师服务
type LecturerService struct {
}

// 创建新讲师
func (lecturerService *LecturerService) CreateLecturer(createLecturerCommand *command.CreateLecturerCommand) (interface{}, error) {
	if err := createLecturerCommand.ValidateCommand(); err != nil {
		return nil, application.ThrowError(application.ARG_ERROR, err.Error())
	}
	transactionContext, err := factory.CreateTransactionContext(nil)
	if err != nil {
		return nil, application.ThrowError(application.TRANSACTION_ERROR, err.Error())
	}
	if err := transactionContext.StartTransaction(); err != nil {
		return nil, application.ThrowError(application.TRANSACTION_ERROR, err.Error())
	}
	defer func() {
		transactionContext.RollbackTransaction()
	}()
	newLecturer := &lecturer.Lecturer{
		Name:         createLecturerCommand.Name,
		Introduction: createLecturerCommand.Introduction,
		Topic:        createLecturerCommand.Topic,
	}
	var lecturerRepository lecturer.LecturerRepository
	if value, err := factory.CreateLecturerRepository(map[string]interface{}{
		"transactionContext": transactionContext,
	}); err != nil {
		return nil, application.ThrowError(application.INTERNAL_SERVER_ERROR, err.Error())
	} else {
		lecturerRepository = value
	}
	if lecturer, err := lecturerRepository.Save(newLecturer); err != nil {
		return nil, application.ThrowError(application.INTERNAL_SERVER_ERROR, err.Error())
	} else {
		if err := transactionContext.CommitTransaction(); err != nil {
			return nil, application.ThrowError(application.TRANSACTION_ERROR, err.Error())
		}
		return lecturer, nil
	}
}

// 返回讲师
func (lecturerService *LecturerService) GetLecturer(getLecturerQuery *query.GetLecturerQuery) (interface{}, error) {
	if err := getLecturerQuery.ValidateQuery(); err != nil {
		return nil, application.ThrowError(application.ARG_ERROR, err.Error())
	}
	transactionContext, err := factory.CreateTransactionContext(nil)
	if err != nil {
		return nil, application.ThrowError(application.TRANSACTION_ERROR, err.Error())
	}
	if err := transactionContext.StartTransaction(); err != nil {
		return nil, application.ThrowError(application.TRANSACTION_ERROR, err.Error())
	}
	defer func() {
		transactionContext.RollbackTransaction()
	}()
	var lecturerRepository lecturer.LecturerRepository
	if value, err := factory.CreateLecturerRepository(map[string]interface{}{
		"transactionContext": transactionContext,
	}); err != nil {
		return nil, application.ThrowError(application.INTERNAL_SERVER_ERROR, err.Error())
	} else {
		lecturerRepository = value
	}
	lecturer, err := lecturerRepository.FindOne(map[string]interface{}{"lecturerId": getLecturerQuery.LecturerId})
	if err != nil {
		return nil, application.ThrowError(application.INTERNAL_SERVER_ERROR, err.Error())
	}
	if lecturer == nil {
		return nil, application.ThrowError(application.RES_NO_FIND_ERROR, fmt.Sprintf("%s", string(getLecturerQuery.LecturerId)))
	} else {
		if err := transactionContext.CommitTransaction(); err != nil {
			return nil, application.ThrowError(application.TRANSACTION_ERROR, err.Error())
		}
		return lecturer, nil
	}
}

// 更新讲师
func (lecturerService *LecturerService) UpdateLecturer(updateLecturerCommand *command.UpdateLecturerCommand) (interface{}, error) {
	if err := updateLecturerCommand.ValidateCommand(); err != nil {
		return nil, application.ThrowError(application.ARG_ERROR, err.Error())
	}
	transactionContext, err := factory.CreateTransactionContext(nil)
	if err != nil {
		return nil, application.ThrowError(application.TRANSACTION_ERROR, err.Error())
	}
	if err := transactionContext.StartTransaction(); err != nil {
		return nil, application.ThrowError(application.TRANSACTION_ERROR, err.Error())
	}
	defer func() {
		transactionContext.RollbackTransaction()
	}()
	var lecturerRepository lecturer.LecturerRepository
	if value, err := factory.CreateLecturerRepository(map[string]interface{}{
		"transactionContext": transactionContext,
	}); err != nil {
		return nil, application.ThrowError(application.INTERNAL_SERVER_ERROR, err.Error())
	} else {
		lecturerRepository = value
	}
	lecturer, err := lecturerRepository.FindOne(map[string]interface{}{"lecturerId": updateLecturerCommand.LecturerId})
	if err != nil {
		return nil, application.ThrowError(application.INTERNAL_SERVER_ERROR, err.Error())
	}
	if lecturer == nil {
		return nil, application.ThrowError(application.RES_NO_FIND_ERROR, fmt.Sprintf("%s", string(updateLecturerCommand.LecturerId)))
	}
	if err := lecturer.Update(tool_funs.SimpleStructToMap(updateLecturerCommand)); err != nil {
		return nil, application.ThrowError(application.BUSINESS_ERROR, err.Error())
	}
	if lecturer, err := lecturerRepository.Save(lecturer); err != nil {
		return nil, application.ThrowError(application.INTERNAL_SERVER_ERROR, err.Error())
	} else {
		if err := transactionContext.CommitTransaction(); err != nil {
			return nil, application.ThrowError(application.TRANSACTION_ERROR, err.Error())
		}
		return lecturer, nil
	}
}

// 移除讲师
func (lecturerService *LecturerService) RemoveLecturer(removeLecturerCommand *command.RemoveLecturerCommand) (interface{}, error) {
	if err := removeLecturerCommand.ValidateCommand(); err != nil {
		return nil, application.ThrowError(application.ARG_ERROR, err.Error())
	}
	transactionContext, err := factory.CreateTransactionContext(nil)
	if err != nil {
		return nil, application.ThrowError(application.TRANSACTION_ERROR, err.Error())
	}
	if err := transactionContext.StartTransaction(); err != nil {
		return nil, application.ThrowError(application.TRANSACTION_ERROR, err.Error())
	}
	defer func() {
		transactionContext.RollbackTransaction()
	}()
	var lecturerRepository lecturer.LecturerRepository
	if value, err := factory.CreateLecturerRepository(map[string]interface{}{
		"transactionContext": transactionContext,
	}); err != nil {
		return nil, application.ThrowError(application.INTERNAL_SERVER_ERROR, err.Error())
	} else {
		lecturerRepository = value
	}
	lecturer, err := lecturerRepository.FindOne(map[string]interface{}{"lecturerId": removeLecturerCommand.LecturerId})
	if err != nil {
		return nil, application.ThrowError(application.INTERNAL_SERVER_ERROR, err.Error())
	}
	if lecturer == nil {
		return nil, application.ThrowError(application.RES_NO_FIND_ERROR, fmt.Sprintf("%s", string(removeLecturerCommand.LecturerId)))
	}
	if lecturer, err := lecturerRepository.Remove(lecturer); err != nil {
		return nil, application.ThrowError(application.INTERNAL_SERVER_ERROR, err.Error())
	} else {
		if err := transactionContext.CommitTransaction(); err != nil {
			return nil, application.ThrowError(application.TRANSACTION_ERROR, err.Error())
		}
		return lecturer, nil
	}
}

// 返回讲师列表
func (lecturerService *LecturerService) ListLecturer(listLecturerQuery *query.ListLecturerQuery) (interface{}, error) {
	if err := listLecturerQuery.ValidateQuery(); err != nil {
		return nil, application.ThrowError(application.ARG_ERROR, err.Error())
	}
	transactionContext, err := factory.CreateTransactionContext(nil)
	if err != nil {
		return nil, application.ThrowError(application.TRANSACTION_ERROR, err.Error())
	}
	if err := transactionContext.StartTransaction(); err != nil {
		return nil, application.ThrowError(application.TRANSACTION_ERROR, err.Error())
	}
	defer func() {
		transactionContext.RollbackTransaction()
	}()
	var lecturerRepository lecturer.LecturerRepository
	if value, err := factory.CreateLecturerRepository(map[string]interface{}{
		"transactionContext": transactionContext,
	}); err != nil {
		return nil, application.ThrowError(application.INTERNAL_SERVER_ERROR, err.Error())
	} else {
		lecturerRepository = value
	}
	if count, lecturers, err := lecturerRepository.Find(tool_funs.SimpleStructToMap(listLecturerQuery)); err != nil {
		return nil, application.ThrowError(application.INTERNAL_SERVER_ERROR, err.Error())
	} else {
		if err := transactionContext.CommitTransaction(); err != nil {
			return nil, application.ThrowError(application.TRANSACTION_ERROR, err.Error())
		}
		return map[string]interface{}{
			"count":     count,
			"lecturers": lecturers,
		}, nil
	}
}

func NewLecturerService(options map[string]interface{}) *LecturerService {
	newLecturerService := &LecturerService{}
	return newLecturerService
}
