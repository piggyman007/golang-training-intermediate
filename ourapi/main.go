package main

import (
	"log"

	"github.com/piggyman007/golang-training-intermediate/ourapi/routers"
)

func main() {
	e := routers.NewRouter()
	log.Fatal(e.Start(":1323"))
}
