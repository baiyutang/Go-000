package main

import (
	"fmt"
	"log"
)

func main() {
	defer func() {
		if e := recover(); e != nil {
			log.Fatal("there are somethins wrong, ", e)
		}
	}()

	id := 1
	service := new(UserService)
	user, err := service.GetNameById(id)
	if err != nil {
		fmt.Print(user)
	}
}
