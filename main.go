package main

import (
	"./cron"
	"./metrics"
)

func main() {
	cron.Start()
	metrics.Start()
}
