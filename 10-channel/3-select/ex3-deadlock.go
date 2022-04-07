package main

func main() {
	ch := make(chan string)
	select {
	case <-ch:
	}
}

//https://go.dev/play/p/TpYK9GEUQYJ