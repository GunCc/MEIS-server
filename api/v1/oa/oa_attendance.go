package oa

import (
	"MEIS-server/global"
	commenReq "MEIS-server/model/commen/request"
	"MEIS-server/model/commen/response"
	oaModel "MEIS-server/model/oa"
	"MEIS-server/utils/upload"
	"fmt"

	"github.com/gin-gonic/gin"
	"github.com/tealeg/xlsx"
	"go.uber.org/zap"
)

type AttendanceApi struct {
}

// 获取考勤列表
func (u *AttendanceApi) GetAttendanceList(ctx *gin.Context) {
	var info commenReq.ListInfo

	err := ctx.ShouldBindJSON(&info)
	if err != nil {
		global.MEIS_LOGGER.Error("获取考勤列表参数错误", zap.Error(err))
		response.FailWithMessage("获取考勤列表参数错误", ctx)
		return
	}

	list, total, err := AttendanceController.GetAttendanceList(info)
	if err != nil {
		global.MEIS_LOGGER.Error("获取考勤列表参数错误", zap.Error(err))
		response.FailWithMessage("获取考勤列表参数错误", ctx)
		return
	}
	response.SuccessWithDetailed(response.ListRes{
		List:     list,
		Total:    total,
		Page:     info.Page,
		PageSize: info.PageSize,
	}, "数据获取成功", ctx)
}

// 获取考勤信息
func (u *AttendanceApi) GetAttendanceInfo(ctx *gin.Context) {
	var info commenReq.GetById
	err := ctx.ShouldBindJSON(&info)
	if err != nil {
		global.MEIS_LOGGER.Error("获取考勤参数错误", zap.Error(err))
		response.FailWithMessage("获取考勤参数错误", ctx)
		return
	}

	user, err := AttendanceController.GetAttendanceInfo(info.ID)
	if err != nil {
		global.MEIS_LOGGER.Error("获取考勤参数错误", zap.Error(err))
		response.FailWithMessage("获取考勤参数错误", ctx)
		return
	}
	response.SuccessWithDetailed(user, "数据获取成功", ctx)
}

// 删除某个考勤
func (u *AttendanceApi) RemoveAttendance(ctx *gin.Context) {
	var info oaModel.OAAttendance
	err := ctx.ShouldBindJSON(&info)
	if err != nil {
		global.MEIS_LOGGER.Error("获取考勤参数错误", zap.Error(err))
		response.FailWithMessage("获取考勤参数错误", ctx)
		return
	}

	err = AttendanceController.RemoveAttendance(info)
	if err != nil {
		global.MEIS_LOGGER.Error("获取考勤参数错误", zap.Error(err))
		response.FailWithMessage("获取考勤参数错误", ctx)
		return
	}
	response.SuccessWithMessage("删除成功", ctx)
}

// 新增某个考勤
func (u *AttendanceApi) CreateAttendance(ctx *gin.Context) {
	var info oaModel.OAAttendance

	err := ctx.ShouldBindJSON(&info)
	if err != nil {
		global.MEIS_LOGGER.Error("获取考勤参数错误", zap.Error(err))
		response.FailWithMessage("获取考勤参数错误", ctx)
		return
	}

	err = AttendanceController.CreateAttendance(info)
	if err != nil {
		global.MEIS_LOGGER.Error("修改考勤错误", zap.Error(err))
		response.FailWithMessage(err.Error(), ctx)
		return
	}
	response.SuccessWithMessage("修改成功", ctx)

}

// 修改某个考勤
func (u *AttendanceApi) UpdateAttendance(ctx *gin.Context) {
	var info oaModel.OAAttendance

	err := ctx.ShouldBindJSON(&info)
	if err != nil {
		global.MEIS_LOGGER.Error("获取考勤参数错误", zap.Error(err))
		response.FailWithMessage("获取考勤参数错误", ctx)
		return
	}

	err = AttendanceController.UpdateAttendance(info)
	if err != nil {
		global.MEIS_LOGGER.Error("修改考勤错误", zap.Error(err))
		response.FailWithMessage(err.Error(), ctx)
		return
	}
	response.SuccessWithMessage("修改成功", ctx)
}

// 导入考勤
func (u *AttendanceApi) Upload(ctx *gin.Context) {
	// 接收文件 参数为文件字段
	_, header, err := ctx.Request.FormFile("file")
	if err != nil {
		response.FailWithMessage("接收文件失败", ctx)
		return
	}

	oss := upload.NewOOS()
	filePath, _, uploadErr := oss.UploadFile(header)
	fmt.Println("file路径", filePath)
	if uploadErr != nil {
		global.MEIS_LOGGER.Error("上传excel失败!", zap.Error(err))
		response.FailWithMessage("上传excel失败", ctx)
		return
	}

	xlFile, err := xlsx.OpenFile(filePath)
	if err != nil {
		global.MEIS_LOGGER.Error("文件解析失败", zap.Error(err))
		response.FailWithMessage(err.Error(), ctx)
		return
	}

	// 处理 Excel 文件数据
	for _, sheet := range xlFile.Sheets {
		for i, row := range sheet.Rows {
			if i == 0 {
				continue
			}
			p_id, _ := row.Cells[1].Int()
			p_uint_id := uint(p_id)
			cells := oaModel.OAAttendance{
				PersonnelID: p_uint_id,
				WorkName:    row.Cells[0].String(),
				Work:        row.Cells[2].String(),
				Working:     row.Cells[3].String(),
				GrandTime:   row.Cells[4].String(),
				Comment:     row.Cells[5].String(),
			}
			err = AttendanceController.CreateAttendance(cells)
			if err != nil {
				global.MEIS_LOGGER.Error("文件解析时，数据不符合规范", zap.Error(err))
				response.FailWithMessage(err.Error(), ctx)
				return
			}
		}
	}

	response.SuccessWithDetailed(filePath, "上传成功", ctx)

}
