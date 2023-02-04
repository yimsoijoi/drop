package usecase

import (
	"context"
	"crypto/aes"
	"encoding/base64"
	"encoding/hex"

	errorLibrary "github.com/yimsoijoi/drop/lib/error"
)

func (usecase usecase) CreateText(ctx context.Context, text *string) (code *string, err error) {
	if text != nil || *text != "" {
		encodedText := base64.StdEncoding.EncodeToString([]byte(*text + "ThisIsMySecret16ThisIsMySecret16"))
		encryptedText := make([]byte, len(encodedText))
		//cipherBlock, err := aes.NewCipher([]byte(viper.GetString(env.AesCipherKey)))
		cipherBlock, err := aes.NewCipher([]byte("ThisIsMySecret16ThisIsMySecret16"))
		if err != nil {
			return nil, errorLibrary.NewError(errorLibrary.ErrCreateText.Error(), err)
		}
		cipherBlock.Encrypt(encryptedText, []byte(encodedText))
		hexStr := hex.EncodeToString(encryptedText)
		key := hexStr[0:4]
		if err := usecase.repo.SetText(ctx, key, *text); err != nil {
			return nil, errorLibrary.NewError(errorLibrary.ErrCreateText.Error(), err)
		}
		return &key, nil
	}
	return nil, errorLibrary.NewError(errorLibrary.ErrNotFoundInput.Error(), errorLibrary.ErrNotFoundInput)
}
