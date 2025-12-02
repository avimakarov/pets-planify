package service_users

type Service struct {
	storage          Storage
	storageUsers     StorageUsers
	storageUserMails StorageUserMails
}

func New(
	storage Storage,
	storageUsers StorageUsers,
	storageUserMails StorageUserMails,
) *Service {
	return &Service{
		storage:          storage,
		storageUsers:     storageUsers,
		storageUserMails: storageUserMails,
	}
}
