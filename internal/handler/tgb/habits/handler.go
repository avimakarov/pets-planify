package handler_tgb_habits

type Handler struct {
	serviceBot    ServiceBot
	serviceHabits ServiceHabits
}

func New(
	serviceBot ServiceBot,
	serviceHabits ServiceHabits,
) *Handler {
	return &Handler{
		serviceBot:    serviceBot,
		serviceHabits: serviceHabits,
	}
}
