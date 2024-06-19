package service

import (
	"content_manage/api/operate"
	"content_manage/internal/biz"
)

// AppService is a content service.
type AppService struct {
	operate.UnimplementedAppServer

	uc *biz.ContentUsecase
}

// NewAppService new a app service.
func NewAppService(uc *biz.ContentUsecase) *AppService {
	return &AppService{uc: uc}
}
