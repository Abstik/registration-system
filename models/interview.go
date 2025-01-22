package models

type Interview struct {
	ID            uint   `gorm:"primaryKey;autoIncrement" json:"-"`
	UserID        uint   `gorm:"not null" json:"-"`
	InterviewTime string `gorm:"not null" json:"interviewTime"`
	Status        int8   `gorm:"not null;default:0" json:"status"` // 0表示等待一面，1表示一面未通过，2表示一面试已通过，3表示二面未通过，4表示二面已通过
}

/*type CustomTime time.Time

const (
	timeFormat = "01-02 15:00"
)

func (t CustomTime) UnmarshalJSON(data []byte) (err error) {
	now, err := time.ParseInLocation(`"`+timeFormat+`"`, string(data), time.Local)
	t = CustomTime(now)
	return
}

func (t CustomTime) MarshalJSON() ([]byte, error) {
	b := make([]byte, 0, len(timeFormat)+2)
	b = append(b, '"')
	b = time.Time(t).AppendFormat(b, timeFormat)
	b = append(b, '"')
	return b, nil
}

func (t CustomTime) String() string {
	return time.Time(t).Format(timeFormat)
}

// Equal 方法比较两个 CustomTime 类型的时间是否相等
func (t CustomTime) Equal(other CustomTime) bool {
	// 比较到分钟级别
	return t.String() == other.String()
}*/
