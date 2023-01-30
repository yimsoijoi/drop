package usecase

import (
	domain2 "github.com/yimsoijoi/drop/internal/domain"
)

type usecase struct {
	repo domain2.RedisRepo
}

func NewUsecase(repo domain2.RedisRepo) domain2.Usecase {
	return usecase{repo: repo}
}
