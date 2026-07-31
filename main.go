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

	// series, err := client.FetchSeries("252322")
	// if err != nil {
	// 	fmt.Println(err)
	// 	return
	// }

	movie, err := client.FetchMovie("791")
	if err != nil {
		fmt.Println(err)
		return
	}

	// fmt.Println(series)
	fmt.Println(movie)
}
