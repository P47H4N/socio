package message

type MessageController struct {
	srv *MessageService
}

func NewController(srv *MessageService) *MessageController {
	return &MessageController{
		srv: srv,
	}
}

