package controller

import (
	"errors"
	"net/http"
	"regexp"

	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
	"gorm.io/gorm"

	"registration_system/logic"
	"registration_system/middleware"
	"registration_system/models"
	"registration_system/response"
)

/*// 选择面试时间
func ChooseInterviewTimeHandler(c *gin.Context) {
	// 获取userID
	userID, err := middleware.GetCurrentUserID(c)
	if err != nil {
		zap.L().Error("获取userID失败", zap.Error(err))
		// 如果用户未登录
		if errors.Is(err, models.ErrorNeedLogin) {
			response.ResponseError(c, models.CodeNeedLogin)
			return
		}
		response.ResponseError(c, models.CodeServerBusy)
	}

	// 获取参数
	p := new(models.InterviewTimeParam)
	if err = c.ShouldBindJSON(p); err != nil {
		zap.L().Error("参数校验失败", zap.Error(err))
		response.ResponseError(c, models.CodeInvalidParam)
	}

	// 新增面试时间
	err = logic.ChooseInterviewTime(p, userID)
	if err != nil {
		zap.L().Error("选择面试时间失败", zap.Error(err))
		response.ResponseError(c, models.CodeServerBusy)
	}
}*/

// 更新面试时间
func UpdateInterviewTimeHandler(c *gin.Context) {
	// 获取userID
	userID, err := middleware.GetCurrentUserID(c)
	if err != nil {
		response.ResponseError(c, http.StatusInternalServerError, models.CodeServerBusy)
		return
	}

	// 获取参数
	p := new(models.InterviewTimeParam)
	if err = c.ShouldBindJSON(p); err != nil {
		zap.L().Error("参数校验失败", zap.Error(err))
		response.ResponseError(c, http.StatusBadRequest, models.CodeInvalidParam)
		return
	}

	// 校验日期格式
	re := regexp.MustCompile(`^(0[1-9]|1[0-2])-(0[1-9]|[12][0-9]|3[01]) ([01]?[0-9]|2[0-3]):([0-5][0-9]) (AM|PM)$`)
	re.MatchString(p.InterviewTime)
	if !re.MatchString(p.InterviewTime) {
		zap.L().Error("日期格式错误", zap.Error(err))
		response.ResponseError(c, http.StatusBadRequest, models.CodeInvalidDateFormat+"，应为mm-dd hh-mm AM")
		return
	}

	// 更新面试时间
	err = logic.UpdateInterviewTime(p, userID)
	if err != nil {
		zap.L().Error("更新面试时间失败", zap.Error(err))
		response.ResponseError(c, http.StatusInternalServerError, models.CodeServerBusy)
		return
	}
	response.ResponseSuccess(c, nil)
	return
}

// 获取面试信息
func GetInterviewInfoHandler(c *gin.Context) {
	// 获取userID
	userID, err := middleware.GetCurrentUserID(c)
	if err != nil {
		zap.L().Error("获取userID失败", zap.Error(err))
		response.ResponseError(c, http.StatusInternalServerError, models.CodeServerBusy)
		return
	}

	// 获取面试信息
	info, err := logic.GetInterviewInfo(userID)
	if err != nil {
		zap.L().Error("获取面试信息失败", zap.Error(err))
		// 如果是未查询到数据（此用户未选择面试时间）
		if errors.Is(err, gorm.ErrRecordNotFound) {
			response.ResponseError(c, http.StatusBadRequest, models.CodeNotChooseInterviewTime)
			return
		}
		response.ResponseError(c, http.StatusInternalServerError, models.CodeServerBusy)
		return
	}
	response.ResponseSuccess(c, info)
	return
}
