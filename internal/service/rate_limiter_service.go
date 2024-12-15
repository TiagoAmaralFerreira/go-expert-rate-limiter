package service

import (
	"time"

	"github.com/TiagoAmaralFerreira/go-expert-rate-limiter/internal/repository"
)

type RateLimiterService struct {
	repo                 *repository.RedisRepository
	maxRequestsPerSecond int
	blockDuration        time.Duration
}

func NewRateLimiterService(repo *repository.RedisRepository, maxReq int, blockDur time.Duration) *RateLimiterService {
	return &RateLimiterService{
		repo:                 repo,
		maxRequestsPerSecond: maxReq,
		blockDuration:        blockDur,
	}
}

func (s *RateLimiterService) IsRateLimited(key string) (bool, error) {
	currentCount, err := s.repo.Increment(key, s.blockDuration)
	if err != nil {
		return false, err
	}
	return currentCount > int64(s.maxRequestsPerSecond), nil
}
