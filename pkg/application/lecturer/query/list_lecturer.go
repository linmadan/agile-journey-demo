package query

import (
	"fmt"

	"github.com/astaxie/beego/validation"
)

type ListLecturerQuery struct {
	// 查询偏离量
	Offset int `json:"offset,omitempty"`
	// 查询限制
	Limit int `json:"limit,omitempty"`
}

func (listLecturerQuery *ListLecturerQuery) ValidateQuery() error {
	valid := validation.Validation{}
	b, err := valid.Valid(listLecturerQuery)
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
