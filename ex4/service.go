package main

type User struct {
	ID   int
	Name string
}

type Database interface {
	GetUserByID(userID int) (*User, error)
	SaveUser(user *User) error
	// Other methods...
}

type UserService struct {
	db Database
}

func NewUserService(db Database) *UserService {
	return &UserService{db: db}
}

func (us *UserService) GetUserByID(userID int) (*User, error) {
	return us.db.GetUserByID(userID)
}

// Use UserService in production with a real Database implementation.
// Use it in tests with a MockDatabase implementation.
