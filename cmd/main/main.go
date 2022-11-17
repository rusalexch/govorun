package main

import (
	"govorun/internal/govorun"
)

func main() {
	gov := govorun.Init("Init word")

	gov.Start()
}
