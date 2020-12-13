package main

import (
    gtu "golang_study/tools/utils"
    "fmt"
)

func init() {
    fmt.Printf("... main init ..."+"\n")
}

func main() {
    fmt.Println("test tools.."+gtu.Apollo_config_server_base_url)
}
