package domain

import "context"

type Usecase interface {
	CreateText(ctx context.Context, text *string) (code *string, err error)
	GetText(ctx context.Context, key *string) (*string, error)
}
