package pong

func Pong(c chan string) {
	c <- "pong"
}