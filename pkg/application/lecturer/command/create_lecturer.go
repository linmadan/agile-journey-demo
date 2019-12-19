package command

import (
	"fmt"

	"github.com/astaxie/beego/validation"
)

type CreateLecturerCommand struct {
	// 讲师姓名
	Name string `json:"name" valid:"Required"`
	// 讲师个人简介
	Introduction string `json:"introduction,omitempty"`
	// 讲师分享的主题
	Topic string `json:"topic,omitempty"`
}

func (createLecturerCommand *CreateLecturerCommand) ValidateCommand() error {
	valid := validation.Validation{}
	b, err := valid.Valid(createLecturerCommand)
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
