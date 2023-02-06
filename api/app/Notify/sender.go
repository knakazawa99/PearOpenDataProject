package Notify

type Sender interface {
	Send(TokenDTO) error
}

type sender struct {
}

// TODO: Mock作って、usecaseの実装を進める
func (s sender) Send(dto TokenDTO) error {
	//TODO implement me
	panic("implement me")
}

func NewSender() Sender {
	return sender{}
}
