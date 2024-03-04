package main

import (
	"heteroflow/cmd/codeaid/app"
)

func main() {
	command := app.NewCodeAidCommand()
	app.Run(command)
}
