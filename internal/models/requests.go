package models

import "push-api/internal/pkg/tools"

type SendVerifyCodeReq struct {
	SendTo string `json:"send_to,omitempty"`
}

func (r SendVerifyCodeReq) Validate() error {
	if !tools.IsValidEmail(r.SendTo) {
		return ErrInvalidEmail
	}
	return nil
}
