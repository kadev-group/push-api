package interfaces

type IService interface {
	Mail() IMailService
}

type IMailService interface {
	SendVerificationCode(email, code string) error
}
