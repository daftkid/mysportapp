package user

type UserRepositoryDB struct {
}

func (s UserRepositoryDB) FindAll() ([]User, error) {
	return nil, nil
}

func NewUserRepositoryDB() UserRepositoryDB {
	return UserRepositoryDB{}
}
