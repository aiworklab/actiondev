package main

import (
	"github.com/aichy126/igo"
	"github.com/aiworklab/actiondev/api"
)

func main() {
	igo.App = igo.NewApp("")
	api.Router(igo.App.Web.Router)
	igo.App.Web.Run()
}
