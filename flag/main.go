package main

import (
	"flag"
	"fmt"
)

func main() {
	host := flag.String("host", "localhost", "database localhost")
	username := flag.String("username", "root", "database username")
	password := flag.String("password", "root", "database password")
	port := flag.Int("port", 3306, "database port")

	flag.Parse()

	fmt.Println("host:", *host)
	fmt.Println("username:", *username)
	fmt.Println("password:", *password)
	fmt.Println("port:", *port)
}
