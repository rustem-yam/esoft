package main

import (
	"github.com/rustem-yam/esoft/internal/app"
)

// @title Esoft API
// @version 1.0
// @accept json
// @produce json

func main() {
	apl, err := app.New()
	if err != nil {
		panic(err)
	}

	if err = apl.Run(); err != nil {
		panic(err)
	}
}
