package service

type Daily interface {
	GetDaily() error
	PublishDaily() error
	UpdateDaily() error
	DeleteDaily() error
	GetComment() error
	CreateComment() error
	UpdateComment() error
	DeleteComment() error
	Like() error
}

type DailyService struct {
}

func NewDailyService() *DailyService {
	return &DailyService{}
}

func (d *DailyService) GetDaily() error {
	// 查找所有日常
	// 分页

	return nil
}

func (d *DailyService) PublishDaily() error {
	// 发布日常

	return nil
}

func (d *DailyService) UpdateDaily() error {
	// 更改日常

	return nil
}

func (d *DailyService) DeleteDaily() error {
	// 删除日常

	return nil
}

func (d *DailyService) GetComment() error {
	// 获取评论

	return nil
}

func (d *DailyService) CreateComment() error {
	// 创建评论

	return nil
}

func (d *DailyService) UpdateComment() error {
	// 更新评论

	return nil
}

func (d *DailyService) DeleteComment() error {
	// 删除评论

	return nil
}

func (d *DailyService) Like() error {
	// 点赞(是否通用)

	return nil
}
