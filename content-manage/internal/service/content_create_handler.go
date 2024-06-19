package service

import (
	"bytes"
	"content_manage/api/operate"
	"content_manage/internal/biz"
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"time"
)

func (a *AppService) CreateContent(ctx context.Context, req *operate.CreateContentReq) (*operate.CreateContentRsp, error) {
	content := req.GetContent()
	uc := a.uc
	id, err := uc.CreateContent(ctx, &biz.Content{
		Title:          content.GetTitle(),
		VideoURL:       content.GetVideoUrl(),
		Author:         content.GetAuthor(),
		Description:    content.GetDescription(),
		Thumbnail:      content.GetThumbnail(),
		Category:       content.GetCategory(),
		Duration:       time.Duration(content.GetDuration()),
		Resolution:     content.GetResolution(),
		FileSize:       content.GetFileSize(),
		Format:         content.GetFormat(),
		Quality:        content.GetQuality(),
		ApprovalStatus: content.GetApprovalStatus(),
	})
	if err != nil {
		return nil, err
	}
	// 执行工作流
	if err := a.ExecFlow(id); err != nil {
		return nil, err
	}

	return &operate.CreateContentRsp{}, nil
}

func (a *AppService) ExecFlow(id int64) error {
	url := "http://localhost:7788/flow/content-flow"
	method := "GET"

	payload := map[string]interface{}{
		"content_id": id,
	}
	data, _ := json.Marshal(payload)

	client := &http.Client{}
	req, err := http.NewRequest(method, url, bytes.NewReader(data))

	if err != nil {
		return err
	}
	req.Header.Add("Content-Type", "application/json")
	_, err = client.Do(req)
	if err != nil {
		fmt.Println(err)
		return err
	}

	return nil
}
