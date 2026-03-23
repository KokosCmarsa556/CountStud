package structerr

type Err struct {
	Message string `json:"error"`
	HasErr  bool   `json:"errBol"`
}

func NewErr(message string) *Err {
	return &Err{
		Message: message,
		HasErr:  true,
	}
}

func (e Err) Error() string {
	return e.Message
}
