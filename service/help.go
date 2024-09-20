package service

import (
	"context"
	"github.com/asynccnu/be-feedback_help/domain"
	"github.com/asynccnu/be-feedback_help/repository"
	"time"
)

type HelpSer struct {
	repo repository.HelpRepository
}

type Service interface {
	GetQuestions(ctx context.Context) ([]domain.FrequentlyAskedQuestion, error)
	FindQuestionByName(ctx context.Context, name string) ([]domain.FrequentlyAskedQuestion, error)
	CreateQuestion(ctx context.Context, q domain.FrequentlyAskedQuestion) error
	ChangeQuestion(ctx context.Context, q domain.FrequentlyAskedQuestion) error
	DeleteQuestion(ctx context.Context, q domain.FrequentlyAskedQuestion) error
	NoteQuestion(ctx context.Context, q domain.Question) error
	NoteEventTracking(ctx context.Context, event domain.EventTracking) error
	NoteMoreFeedbackSearch(ctx context.Context, search domain.EventSearchQuestion) error
	NoteMoreFeedbackSearchSkip(ctx context.Context, search domain.EventQuestion) error
}

func NewFeedbackHelpService(repo repository.HelpRepository) Service {
	return &HelpSer{repo: repo}
}

func (ser *HelpSer) GetQuestions(ctx context.Context) ([]domain.FrequentlyAskedQuestion, error) {
	return ser.repo.GetQuestions(ctx)
}
func (ser *HelpSer) FindQuestionByName(ctx context.Context, name string) ([]domain.FrequentlyAskedQuestion, error) {
	return ser.repo.FindQuestionByName(ctx, name)
}
func (ser *HelpSer) CreateQuestion(ctx context.Context, q domain.FrequentlyAskedQuestion) error {
	q.Utime = time.Now()
	q.Ctime = time.Now()
	return ser.repo.CreateQuestion(ctx, q)
}
func (ser *HelpSer) ChangeQuestion(ctx context.Context, q domain.FrequentlyAskedQuestion) error {
	q.Utime = time.Now()
	return ser.repo.ChangeQuestion(ctx, q)
}
func (ser *HelpSer) DeleteQuestion(ctx context.Context, q domain.FrequentlyAskedQuestion) error {
	return ser.repo.DeleteQuestion(ctx, q)
}
func (ser *HelpSer) NoteQuestion(ctx context.Context, q domain.Question) error {
	q.Ctime = time.Now()
	return ser.repo.NoteQuestion(ctx, q)
}
func (ser *HelpSer) NoteEventTracking(ctx context.Context, event domain.EventTracking) error {
	event.Ctime = time.Now()
	return ser.repo.NoteEventTracking(ctx, event)
}
func (ser *HelpSer) NoteMoreFeedbackSearch(ctx context.Context, search domain.EventSearchQuestion) error {
	search.Ctime = time.Now()
	return ser.repo.NoteMoreFeedbackSearch(ctx, search)
}
func (ser *HelpSer) NoteMoreFeedbackSearchSkip(ctx context.Context, search domain.EventQuestion) error {
	search.Ctime = time.Now()
	return ser.repo.NoteMoreFeedbackSearchSkip(ctx, search)
}