package service_tasks

type Service struct {
	storageTasks StorageTasks
}

func New(storageTasks StorageTasks) *Service {
	return &Service{
		storageTasks: storageTasks,
	}
}
