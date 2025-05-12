package credentials

type Credentials struct {
	id        string
	email     string
	hpassword string
}

type CredentialsRepository interface {
	CreateCredentials(credentials Credentials) (int, error)
	ReadCredentials(int) (Credentials, error)
	UpdateCredentials(credentials Credentials) (int, error)
	DeleteCredentials(int) error
}
