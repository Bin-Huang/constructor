package services

import (
	"log"

	"github.com/Bin-Huang/make-constructor/test/repositories"
)

//go:generate go run github.com/Bin-Huang/make-constructor@v0.1.0
type UserService struct {
	baseService

	userRepository *repositories.UserRepository
	proRepository  *repositories.ProRepository

	emailService *EmailService

	logger *log.Logger

	debug bool
}
