package controllers

import (
	"encoding/json"

	"github.com/astaxie/beego"
	"github.com/linmadan/agile-journey-demo/pkg/application/lecturer/command"
	"github.com/linmadan/agile-journey-demo/pkg/application/lecturer/query"
	"github.com/linmadan/agile-journey-demo/pkg/application/lecturer/service"
	"github.com/linmadan/egglib-go/web/beego/utils"
)

type LecturerController struct {
	beego.Controller
}

func (controller *LecturerController) CreateLecturer() {
	lecturerService := service.NewLecturerService(nil)
	createLecturerCommand := &command.CreateLecturerCommand{}
	json.Unmarshal(controller.Ctx.Input.GetData("requestBody").([]byte), createLecturerCommand)
	data, err := lecturerService.CreateLecturer(createLecturerCommand)
	var response utils.JsonResponse
	if err != nil {
		response = utils.ResponseError(controller.Ctx, err)
	} else {
		response = utils.ResponseData(controller.Ctx, data)
	}
	controller.Data["json"] = response
	controller.ServeJSON()
}

func (controller *LecturerController) UpdateLecturer() {
	lecturerService := service.NewLecturerService(nil)
	updateLecturerCommand := &command.UpdateLecturerCommand{}
	json.Unmarshal(controller.Ctx.Input.GetData("requestBody").([]byte), updateLecturerCommand)
	lecturerId, _ := controller.GetInt64(":lecturerId")
	updateLecturerCommand.LecturerId = lecturerId
	data, err := lecturerService.UpdateLecturer(updateLecturerCommand)
	var response utils.JsonResponse
	if err != nil {
		response = utils.ResponseError(controller.Ctx, err)
	} else {
		response = utils.ResponseData(controller.Ctx, data)
	}
	controller.Data["json"] = response
	controller.ServeJSON()
}

func (controller *LecturerController) GetLecturer() {
	lecturerService := service.NewLecturerService(nil)
	getLecturerQuery := &query.GetLecturerQuery{}
	lecturerId, _ := controller.GetInt64(":lecturerId")
	getLecturerQuery.LecturerId = lecturerId
	data, err := lecturerService.GetLecturer(getLecturerQuery)
	var response utils.JsonResponse
	if err != nil {
		response = utils.ResponseError(controller.Ctx, err)
	} else {
		response = utils.ResponseData(controller.Ctx, data)
	}
	controller.Data["json"] = response
	controller.ServeJSON()
}

func (controller *LecturerController) RemoveLecturer() {
	lecturerService := service.NewLecturerService(nil)
	removeLecturerCommand := &command.RemoveLecturerCommand{}
	lecturerId, _ := controller.GetInt64(":lecturerId")
	removeLecturerCommand.LecturerId = lecturerId
	data, err := lecturerService.RemoveLecturer(removeLecturerCommand)
	var response utils.JsonResponse
	if err != nil {
		response = utils.ResponseError(controller.Ctx, err)
	} else {
		response = utils.ResponseData(controller.Ctx, data)
	}
	controller.Data["json"] = response
	controller.ServeJSON()
}

func (controller *LecturerController) ListLecturer() {
	lecturerService := service.NewLecturerService(nil)
	listLecturerQuery := &query.ListLecturerQuery{}
	offset, _ := controller.GetInt("offset")
	listLecturerQuery.Offset = offset
	limit, _ := controller.GetInt("limit")
	listLecturerQuery.Limit = limit
	data, err := lecturerService.ListLecturer(listLecturerQuery)
	var response utils.JsonResponse
	if err != nil {
		response = utils.ResponseError(controller.Ctx, err)
	} else {
		response = utils.ResponseData(controller.Ctx, data)
	}
	controller.Data["json"] = response
	controller.ServeJSON()
}
