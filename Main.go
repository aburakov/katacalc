package main

import (
	"Calculator/ServiceLogic"
)

func main() {
	instance := ServiceLogic.CalcReader{}
	instance.Read()
}
