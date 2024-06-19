package service

import (
	"content_manage/api/operate"
	"content_manage/internal/biz"
	"context"
)

func (a *AppService) FindContent(ctx context.Context, req *operate.FindContentReq) (*operate.FindContentRsp, error) {
	findParams := &biz.FindParams{
		ID:       req.GetId(),
		Author:   req.GetAuthor(),
		Title:    req.GetTitle(),
		Page:     req.GetPage(),
		PageSize: req.GetPageSize(),
	}
	uc := a.uc
	results, total, err := uc.FindContent(ctx, findParams)
	if err != nil {
		return nil, err
	}

	contents := make([]*operate.Content, 0, len(results))
	for _, result := range results {
		contents = append(contents, &operate.Content{
			Id:             result.ID,
			Title:          result.Title,
			VideoUrl:       result.VideoURL,
			Author:         result.Author,
			Description:    result.Description,
			Thumbnail:      result.Thumbnail,
			Category:       result.Category,
			Duration:       result.Duration.Microseconds(),
			Resolution:     result.Resolution,
			FileSize:       result.FileSize,
			Format:         result.Format,
			Quality:        result.Quality,
			ApprovalStatus: result.ApprovalStatus,
		})
	}

	return &operate.FindContentRsp{
		Contents: contents,
		Total:    total,
	}, nil
}
