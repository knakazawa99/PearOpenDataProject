package request

type PearUpdate struct {
	ID             int    `param:"id"`
	ReleaseNote    string `json:"release_note"`
	ReleaseComment string `json:"release_comment"`
	ReleaseFlag    bool   `json:"release_flag"`
}
