package biz

import (
	"context"
	"errors"
	"time"

	"github.com/go-kratos/kratos/v2/log"
	"golang.org/x/sync/errgroup"
)

// Content is a Content model.
type Content struct {
	ID             int64         `json:"id"`
	ContentID      string        `json:"content_id"`
	Title          string        `json:"title" binding:"required"`
	VideoURL       string        `json:"video_url" binding:"required"`
	Author         string        `json:"author" binding:"required"`
	Description    string        `json:"description"`
	Thumbnail      string        `json:"thumbnail"`
	Category       string        `json:"category"`
	Duration       time.Duration `json:"duration"`
	Resolution     string        `json:"resolution"`
	FileSize       int64         `json:"file_size"`
	Format         string        `json:"format"`
	Quality        int32         `json:"quality"`
	ApprovalStatus int32         `json:"approval_status"`
	UpdatedAt      time.Time     `json:"updated_at"`
	CreatedAt      time.Time     `json:"created_at"`
}

type FindParams struct {
	ID       int64
	Author   string
	Title    string
	Page     int32
	PageSize int32
}

type ContentIndex struct {
	ID        int64
	ContentID string
}

// ContentRepo is a Content repo.
type ContentRepo interface {
	Create(ctx context.Context, c *Content) (int64, error)
	Update(ctx context.Context, id int64, c *Content) error
	IsExist(ctx context.Context, id int64) (bool, error)
	Delete(ctx context.Context, id int64) error
	Find(ctx context.Context, params *FindParams) ([]*Content, int64, error)
	FindIndex(ctx context.Context, params *FindParams) ([]*ContentIndex, int64, error)
	First(ctx context.Context, idx *ContentIndex) (*Content, error)
}

// ContentUsecase is a Content usecase.
type ContentUsecase struct {
	repo ContentRepo
	log  *log.Helper
}

// NewContentUsecase new a Content usecase.
func NewContentUsecase(repo ContentRepo, logger log.Logger) *ContentUsecase {
	return &ContentUsecase{repo: repo, log: log.NewHelper(logger)}
}

// CreateContent creates a Content, and returns the new Content.
func (uc *ContentUsecase) CreateContent(ctx context.Context, c *Content) (int64, error) {
	uc.log.WithContext(ctx).Infof("CreateContent: %v", c)
	return uc.repo.Create(ctx, c)
}

// UpdateContent updates a Content, and returns the new Content.
func (uc *ContentUsecase) UpdateContent(ctx context.Context, c *Content) error {
	uc.log.WithContext(ctx).Infof("UpdateContent: %v", c)
	return uc.repo.Update(ctx, c.ID, c)
}

// DeleteContent deletes a Content by ID.
func (uc *ContentUsecase) DeleteContent(ctx context.Context, id int64) error {
	uc.log.WithContext(ctx).Infof("DeleteContent: %d", id)
	repo := uc.repo
	isExist, err := repo.IsExist(ctx, id)
	if err != nil {
		return err
	}
	if !isExist {
		return errors.New("内容不存在")
	}

	return uc.repo.Delete(ctx, id)
}

// FindContent finds a Content by ID.
func (uc *ContentUsecase) FindContent(ctx context.Context, params *FindParams) ([]*Content, int64, error) {
	uc.log.WithContext(ctx).Infof("FindContent: %v", params)
	repo := uc.repo
	indices, total, err := repo.FindIndex(ctx, params)
	if err != nil {
		return nil, 0, err
	}
	var eg errgroup.Group
	contents := make([]*Content, len(indices))
	for index, idx := range indices {
		tempIndex := index
		tempIdx := idx
		eg.Go(func() error {
			content, err := repo.First(ctx, tempIdx)
			if err != nil {
				return err
			}
			contents[tempIndex] = content
			return nil
		})
	}
	if err := eg.Wait(); err != nil {
		return nil, 0, err
	}
	return contents, total, nil
}
