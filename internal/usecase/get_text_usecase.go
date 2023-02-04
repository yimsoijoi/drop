package usecase

import (
	"context"

	errorLibrary "github.com/yimsoijoi/drop/lib/error"
)

func (usecase usecase) GetText(ctx context.Context, key *string) (*string, error) {
	if key != nil {
		text, err := usecase.repo.GetText(ctx, key)
		if err != nil {
			return returnErrGetText(err)
		}
		return text, nil
	}
	return returnErrGetText(errorLibrary.ErrNotFoundInput)
}

func returnErrGetText(err error) (*string, error) {
	return nil, errorLibrary.NewError(errorLibrary.ErrUseCaseGetText.Error(), err)
}
