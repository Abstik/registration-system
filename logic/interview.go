package logic

import (
	"registration_system/dao/mysql"
	"registration_system/models"
)

/*// 选取面试时间
func ChooseInterviewTime(p *models.InterviewTimeParam, userID int64) error {
	interview := &models.Interview{
		InterviewTime: p.InterviewTime,
		Status:        0,
		UserID:        uint(userID),
	}

	return mysql.CreateInterview(interview)
}*/

// 更新面试时间
func UpdateInterviewTime(p *models.InterviewTimeParam, userID int64) error {
	interview := &models.Interview{
		InterviewTime: p.InterviewTime,
		UserID:        uint(userID),
	}

	return mysql.UpdateInterviewTime(interview)
}

func GetInterviewInfo(userID int64) (*models.Interview, error) {
	return mysql.GetInterviewInfo(userID)
}
