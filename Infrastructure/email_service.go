package infrastructure

import (
	"github.com/asaskevich/govalidator"
	"github.com/solo21-12/go-server-less-API/Domain/interfaces"
	"github.com/solo21-12/go-server-less-API/config"
)

type emailService struct {
	env config.Env
}

func NewEmailService(env config.Env) interfaces.EmailService {
	return &emailService{
		env: env,
	}
}

func (es *emailService) IsValidEmail(email string) bool {
	return govalidator.IsEmail(email)
}
