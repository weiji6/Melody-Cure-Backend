package service

type Other interface {
	PrivacySetting() (string, error)
	AboutUs() (string, error)
	FeedBack() (string, error)
}

type OtherService struct {
}

func NewOtherService() *OtherService {
	return &OtherService{}
}

func (o *OtherService) PrivacySetting() (string, error) {
	return "", nil
}

func (o *OtherService) FeedBack() (string, error) {
	return "", nil
}

func (o *OtherService) AboutUs() (string, error) {
	return "", nil
}
