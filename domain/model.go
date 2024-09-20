package domain

import "time"

type FrequentlyAskedQuestion struct {
	Id         int64  `gorm:"primaryKey,autoIncrement"`
	Question   string `json:"question"`
	Answer     string `json:"answer"`
	Utime      time.Time
	Ctime      time.Time
	ClickTimes int //记录该问题点击次数  //More_feedback_Q&A
}

type EventTracking struct {
	Id    int64 `gorm:"primaryKey,autoIncrement"`
	Ctime time.Time
	Event int8 `json:"event"`
}

// More_feedback_search
type EventSearchQuestion struct {
	Id       int64 `gorm:"primaryKey,autoIncrement"`
	Ctime    time.Time
	Question string `json:"question"`
}

// More_feedback_search_skip
type EventQuestion struct {
	Id         int64 `gorm:"primaryKey,autoIncrement"`
	Ctime      time.Time
	QuestionId int64
}

// 问题解决情况
type Question struct {
	Id         int64 `gorm:"primaryKey,autoIncrement"`
	QuestionId int64 `json:"question_id"`
	IfOver     bool  `json:"if_over"`
	Ctime      time.Time
}
