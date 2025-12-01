package service_habits

type Service struct {
	storageHabits StorageHabits
}

func New(
	storageHabits StorageHabits,
) *Service {
	return &Service{
		storageHabits: storageHabits,
	}
}
