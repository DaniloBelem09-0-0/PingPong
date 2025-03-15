package ping

func Ping(c chan string) {
	c <- "ping"
}