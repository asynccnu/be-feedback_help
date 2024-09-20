package cache

import (
	"context"
	"encoding/json"
	"github.com/asynccnu/be-feedback_help/domain"
	"time"
)

func (cache *RedisCache) Set(ctx context.Context, question []domain.FrequentlyAskedQuestion) error {
	val, err := json.Marshal(question)
	if err != nil {
		return err
	}
	return cache.cmd.Set(ctx, "questions", val, time.Minute*30).Err()
}

func (cache *RedisCache) Get(ctx context.Context) ([]domain.FrequentlyAskedQuestion, error) {
	val, err := cache.cmd.Get(ctx, "questions").Bytes()
	if err != nil {
		return []domain.FrequentlyAskedQuestion{}, err
	}
	var q []domain.FrequentlyAskedQuestion
	err = json.Unmarshal(val, &q)
	return q, err
}
