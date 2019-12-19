package command

import (
	"fmt"

	"github.com/astaxie/beego/validation"
)

type UpdateLecturerCommand struct {
	// 唯一标识ID
	LecturerId int64 `json:"id,omitempty"`
	// 讲师个人简介
	Introduction string `json:"introduction,omitempty"`
	// 讲师分享的主题
	Topic string `json:"topic,omitempty"`
}

func (updateLecturerCommand *UpdateLecturerCommand) ValidateCommand() error {
	valid := validation.Validation{}
	b, err := valid.Valid(updateLecturerCommand)
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
