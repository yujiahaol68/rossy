package socket

// Enable control socket status is open or not
var Enable = false

// Notificater will start when Listen(isBlock) and Push msg to User
type Notificater interface {
	Listen(c chan string)
	Push(msg string)
}
