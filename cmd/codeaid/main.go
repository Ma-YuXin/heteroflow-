package main

import (
	"heterflow/cmd/codeaid/app"
)

func main() {
	command := app.NewCodeAidCommand()
	app.Run(command)
}
