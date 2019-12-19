package lecturer

// 敏捷之旅的分享讲师
type Lecturer struct {
	// 唯一标识ID
	Id int64 `json:"id"`
	// 讲师姓名
	Name string `json:"name"`
	// 讲师个人简介
	Introduction string `json:"introduction"`
	// 讲师分享的主题
	Topic string `json:"topic"`
}

type LecturerRepository interface {
	Save(lecturer *Lecturer) (*Lecturer, error)
	Remove(lecturer *Lecturer) (*Lecturer, error)
	FindOne(queryOptions map[string]interface{}) (*Lecturer, error)
	Find(queryOptions map[string]interface{}) (int64, []*Lecturer, error)
}

func (lecturer *Lecturer) Identify() interface{} {
	if lecturer.Id == 0 {
		return nil
	}
	return lecturer.Id
}

func (lecturer *Lecturer) Update(data map[string]interface{}) error {
	if introduction, ok := data["introduction"]; ok {
		lecturer.Introduction = introduction.(string)
	}
	if topic, ok := data["topic"]; ok {
		lecturer.Topic = topic.(string)
	}
	return nil
}
