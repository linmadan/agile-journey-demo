package query

import (
	"fmt"

	"github.com/astaxie/beego/validation"
)

type GetLecturerQuery struct {
	// 讲师标识
	LecturerId int64 `json:"lecturerId" valid:"Required"`
}

func (getLecturerQuery *GetLecturerQuery) ValidateQuery() error {
	valid := validation.Validation{}
	b, err := valid.Valid(getLecturerQuery)
	if err != nil {
		return err
	}
	if !b {
		for _, validErr := range valid.Errors {
			return fmt.Errorf("%s  %s", validErr.Key, validErr.Message)
		}
	}
	return nil
}
