package main

import "github.com/hjoshi123/temporal-loan-app/cmd/api"

func main() {
	if err := api.Execute(); err != nil {
		panic(err)
	}
}
