package command

import (
	"fmt"

	"github.com/astaxie/beego/validation"
)

type RemoveLecturerCommand struct {
	// 讲师标识
	LecturerId int64 `json:"lecturerId" valid:"Required"`
}

func (removeLecturerCommand *RemoveLecturerCommand) ValidateCommand() error {
	valid := validation.Validation{}
	b, err := valid.Valid(removeLecturerCommand)
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
