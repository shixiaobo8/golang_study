package main

import (
    "fmt"
)

type User struct {
    username string
    password string
}


func main (){
   // u := User{}
    u1 := new(User)
    u1.username = "bob"
    fmt.Println(u1)
}
