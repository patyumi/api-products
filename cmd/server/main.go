package main

import "github.com/patyumi/api-products/configs"

func main() {
	config, _ := configs.LoadConfig(".")
	println(config.DBDriver)

}
