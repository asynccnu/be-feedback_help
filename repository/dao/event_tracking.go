package dao

import (
	"context"
	"github.com/asynccnu/be-feedback_help/domain"
	"gorm.io/gorm"
	"sync"
	"time"
)

const (
	MoreFeedback              = 0
	MoreFeedbackButton        = 1
	MoreFeedbackAcgroup       = 2
	More_feedback_search_skip = 3
)

var mu sync.Mutex

// 记录问题解决情况
func (dao *GormDao) NoteQuestion(ctx context.Context, q domain.Question) error {
	return dao.db.WithContext(ctx).Create(&q).Error
}

// 记录事件
func (dao *GormDao) NoteEventTracking(ctx context.Context, event domain.EventTracking) error {
	return dao.db.WithContext(ctx).Create(&event).Error
}

func (dao *GormDao) NoteMoreFeedbackSearch(ctx context.Context, search domain.EventSearchQuestion) error {
	return dao.db.WithContext(ctx).Create(&search).Error
}

func (dao *GormDao) NoteMoreFeedbackSearchSkip(ctx context.Context, search domain.EventQuestion) error {
	return dao.db.WithContext(ctx).Create(&search).Error
}

func (dao *GormDao) NoteMoreFeedbackQA(ctx context.Context, QuestionId int64) error {
	mu.Lock()
	defer mu.Unlock()

	// 更新库存数量
	if err := dao.db.WithContext(ctx).
		Model(&domain.FrequentlyAskedQuestion{}).
		Where("ID = ?", QuestionId).
		Update("Click_Times", gorm.Expr("Click_Times + 1")).Error; err != nil {
		return err
	}

	return nil
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
