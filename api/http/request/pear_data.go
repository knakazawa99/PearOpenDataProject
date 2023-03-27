package request

type PearUpdate struct {
	ID             int    `param:"id"`
	ReleaseNote    string `json:"release_note"`
	ReleaseComment string `json:"release_comment"`
	ReleaseFlag    bool   `json:"release_flag"`
}

type PearCreate struct {
	// TODO: formnize
	ReleaseNote    string `json:"release_note" form:"release_note"`
	ReleaseComment string `json:"release_comment" form:"release_comment"`
	ReleaseFlag    bool   `json:"release_flag" form:"release_flag"`
	Version        string `json:"version" form:"version"`
}
