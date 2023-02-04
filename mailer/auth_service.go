package mailer

type AuthService interface {
	IsValidLogin(authCredentials string) bool
}
