package main

import (
	"anime-meta/lib/core/env"
	"anime-meta/lib/tvdb"
	"fmt"
)

func main() {
	env.Init()

	client, err := tvdb.Authorize()
	if err != nil {
		fmt.Errorf("%w", err)
		return
	}

	media, err := client.GetMedia("252322", "series")

	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Println(media)
}
