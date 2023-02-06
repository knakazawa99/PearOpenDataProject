package Notify

import "api/app/entity"

type TokenDTO struct {
	Token string
	Email entity.Email
}
