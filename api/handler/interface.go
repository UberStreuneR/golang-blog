package handler

type Handler interface {
	GetUser()
	GetUsers()
	AddUser()

	GetPost()
	GetPosts()
	AddPost()

	SignIn()
	SignUp()
}
