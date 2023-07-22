package handlers

import (
	"context"
	"fmt"

	"github.com/intxff/mango/categories/internal/dao"
	"github.com/intxff/mango/categories/internal/models"
	"github.com/intxff/mango/common/log"
	rpc "github.com/intxff/mango/common/protos/categories"
	"go.uber.org/zap"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"google.golang.org/protobuf/types/known/timestamppb"
	"gorm.io/gen/field"
	"gorm.io/gorm"
)

const (
	UpdateMaskName        = "name"
	UpdateMaskPostsCount  = "posts_count"
	UpdateMaskSlogan      = "slogan"
	UpdateMaskDescription = "description"
	UpdateMaskParent      = "parent"

	OrderByID        = "id"
	OrderByName      = "name"
	OrderByCreatedAt = "created_at"
)

var _ rpc.CategoriesServer = (*CategoriesServer)(nil)

type CategoriesServer struct {
	rpc.UnimplementedCategoriesServer
	log.Logger
	ReadDB  *gorm.DB
	WriteDB *gorm.DB
}

func (t *CategoriesServer) Create(
	ctx context.Context, in *rpc.CreateReq,
) (*rpc.CreateResp, error) {
	categoryModel := &models.Category{
		Name:        in.Name,
		Slogan:      in.Slogan,
		Parent:      *in.Parent,
		Description: in.Description,
	}
	query := dao.Use(t.WriteDB)
	err := query.Category.WithContext(ctx).Create(categoryModel)
	if err != nil {
		errCreate := fmt.Errorf("[Categories]: failed to create category %v", in.Name)
		t.Error(errCreate.Error(), zap.Error(err))
		return &rpc.CreateResp{}, status.Error(codes.Aborted, errCreate.Error())
	}
	return &rpc.CreateResp{}, nil
}

func (t *CategoriesServer) Get(
	ctx context.Context, in *rpc.GetReq,
) (*rpc.GetResp, error) {
	query := dao.Use(t.ReadDB)
	var (
		categories []*models.Category
		err        error
	)
	switch in.Index.(type) {
	case *rpc.GetReq_CategoryId:
		categories, err = query.Category.WithContext(ctx).Where(query.Category.CategoryID.Eq(in.GetCategoryId())).Find()
	case *rpc.GetReq_Name:
		categories, err = query.Category.WithContext(ctx).Where(query.Category.Name.Eq(in.GetName())).Find()
	}
	if err != nil {
		errGet := fmt.Errorf("[Categories]: failed to get categories with %v", in.Index)
		t.Error(errGet.Error(), zap.Error(err))
		return &rpc.GetResp{}, status.Error(codes.Aborted, errGet.Error())
	}

	categoriesResp := make([]*rpc.GetResp_Category, 0, len(categories))
	switch in.View {
	case rpc.GetReq_UNSPECIFIED:
		fallthrough
	case rpc.GetReq_BASIC:
		for _, category := range categories {
			categoriesResp = append(categoriesResp, &rpc.GetResp_Category{
				CategoryId: category.CategoryID,
				Name:       category.Name,
			})
		}
	case rpc.GetReq_SIMPLE:
		for _, category := range categories {
			categoriesResp = append(categoriesResp, &rpc.GetResp_Category{
				CategoryId: category.CategoryID,
				Name:       category.Name,
				Slogan:     category.Slogan,
				Desciption: &category.Description,
			})
		}
	case rpc.GetReq_FULL:
		for _, category := range categories {
			categoriesResp = append(categoriesResp, &rpc.GetResp_Category{
				CategoryId: category.CategoryID,
				Name:       category.Name,
				PostsCount: &category.PostsCount,
				Slogan:     category.Slogan,
				Desciption: &category.Description,
				Parent:     &category.Parent,
				CreatedAt:  timestamppb.New(category.CreatedAt),
				DeletedAt:  timestamppb.New(category.DeletedAt.Time),
				UpdatedAt:  timestamppb.New(category.UpdatedAt),
			})
		}
	}

	return &rpc.GetResp{Categories: categoriesResp}, nil
}

func (t *CategoriesServer) Update(
	ctx context.Context, in *rpc.UpdateReq,
) (*rpc.UpdateResp, error) {
	query := dao.Use(t.WriteDB)

	maskMap := map[string]struct {
		expr  field.Expr
		value any
	}{
		UpdateMaskName:        {query.Category.Name, in.Name},
		UpdateMaskPostsCount:  {query.Category.PostsCount, in.PostsCount},
		UpdateMaskParent:      {query.Category.Parent, in.Parent},
		UpdateMaskSlogan:      {query.Category.Slogan, in.Slogan},
		UpdateMaskDescription: {query.Category.Description, in.Desciption},
	}

	_, err := query.Category.WithContext(ctx).
		Where(query.Category.CategoryID.Eq(in.CategoryId)).
		Update(maskMap[in.UpdateMask].expr, maskMap[in.UpdateMask].value)

	if err != nil {
		errUpdate := fmt.Errorf("[Categories]: failed to update %v of category %v", in.UpdateMask, in.CategoryId)
		t.Error(errUpdate.Error(), zap.Error(err))
		return nil, status.Error(codes.Aborted, errUpdate.Error())
	}

	return &rpc.UpdateResp{}, nil
}

func (t *CategoriesServer) List(
	ctx context.Context, in *rpc.ListReq,
) (*rpc.ListResp, error) {
	query := dao.Use(t.ReadDB)

	orderMap := map[string]field.Expr{
		OrderByID:        query.Category.CategoryID,
		OrderByName:      query.Category.Name,
		OrderByCreatedAt: query.Category.CreatedAt,
	}

	modelCategories, err := query.Category.WithContext(ctx).
		Offset(int(in.PageToken)).
		Limit(int(in.PageSize)).
		Order(orderMap[in.OrderBy]).
		Find()
	if err != nil {
		errList := fmt.Errorf("[Categories]: failed to list all categories")
		t.Error(errList.Error(), err)
		return nil, status.Error(codes.Aborted, errList.Error())
	}
	respCategories := make([]*rpc.ListResp_Category, 0, len(modelCategories))
	for _, category := range modelCategories {
		respCategories = append(respCategories, &rpc.ListResp_Category{
			CategoryId:  category.CategoryID,
			Name:        category.Name,
			PostsCount:  category.PostsCount,
			Parent:      category.Parent,
			Slogan:      category.Slogan,
			Description: category.Description,
			CreatedAt:   timestamppb.New(category.CreatedAt),
			DeletedAt:   timestamppb.New(category.DeletedAt.Time),
			UpdatedAt:   timestamppb.New(category.UpdatedAt),
		})
	}

	out := &rpc.ListResp{
		Categories:    respCategories,
		NextPageToken: in.PageToken + in.PageSize,
	}
	if in.NeedTotalSize {
		count, err := query.Category.WithContext(ctx).
			Count()
		if err != nil {
			errCount := fmt.Errorf("[Categories]: failed to count all categories")
			t.Error(errCount.Error(), err)
			return nil, status.Error(codes.Aborted, errCount.Error())
		}
		*out.TotalSize = uint64(count)
	}
	return out, nil
}
