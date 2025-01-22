package mysql

import "registration_system/models"

/*// 创建面试时间
func CreateInterview(interview *models.Interview) error {
	return DB.Create(interview).Error
}*/

// 更新面试时间
func UpdateInterviewTime(interview *models.Interview) error {
	existingInterview := new(models.Interview)
	result := DB.Model(&models.Interview{}).Where("user_id = ?", interview.UserID).First(existingInterview)
	// 如果未查询到数据（当前用户还未选择面试时间）
	if result.RowsAffected == 0 {
		// 新增当前用户的面试时间
		err := DB.Create(interview).Error
		return err
	}

	// 如果查询到数据，但用户当前面试时间与修改后相同，则不更新
	if existingInterview.InterviewTime == interview.InterviewTime {
		return nil
	}

	// 修改面试时间
	err := DB.Model(&models.Interview{}).Where("user_id = ?", interview.UserID).Update("interview_time", interview.InterviewTime).Error
	return err
}

// 获取面试信息
func GetInterviewInfo(userID int64) (*models.Interview, error) {
	interview := new(models.Interview)
	err := DB.Where("user_id = ?", userID).First(interview).Error
	return interview, err
}
