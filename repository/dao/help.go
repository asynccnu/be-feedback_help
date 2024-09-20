package dao

import (
	"context"
	"github.com/asynccnu/be-feedback_help/domain"
	"time"
)

// 获取常用问题
func (dao *GormDao) GetQuestions(ctx context.Context) ([]domain.FrequentlyAskedQuestion, error) {
	var FSQ []domain.FrequentlyAskedQuestion
	err := dao.db.WithContext(ctx).Model(&FrequentlyAskedQuestion{}).Find(&FSQ).Limit(10).Error
	return FSQ, err

}

// 搜索
func (dao *GormDao) FindQuestionsByName(ctx context.Context, name string) ([]domain.FrequentlyAskedQuestion, error) {
	var FSQ []domain.FrequentlyAskedQuestion
	err := dao.db.WithContext(ctx).Model(&FrequentlyAskedQuestion{}).Where("Question like ?", "%"+name+"%").Find(&FSQ).Error
	return FSQ, err
}

// 创建问题
func (dao *GormDao) CreateQuestion(ctx context.Context, q domain.FrequentlyAskedQuestion) error {
	return dao.db.WithContext(ctx).Create(&q).Error
}

// 更改问题
func (dao *GormDao) ChangeQuestion(ctx context.Context, q domain.FrequentlyAskedQuestion) error {
	return dao.db.WithContext(ctx).Model(&FrequentlyAskedQuestion{}).Where("Id = ?", q.Id).Updates(q).Error
}

// 删除问题
func (dao *GormDao) DeleteQuestion(ctx context.Context, q domain.FrequentlyAskedQuestion) error {
	return dao.db.WithContext(ctx).Model(&FrequentlyAskedQuestion{}).Delete(q).Error
}

// 常见问题
type FrequentlyAskedQuestion struct {
	Id         int64  `gorm:"primaryKey,autoIncrement"`
	Question   string `json:"question"`
	Answer     string `json:"answer"`
	Utime      time.Time
	Ctime      time.Time
	ClickTimes int //记录该问题点击次数  //More_feedback_Q&A
}
