package request

type ReqeustEmail struct {
	Organization string `json:"organization"`
	Name         string `json:"name"`
	Email        string `json:"email"`
}

type TokenWithDownload struct {
	Email   string `param:"email" json:"email"`
	Token   string `param:"token" json:"token"`
	Version string `param:"version" json:"version"`
}

type AdminAuth struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

type DeleteAdminAuth struct {
	ID int `param:"id"`
}
