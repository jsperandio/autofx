package example

type UserDB struct{}

func NewUserDB() *UserDB {
	return &UserDB{}
}

func (u *UserDB) Get(id string) (string, error) {
	return "John Doe", nil
}

func (u *UserDB) GetAll() ([]string, error) {
	return []string{"John Doe", "Jane Doe"}, nil
}

func (u *UserDB) Set(id string, value string) error {
	return nil
}
