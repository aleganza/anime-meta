package tvdb

// === AUTH ===

type Client struct {
	Token string
}

type LoginResponse struct {
	Status string
	Data   struct {
		Token string
	}
}

// === SERIES ===

type Alias struct {
	Language string
	Name     string
}

type MediaStatus struct {
	Id          uint
	Name        string
	RecordType  string
	KeepUpdated bool
}

type TvdbSeriesResponse struct {
	Status string
	Data   struct {
		Id                   uint
		Name                 string
		Slug                 string
		Image                string
		NameTranslations     []string
		OverviewTranslations []string
		Aliases              []Alias

		FirstAired string
		LastAired  string
		NextAired  string

		Score uint

		Status MediaStatus

		OriginalCountry  string
		OriginalLanguage string

		AverageRuntime uint

		Episodes []struct {
			Id             uint
			SeriesId       uint
			Name           string
			Aired          string
			Runtime        *uint
			Overview       string
			Image          string
			Number         uint
			SeasonNumber   uint
			AbsoluteNumber uint
			Year           string
		}

		Overview string
		Year     string
	}
}

type TvdbMovieResponse struct {
	Status string
	Data   struct {
		Id                   uint
		Name                 string
		Slug                 string
		Image                string
		NameTranslations     []string
		OverviewTranslations []string
		Aliases              []Alias

		Score   uint
		Runtime uint

		Status MediaStatus

		LastUpdated string
		Year        string

		Trailers []struct {
			Id       uint
			Name     string
			URL      string
			Language string
			Runtime  uint
		}

		Genres []struct {
			Id   uint
			Name string
			Slug string
		}

		Releases []struct {
			Country string
			Date    string
			Detail  *string
		}

		OriginalCountry  string
		OriginalLanguage string

		Studios []struct {
			Id   uint
			Name string
		}

		Budget    string
		BoxOffice string

		Overview string
	}
}
