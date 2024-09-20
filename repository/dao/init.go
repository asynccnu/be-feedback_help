package dao

import (
	"context"
	"github.com/asynccnu/be-feedback_help/domain"
	"gorm.io/gorm"
)

func InitTables(db *gorm.DB) error {
	db.AutoMigrate(&domain.FrequentlyAskedQuestion{})
	db.AutoMigrate(&domain.Question{})
	db.AutoMigrate(&domain.EventQuestion{})
	db.AutoMigrate(&domain.EventTracking{})
	db.AutoMigrate(&domain.EventSearchQuestion{})

	return db.Error
}

type Dao interface {
	GetQuestions(ctx context.Context) ([]domain.FrequentlyAskedQuestion, error)
	FindQuestionsByName(ctx context.Context, name string) ([]domain.FrequentlyAskedQuestion, error)
	CreateQuestion(ctx context.Context, q domain.FrequentlyAskedQuestion) error
	ChangeQuestion(ctx context.Context, q domain.FrequentlyAskedQuestion) error
	DeleteQuestion(ctx context.Context, q domain.FrequentlyAskedQuestion) error
	NoteQuestion(ctx context.Context, q domain.Question) error
	NoteEventTracking(ctx context.Context, event domain.EventTracking) error
	NoteMoreFeedbackSearch(ctx context.Context, search domain.EventSearchQuestion) error
	NoteMoreFeedbackSearchSkip(ctx context.Context, search domain.EventQuestion) error
	NoteMoreFeedbackQA(ctx context.Context, QuestionId int64) error
}

type GormDao struct {
	db *gorm.DB
}

func NewFeedbackHelpGormDao(db *gorm.DB) Dao {
	return &GormDao{db: db}
}
