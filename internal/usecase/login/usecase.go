package usecase_login

type Usecase struct {
	serviceUsers                   ServiceUsers
	serviceUserMails               ServiceUserMails
	queueConfirmationCodeRequested QueueConfirmationCodeRequested
}

func New(
	serviceUsers ServiceUsers,
	serviceUserMails ServiceUserMails,
	queueConfirmationCodeRequested QueueConfirmationCodeRequested,
) *Usecase {
	return &Usecase{
		serviceUsers:                   serviceUsers,
		serviceUserMails:               serviceUserMails,
		queueConfirmationCodeRequested: queueConfirmationCodeRequested,
	}
}
