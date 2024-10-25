package main

import (
	"github.com/Redvin-dt/Lists-for-all-life-cases.git/internal/setup"
)

func main() {
	router := setup.SetupRouters()

	router.Run(":8080")
}
