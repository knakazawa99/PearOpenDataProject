package request

type ReqeustEmail struct {
	Email string `json:"notify"`
}

type TokenWithDownload struct {
	Email   string `json:"email"`
	Token   string `json:"token"`
	Version string `json:"version"`
}
