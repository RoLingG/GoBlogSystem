package main

import (
	"GoRoLingG/utils/pwd"
	"fmt"
)

func main() {
	fmt.Println(pwd.HashPwd("1234"))
	fmt.Println(pwd.CheckPwd("$2a$04$nRyvfsgAn/NMn/xcWx6J.eILkDWf88P/7lZRZXrGFzBNrQXqJcGFm", "1234"))
}
