package biz

type UserRepo interface {
	CreateUser() error
}

type UserUseCase struct{}

func NewUserUseCase() *UserUseCase {
	return &UserUseCase{}
}

func (u *UserUseCase) UserRegister() {}
