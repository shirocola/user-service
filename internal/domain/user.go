package domain

type User struct {
    ID    string
    Name  string
    Email string
}

type UserRepository interface {
    GetByID(id string) (*User, error)
    Create(user *User) error
}
