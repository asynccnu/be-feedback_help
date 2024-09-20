package cache

import (
	"context"
	"github.com/asynccnu/be-feedback_help/domain"
	"github.com/redis/go-redis/v9"
)

var ErrKeyNotExists = redis.Nil

type RedisCache struct {
	cmd redis.Cmdable
}

type Cache interface {
	Set(ctx context.Context, question []domain.FrequentlyAskedQuestion) error
	Get(ctx context.Context) ([]domain.FrequentlyAskedQuestion, error)
}

func NewFeedbackHelpRedisCache(cmd redis.Cmdable) Cache {
	return &RedisCache{cmd: cmd}
}
