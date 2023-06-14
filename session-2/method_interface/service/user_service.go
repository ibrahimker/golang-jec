package service

type User struct {
	Name string
	age  int
}

type UserService struct {
	UsersDB []User
}

func NewUserService(db []User) *UserService {
	return &UserService{UsersDB: db}
}

func (u *UserService) Register(user User) {
	u.UsersDB = append(u.UsersDB, user)
}

func (u *UserService) Get() []User {
	return u.UsersDB
}
