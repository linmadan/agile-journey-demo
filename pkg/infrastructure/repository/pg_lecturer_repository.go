package repository

import (
	"fmt"

	"github.com/go-pg/pg"
	"github.com/linmadan/agile-journey-demo/pkg/domain/lecturer"
	"github.com/linmadan/agile-journey-demo/pkg/infrastructure/pg/models"
	pgTransaction "github.com/linmadan/egglib-go/transaction/pg"
)

type LecturerRepository struct {
	transactionContext *pgTransaction.TransactionContext
}

func (repository *LecturerRepository) nextIdentify() (int64, error) {
	return 0, nil
}

func (repository *LecturerRepository) Save(lecturer *lecturer.Lecturer) (*lecturer.Lecturer, error) {
	tx := repository.transactionContext.PgTx
	if lecturer.Identify() == nil {
		_, err := repository.nextIdentify()
		if err != nil {
			return lecturer, err
		}
		if _, err := tx.QueryOne(
			pg.Scan(&lecturer.Id, &lecturer.Name, &lecturer.Introduction, &lecturer.Topic),
			"INSERT INTO lecturers (id, name, introduction, topic) VALUES (DEFAULT, ?, ?, ?) RETURNING id, name, introduction, topic",
			lecturer.Name, lecturer.Introduction, lecturer.Topic); err != nil {
			return lecturer, err
		}
	} else {
		if _, err := tx.QueryOne(
			pg.Scan(&lecturer.Name, &lecturer.Introduction, &lecturer.Topic),
			"UPDATE lecturers SET name=?, introduction=?, topic=? WHERE id=? RETURNING name, introduction, topic",
			lecturer.Name, lecturer.Introduction, lecturer.Topic, lecturer.Identify()); err != nil {
			return lecturer, err
		}
	}
	return lecturer, nil
}

func (repository *LecturerRepository) Remove(lecturer *lecturer.Lecturer) (*lecturer.Lecturer, error) {
	tx := repository.transactionContext.PgTx
	lecturerModel := new(models.Lecturer)
	lecturerModel.Id = lecturer.Identify().(int64)
	if _, err := tx.Model(lecturerModel).WherePK().Delete(); err != nil {
		return lecturer, err
	}
	return lecturer, nil
}

func (repository *LecturerRepository) FindOne(queryOptions map[string]interface{}) (*lecturer.Lecturer, error) {
	tx := repository.transactionContext.PgTx
	lecturerModel := new(models.Lecturer)
	query := tx.Model(lecturerModel)
	if lecturerId, ok := queryOptions["lecturerId"]; ok {
		query = query.Where("lecturer.id = ?", lecturerId)
	}
	if err := query.Limit(1).Select(); err != nil {
		return nil, err
	}
	if lecturerModel.Id == 0 {
		return nil, nil
	} else {
		return &lecturer.Lecturer{
			Id:           lecturerModel.Id,
			Name:         lecturerModel.Name,
			Introduction: lecturerModel.Introduction,
			Topic:        lecturerModel.Topic,
		}, nil
	}
}

func (repository *LecturerRepository) Find(queryOptions map[string]interface{}) (int64, []*lecturer.Lecturer, error) {
	tx := repository.transactionContext.PgTx
	var lecturerModels []*models.Lecturer
	var lecturers []*lecturer.Lecturer
	query := tx.Model(&lecturerModels)
	if offset, ok := queryOptions["offset"]; ok {
		offset := offset.(int)
		if offset > -1 {
			query = query.Offset(offset)
		}
	} else {
		query = query.Offset(0)
	}
	if limit, ok := queryOptions["limit"]; ok {
		limit := limit.(int)
		if limit > -1 {
			query = query.Limit(limit)
		}
	} else {
		query = query.Limit(20)
	}
	if count, err := query.Order("id DESC").SelectAndCount(); err != nil {
		return 0, nil, err
	} else {
		for _, lecturerModel := range lecturerModels {
			lecturers = append(lecturers, &lecturer.Lecturer{
				Id:           lecturerModel.Id,
				Name:         lecturerModel.Name,
				Introduction: lecturerModel.Introduction,
				Topic:        lecturerModel.Topic,
			})
		}
		return int64(count), lecturers, nil
	}
}

func NewLecturerRepository(transactionContext *pgTransaction.TransactionContext) (*LecturerRepository, error) {
	if transactionContext == nil {
		return nil, fmt.Errorf("transactionContext参数不能为nil")
	} else {
		return &LecturerRepository{
			transactionContext: transactionContext,
		}, nil
	}
}
