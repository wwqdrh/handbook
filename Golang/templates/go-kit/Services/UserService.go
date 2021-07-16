package Services

import "errors"

type IUserService interface {
	GetName(userId int) string
	DelUser(userId int) error
}

type UserService struct{}

func (s *UserService) GetName(userId int) string {
	if userId == 101 {
		return "admin"
	}
	return "guest"
}

func (s *UserService) DelUser(userId int) error {
	if userId == 101 {
		return errors.New("无权限")
	}
	return nil
}
