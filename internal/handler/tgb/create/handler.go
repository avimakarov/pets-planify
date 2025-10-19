package handler_tgb_create

// todo нужна дополнительная абстракция как сделано с api
const (
	SlugCommandHandler = "create"
)

type Handler struct {
	serviceBot   ServiceBot
	serviceTasks ServiceTasks
}

func New(
	serviceBot ServiceBot,
	serviceTasks ServiceTasks,
) *Handler {
	return &Handler{
		serviceBot:   serviceBot,
		serviceTasks: serviceTasks,
	}
}
