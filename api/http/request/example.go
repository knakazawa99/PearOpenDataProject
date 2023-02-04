package request

type SampleGetRequest struct {
	ID      int    `param:"id"`
	Name    string `query:"name"`
	Remarks string `query:"remarks"`
}

type SamplePostRequest struct {
	ID      int    `param:"id"`
	Name    string `json:"name"`
	Remarks string `json:"remarks"`
}
