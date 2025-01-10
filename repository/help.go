package repository

import (
	"context"
	"github.com/asynccnu/be-feedback_help/domain"
	"github.com/asynccnu/be-feedback_help/pkg/logger"
	"github.com/asynccnu/be-feedback_help/repository/cache"
	"github.com/asynccnu/be-feedback_help/repository/dao"
	"time"
)

type HelpRepository interface {
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

type CachedRepository struct {
	dao   dao.Dao
	cache cache.Cache
	l     logger.Logger
}

func NewFeedbackHelpHelpRepository(dao dao.Dao, cache cache.Cache, l logger.Logger) HelpRepository {
	return &CachedRepository{dao: dao, cache: cache, l: l}
}

func (repo *CachedRepository) UpdateCache() error { //更新缓存，不对外暴露，在内部处理错误
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*2)
	defer cancel()
	cache, err := repo.dao.GetQuestions(ctx)
	if err != nil {
		repo.l.Error("缓存更新失败,获取问题失败", logger.Error(err))
		return err
	}
	err = repo.cache.Set(ctx, cache)
	if err != nil {
		repo.l.Error("缓存更新失败，存储问题失败", logger.Error(err))
		return err
	}
	return nil
}

func (repo *CachedRepository) GetQuestions(ctx context.Context) ([]domain.FrequentlyAskedQuestion, error) {
	res, err := repo.cache.Get(ctx)
	if err == nil {
		return res, nil
	}
	if err != cache.ErrKeyNotExists {
		// redis崩溃或者网络错误，用户量不大，MySQL撑得住，所以不降级处理
		repo.l.Error("访问Redis失败，查询常见问题缓存", logger.Error(err))
	}
	result, err := repo.dao.GetQuestions(ctx)
	// 异步回写
	go func() {
		ctx, cancel := context.WithTimeout(context.Background(), time.Second*1)
		defer cancel()
		er := repo.cache.Set(ctx, res)
		if er != nil {
			repo.l.Error("回写获取常见问题缓存失败", logger.Error(err))
		}
	}()
	return result, nil
}

func (repo *CachedRepository) FindQuestionByName(ctx context.Context, name string) ([]domain.FrequentlyAskedQuestion, error) {
	err := repo.dao.NoteMoreFeedbackSearch(ctx, domain.EventSearchQuestion{
		Ctime:    time.Now(),
		Question: name,
	})
	if err != nil {
		return nil, err
	}
	return repo.dao.FindQuestionsByName(ctx, name)
}

func (repo *CachedRepository) CreateQuestion(ctx context.Context, q domain.FrequentlyAskedQuestion) error {
	err := repo.dao.CreateQuestion(ctx, q)
	repo.UpdateCache()
	return err
}

func (repo *CachedRepository) ChangeQuestion(ctx context.Context, q domain.FrequentlyAskedQuestion) error {
	err := repo.dao.ChangeQuestion(ctx, q)
	repo.UpdateCache()
	return err
}

func (repo *CachedRepository) DeleteQuestion(ctx context.Context, q domain.FrequentlyAskedQuestion) error {
	err := repo.dao.DeleteQuestion(ctx, q)
	repo.UpdateCache()
	return err
}

func (repo *CachedRepository) NoteQuestion(ctx context.Context, q domain.Question) error {
	return repo.dao.NoteQuestion(ctx, q)
}

func (repo *CachedRepository) NoteEventTracking(ctx context.Context, event domain.EventTracking) error {
	return repo.dao.NoteEventTracking(ctx, event)
}

func (repo *CachedRepository) NoteMoreFeedbackSearch(ctx context.Context, search domain.EventSearchQuestion) error {
	return repo.dao.NoteMoreFeedbackSearch(ctx, search)
}

func (repo *CachedRepository) NoteMoreFeedbackSearchSkip(ctx context.Context, search domain.EventQuestion) error {
	err := repo.dao.NoteMoreFeedbackQA(ctx, search.QuestionId)
	if err != nil {
		return err
	}
	return repo.dao.NoteMoreFeedbackSearchSkip(ctx, search)
}
