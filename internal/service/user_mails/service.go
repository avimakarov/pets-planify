package service_user_mails

type Service struct {
	storageUserMails StorageUserMails
}

func New(storageUserMails StorageUserMails) *Service {
	return &Service{
		storageUserMails: storageUserMails,
	}
}
