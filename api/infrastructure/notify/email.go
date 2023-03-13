package notify

type EmailDTO struct {
	Email          string
	MessageContent string
}

type EmailSender interface {
	Send(token EmailDTO) error
}

type email struct {
}

func (e email) Send(email EmailDTO) error {
	//TODO implement me
	return nil
}
