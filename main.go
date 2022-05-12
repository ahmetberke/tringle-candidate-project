package main

import (
	"github.com/ahmetberke/tringle-candidate-project/configs"
	"github.com/ahmetberke/tringle-candidate-project/internal/api"
)

func init() {
	configs.Manager.Setup()
}

func main() {
	app := api.NewAPI()
	err := app.Run()
	if err != nil {
		panic(err)
	}
}
