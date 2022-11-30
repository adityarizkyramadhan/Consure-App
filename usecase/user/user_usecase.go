package user

type UserUsecase interface {
	SignIn(string, string) (string, error)
	SignUp(string, string, string) (string, error)
}
