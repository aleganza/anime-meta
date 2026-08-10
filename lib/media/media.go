package media

import (
	"anime-meta/lib/core/threads"
	"anime-meta/lib/tvdb" 
)

func FetchSeries(tvdbClient *tvdb.Client, id string) any {
	results := threads.RunParallel(
		threads.ParallelFetchFunc{
			Name: "tvdb_extended",
			Func: func() (any, error) {
				return tvdbClient.FetchSeriesExtended(id)
			},
		},
		threads.ParallelFetchFunc{
			Name: "tvdb_episodes",
			Func: func() (any, error) {
				return tvdbClient.FetchSeriesEpisodes(id)
			},
		},
	)

	for _, result := range results {
		if result.Name == "tvdb_extended" {
			a := result.Body
		}
	}

	return "s"
}