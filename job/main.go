package job

type PlaidJobRequest struct {
	Name   string `json:"name"`
	Email  string `json:"email"`
	Resume string `json:"resume"`
	Github string `json:"github"`
}

func New(name, email, resume, github string) PlaidJobRequest {
	return PlaidJobRequest{
		Name:   name,
		Email:  email,
		Resume: resume,
		Github: github,
	}
}
