package models

type Lecturer struct {
	TableName string `pg:"lecturers,alias:lecturer"`
	// 唯一标识ID
	Id int64 `pg:"pk"`
	// 讲师姓名
	Name string
	// 讲师个人简介
	Introduction string
	// 讲师分享的主题
	Topic string
}
