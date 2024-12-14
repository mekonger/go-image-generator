package repo

type UserRepo interface {
	Hello(name string) string
	PostHello(name string, email string) string
}

func NewUserRepo() UserRepo {
	return &userRepo{}
}

type userRepo struct{}

func (ur *userRepo) Hello(name string) string {
	return "Hello, " + name
}

func (ur *userRepo) PostHello(name string, email string) string {
	return "Hello, " + name + " - your email is " + email
}
