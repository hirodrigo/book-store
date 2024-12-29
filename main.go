package main

import (
	"bookstore/bootstrap"

	"go.uber.org/fx"
)

func main() {
	fx.New(bootstrap.Module).Run()
}
