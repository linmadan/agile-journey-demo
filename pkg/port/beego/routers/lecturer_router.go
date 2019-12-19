package routers

import (
	"github.com/astaxie/beego"
	"github.com/linmadan/agile-journey-demo/pkg/port/beego/controllers"
)

func init() {
	beego.Router("/lecturers/", &controllers.LecturerController{}, "Post:CreateLecturer")
	beego.Router("/lecturers/:lecturerId", &controllers.LecturerController{}, "Put:UpdateLecturer")
	beego.Router("/lecturers/:lecturerId", &controllers.LecturerController{}, "Get:GetLecturer")
	beego.Router("/lecturers/:lecturerId", &controllers.LecturerController{}, "Delete:RemoveLecturer")
	beego.Router("/lecturers/", &controllers.LecturerController{}, "Get:ListLecturer")
}
