package tvdb

type Client struct {
	Token string
}

type LoginResponse struct {
	Status string `json:"status"`
	Data   struct {
		Token string `json:"token"`
	} `json:"data"`
}
