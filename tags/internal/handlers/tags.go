package handlers

import (
	"context"
	"fmt"

	"github.com/intxff/mango/common/log"
	rpc "github.com/intxff/mango/common/protos/tags"
	"github.com/intxff/mango/tags/internal/dao"
	"github.com/intxff/mango/tags/internal/models"
	"go.uber.org/zap"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"google.golang.org/protobuf/types/known/timestamppb"
	"gorm.io/gen/field"
	"gorm.io/gorm"
)

const (
	UpdateMaskName       = "name"
	UpdateMaskPostsCount = "posts_count"

	OrderByID        = "id"
	OrderByName      = "name"
	OrderByCreatedAt = "created_at"
)

var _ rpc.TagsServer = (*TagsServer)(nil)

type TagsServer struct {
	rpc.UnimplementedTagsServer
	log.Logger
	ReadDB  *gorm.DB
	WriteDB *gorm.DB
}

func (t *TagsServer) Create(
	ctx context.Context, in *rpc.CreateReq,
) (*rpc.CreateResp, error) {
	tagModel := &models.Tag{
		Name: in.Name,
	}
	query := dao.Use(t.WriteDB)
	err := query.Tag.WithContext(ctx).Create(tagModel)
	if err != nil {
		errCreate := fmt.Errorf("[Tags]: failed to create tag %v", in.Name)
		t.Error(errCreate.Error(), zap.Error(err))
		return &rpc.CreateResp{}, status.Error(codes.Aborted, errCreate.Error())
	}
	return &rpc.CreateResp{}, nil
}

func (t *TagsServer) Get(
	ctx context.Context, in *rpc.GetReq,
) (*rpc.GetResp, error) {
	query := dao.Use(t.ReadDB)
	var (
		tags []*models.Tag
		err  error
	)
	switch in.Index.(type) {
	case *rpc.GetReq_TagId:
		tags, err = query.Tag.WithContext(ctx).Where(query.Tag.TagID.Eq(in.GetTagId())).Find()
	case *rpc.GetReq_Name:
		tags, err = query.Tag.WithContext(ctx).Where(query.Tag.Name.Eq(in.GetName())).Find()
	}
	if err != nil {
		errGet := fmt.Errorf("[Tags]: failed to get tags with %v", in.Index)
		t.Error(errGet.Error(), zap.Error(err))
		return &rpc.GetResp{}, status.Error(codes.Aborted, errGet.Error())
	}

	tagsResp := make([]*rpc.GetResp_Tag, 0, len(tags))
	switch in.View {
	case rpc.GetReq_UNSPECIFIED:
		fallthrough
	case rpc.GetReq_BASIC:
		for _, tag := range tags {
			tagsResp = append(tagsResp, &rpc.GetResp_Tag{
				TagId: tag.TagID,
				Name:  tag.Name,
			})
		}
	case rpc.GetReq_FULL:
		for _, tag := range tags {
			tagsResp = append(tagsResp, &rpc.GetResp_Tag{
				TagId:      tag.TagID,
				Name:       tag.Name,
				PostsCount: &tag.PostsCount,
				CreatedAt:  timestamppb.New(tag.CreatedAt),
				DeletedAt:  timestamppb.New(tag.DeletedAt.Time),
				UpdatedAt:  timestamppb.New(tag.UpdatedAt),
			})
		}
	}

	return &rpc.GetResp{Tags: tagsResp}, nil
}

func (t *TagsServer) Update(
	ctx context.Context, in *rpc.UpdateReq,
) (*rpc.UpdateResp, error) {
	query := dao.Use(t.WriteDB)

	maskMap := map[string]struct {
		expr  field.Expr
		value any
	}{
		UpdateMaskName:       {query.Tag.Name, in.Name},
		UpdateMaskPostsCount: {query.Tag.PostsCount, in.PostsCount},
	}

	_, err := query.Tag.WithContext(ctx).
		Where(query.Tag.TagID.Eq(in.TagId)).
		Update(maskMap[in.UpdateMask].expr, maskMap[in.UpdateMask].value)

	if err != nil {
		errUpdate := fmt.Errorf("[Tags]: failed to update %v of tag %v", in.UpdateMask, in.TagId)
		t.Error(errUpdate.Error(), zap.Error(err))
		return nil, status.Error(codes.Aborted, errUpdate.Error())
	}

	return &rpc.UpdateResp{}, nil
}

func (t *TagsServer) List(
	ctx context.Context, in *rpc.ListReq,
) (*rpc.ListResp, error) {
	query := dao.Use(t.ReadDB)

	orderMap := map[string]field.Expr{
		OrderByID:        query.Tag.TagID,
		OrderByName:      query.Tag.Name,
		OrderByCreatedAt: query.Tag.CreatedAt,
	}

	modelTags, err := query.Tag.WithContext(ctx).
		Offset(int(in.PageToken)).
		Limit(int(in.PageSize)).
		Order(orderMap[in.OrderBy]).
		Find()
	if err != nil {
		errList := fmt.Errorf("[Tags]: failed to list all tags")
		t.Error(errList.Error(), err)
		return nil, status.Error(codes.Aborted, errList.Error())
	}
	respTags := make([]*rpc.ListResp_Tag, 0, len(modelTags))
	for _, tag := range modelTags {
		respTags = append(respTags, &rpc.ListResp_Tag{
			TagId:      tag.TagID,
			Name:       tag.Name,
			PostsCount: tag.PostsCount,
			CreatedAt:  timestamppb.New(tag.CreatedAt),
			DeletedAt:  timestamppb.New(tag.DeletedAt.Time),
			UpdatedAt:  timestamppb.New(tag.UpdatedAt),
		})
	}

    out := &rpc.ListResp{
		Tags:          respTags,
		NextPageToken: in.PageToken + in.PageSize,
	}
	if in.NeedTotalSize {
		count, err := query.Tag.WithContext(ctx).
			Count()
		if err != nil {
			errCount := fmt.Errorf("[Tags]: failed to count all tags")
			t.Error(errCount.Error(), err)
			return nil, status.Error(codes.Aborted, errCount.Error())
		}
        *out.TotalSize = uint64(count)
	}
	return out, nil
}
