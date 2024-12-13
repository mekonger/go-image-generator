package repo

type UserRepo interface {
	Hello(name string) string
}

func NewUserRepo() UserRepo {
	return &userRepo{}
}

type userRepo struct{}

func (ur *userRepo) Hello(name string) string {
	return "Hello, " + name
}
