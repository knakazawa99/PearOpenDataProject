package request

type ReqeustEmail struct {
	Email string `json:"email"`
}

type TokenWithDownload struct {
	Email   string `param:"email" json:"email"`
	Token   string `param:"token" json:"token"`
	Version string `param:"version" json:"version"`
}
