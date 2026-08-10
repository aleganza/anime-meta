package main

import (
	"anime-meta/lib/core/env"
	"anime-meta/lib/media"
	"anime-meta/lib/tvdb"
	"fmt"
)

func main() {
	env.Init()

	tvdbClient, err := tvdb.Authorize()
	if err != nil {
		fmt.Errorf("%w", err)
		return
	}

	series := media.FetchSeries(&tvdbClient, "252322")

	// movie, err := client.FetchMovie("791")
	// if err != nil {
	// 	fmt.Println(err)
	// 	return
	// }

	fmt.Println(series)
	// fmt.Println(movie)
}
