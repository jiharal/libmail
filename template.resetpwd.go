package libmail

import (
	"github.com/matcornic/hermes/v2"
)

// ResetPasswordModel is ...
type ResetPasswordModel struct {
	MimeType MimeType
	Hermes   hermes.Hermes
	Email    hermes.Email
}

// ResetPassword template reset password
func ResetPassword(opt ResetPasswordModel) (string, error) {
	hermes := opt.Hermes
	email := opt.Email
	switch opt.MimeType {
	case MimeTypeHTML:
		emailBody, err := hermes.GenerateHTML(email)
		if err != nil {
			return "", err
		}
		return emailBody, nil
	default:
		emailText, err := hermes.GeneratePlainText(email)
		if err != nil {
			return "", err
		}

		return emailText, nil
	}
}
