package usecase

import (
	"context"
	error2 "github.com/yimsoijoi/drop/lib/error"
)

func (usecase usecase) GetText(ctx context.Context, key *string) (*string, error) {
	if key != nil {
		text, err := usecase.repo.GetText(ctx, key)
		if err != nil {
			return nil, error2.NewError(error2.ErrGetText.Error(), err)
		}
		return text, nil
	}
	return nil, error2.NewError(error2.ErrNotFoundInput.Error(), error2.ErrNotFoundInput)
}
