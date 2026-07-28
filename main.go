package main

import (
	"anime-meta/lib/core/env"
	"anime-meta/lib/tvdb"
)

func main() {
	env.Init()

	tvdb.Authorize()

}
