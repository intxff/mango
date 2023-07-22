package handlers

import (
	"context"
	"fmt"
	"time"

	rpc "github.com/intxff/mango/common/protos/users"
	"github.com/intxff/mango/users/internal/dao"
	"github.com/intxff/mango/users/internal/models"
	"go.uber.org/zap"
	"golang.org/x/crypto/bcrypt"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"google.golang.org/protobuf/types/known/timestamppb"
	"gorm.io/gen/field"
	"gorm.io/gorm"
)

// ensure server implemented
var _ rpc.UsersServer = (*UsersServer)(nil)

const (
	FOLLOW = "follow"
    BLOCK = "block"
)

type UsersServer struct {
	rpc.UnimplementedUsersServer
    zap.Logger
	ReadDB  *gorm.DB
	WriteDB *gorm.DB
}

func (u *UsersServer) Create(
	ctx context.Context, in *rpc.CreateReq,
) (*rpc.CreateResp, error) {
	hashedPassword, err := HashPassword(in.Password)
	if err != nil {
		errHashPassword := fmt.Errorf("[USERS]: failed to hash password %v", in.Password)
		u.Error(errHashPassword.Error(), zap.Error(err))
		return nil, status.Error(codes.InvalidArgument, errHashPassword.Error())
	}

	user := &models.User{
		Name:     in.Name,
		Password: hashedPassword,
		Email:    in.Email,
	}

	query := dao.Use(u.WriteDB)
	err = query.User.WithContext(ctx).Create(user)
	if err != nil {
		errCreate := fmt.Errorf("[USERS]: failed to create user %v", in.Name)
		u.Error(errCreate.Error(), zap.Error(err))
		return nil, status.Error(codes.Aborted, errCreate.Error())
	}
	return &rpc.CreateResp{}, nil
}

func (u *UsersServer) Delete(
	ctx context.Context, in *rpc.DeleteReq,
) (*rpc.DeleteResp, error) {
	query := dao.Use(u.WriteDB)

	_, err := query.User.WithContext(ctx).
		Where(query.User.UserID.Eq(in.UserId)).
		Delete()
	if err != nil {
		errDelete := fmt.Errorf("[USERS]: failed to delete user with user_id %v", in.UserId)
		u.Error(errDelete.Error(), zap.Error(err))
		return nil, status.Error(codes.Aborted, errDelete.Error())
	}

	return nil, nil
}

func (u *UsersServer) BatchDelete(
	ctx context.Context, in *rpc.BatchDeleteReq,
) (*rpc.DeleteResp, error) {
	query := dao.Use(u.WriteDB)

	_, err := query.User.WithContext(ctx).
		Where(query.User.UserID.In(in.UserId...)).
		Delete()
	if err != nil {
		errDelete := fmt.Errorf("[USERS]: failed to delete user in user_id %v", in.UserId)
		u.Error(errDelete.Error(), zap.Error(err))
		return nil, status.Error(codes.Aborted, errDelete.Error())
	}

	return nil, nil
}

func (u *UsersServer) Get(
	ctx context.Context, in *rpc.GetReq,
) (*rpc.GetResp, error) {
	query := dao.Use(u.ReadDB)
	user, err := query.User.WithContext(ctx).
		Where(query.User.UserID.Eq(in.UserId)).
		First()
	if err != nil {
		errGet := fmt.Errorf("[USERS]: failed to find user with user_id %v", in.UserId)
		u.Error(errGet.Error(), zap.Error(err))
		return nil, status.Error(codes.NotFound, errGet.Error())
	}

	out := &rpc.GetResp{}
	switch in.GetView() {
	case rpc.GetReq_UNSPECIFIED:
		fallthrough
	case rpc.GetReq_BASIC:
		out.UserId = user.UserID
		out.Name = user.Name
	case rpc.GetReq_SIMPLE:
		out.UserId = user.UserID
		out.Name = user.Name
		out.Slogan = &user.Slogan
		out.FollowersCount = &user.FollowersCount
		out.FollowingCount = &user.FollowingCount
		out.PostsCount = &user.PostsCount
		out.LikeCount = &user.LikeCount
		out.Online = &user.Online
	case rpc.GetReq_FULL:
		out.UserId = user.UserID
		out.Name = user.Name
		out.Slogan = &user.Slogan
		out.FollowersCount = &user.FollowersCount
		out.FollowingCount = &user.FollowingCount
		out.PostsCount = &user.PostsCount
		out.LikeCount = &user.LikeCount
		out.DislikeCount = &user.DislikeCount
		out.Online = &user.Online
		out.AccountStatus = &user.AccountStatus
		out.CreatedAt = timestamppb.New(user.CreatedAt)
		out.DeletedAt = timestamppb.New(user.DeletedAt.Time)
		out.UpdatedAt = timestamppb.New(user.UpdatedAt)
	}
	return out, nil
}

func (u *UsersServer) List(
	ctx context.Context, in *rpc.ListReq,
) (*rpc.ListResp, error) {
	query := dao.Use(u.ReadDB)
	OrderMap := map[string]field.Expr{
		"user_id":        query.User.UserID,
		"name":           query.User.Name,
		"email":          query.User.Email,
		"create_at":      query.User.CreatedAt,
		"followers":      query.User.FollowersCount,
		"account_status": query.User.AccountStatus,
	}

	usersModel, err := query.User.WithContext(ctx).
		Order(OrderMap[in.OrderBy]).
		Offset(int(in.PageToken)).
		Limit(int(in.PageSize)).
		Find()

	if err != nil {
		errList := fmt.Errorf("[USERS]: failed to list users")
		u.Error(errList.Error(), zap.Error(err))
		return nil, status.Error(codes.Aborted, errList.Error())
	}

	users := make([]*rpc.ListResp_User, 0, len(usersModel))
	for _, user := range usersModel {
		users = append(users, &rpc.ListResp_User{
			UserId:         user.UserID,
			Name:           user.Name,
			Email:          user.Email,
			Slogan:         user.Slogan,
			FollowersCount: user.FollowersCount,
			FollowingCount: user.FollowingCount,
			LikeCount:      user.LikeCount,
			DislikeCount:   user.DislikeCount,
			PostsCount:     user.PostsCount,
			Online:         user.Online,
			AccountStatus:  user.AccountStatus,
			CreatedAt:      timestamppb.New(user.CreatedAt),
			UpdatedAt:      timestamppb.New(user.UpdatedAt),
			DeletedAt:      timestamppb.New(user.DeletedAt.Time),
		})
	}

	out := &rpc.ListResp{
		Users:         users,
		NextPageToken: in.PageToken + in.PageSize,
	}

	if in.NeedTotalSize {
		count, err := query.User.WithContext(ctx).Count()
		if err != nil {
			errTotalSize := fmt.Errorf("[USERS]: failed to get total size")
			u.Error(errTotalSize.Error(), zap.Error(err))
			return nil, status.Error(codes.Aborted, errTotalSize.Error())
		}
		*out.TotalSize = uint64(count)
	}

	return out, nil
}

func (u *UsersServer) ListFollowers(
	ctx context.Context, in *rpc.ListFollowersReq,
) (*rpc.ListFollowersResp, error) {
	userQuery := dao.Use(u.ReadDB).User
	usersRelationQuery := dao.Use(u.ReadDB).UsersRelation

	// get followers
	var followers []struct {
		UserID    uint64
		Name      string
		Slogan    string
		CreatedAt time.Time
	}

	err := usersRelationQuery.WithContext(ctx).
		Select(
			userQuery.UserID,
			userQuery.Name,
			userQuery.Slogan,
			usersRelationQuery.CreatedAt,
		).
		Where(
			usersRelationQuery.RelatedUserID.Eq(in.UserId),
			usersRelationQuery.Relation.Eq(FOLLOW),
		).
		LeftJoin(
			userQuery,
			userQuery.UserID.EqCol(usersRelationQuery.UserID),
		).
		Order(usersRelationQuery.CreatedAt).
		Offset(int(in.PageToken)).
		Limit(int(in.PageSize)).
		Scan(&followers)

	if err != nil {
		errFollowers := fmt.Errorf("[USERS]: failed to get followers for user with user id %v", in.UserId)
		u.Error(errFollowers.Error(), zap.Error(err))
		return nil, status.Error(codes.Aborted, errFollowers.Error())
	}

	out := &rpc.ListFollowersResp{
		Followers:    make([]*rpc.ListFollowersResp_Follower, 0, len(followers)),
		NextPageToke: in.PageSize + in.PageToken,
	}

	for _, follower := range followers {
		out.Followers = append(out.Followers, &rpc.ListFollowersResp_Follower{
			UserId:    follower.UserID,
			Name:      follower.Name,
			Slogan:    follower.Slogan,
			CreatedAt: timestamppb.New(follower.CreatedAt),
		})
	}

	// get total size
	if in.NeedTotalSize {
		count, err := usersRelationQuery.WithContext(ctx).
			Select(
				userQuery.UserID,
				userQuery.Name,
				userQuery.Slogan,
				usersRelationQuery.CreatedAt,
			).
			Where(
				usersRelationQuery.UserID.Eq(in.UserId),
				usersRelationQuery.Relation.Eq(FOLLOW),
			).
			LeftJoin(
				userQuery,
				userQuery.UserID.EqCol(usersRelationQuery.RelatedUserID),
			).
			Count()

		if err != nil {
			errTotalSize := fmt.Errorf("[USERS]: failed to get total size of followers")
			u.Error(errTotalSize.Error(), zap.Error(err))
			return nil, status.Error(codes.Aborted, errTotalSize.Error())
		}

		*out.TotalSize = uint64(count)
	}

	return out, nil
}

func (u *UsersServer) UpdatePassword(
	ctx context.Context, in *rpc.UpdatePasswordReq,
) (*rpc.UpdatePasswordResp, error) {
	query := dao.Use(u.WriteDB).User

	hashedPassword, err := HashPassword(in.Password)
	if err != nil {
		errHashPassword := fmt.Errorf("[USERS]: failed to hash password %v", in.Password)
		u.Error(errHashPassword.Error(), zap.Error(err))
		return nil, status.Error(codes.InvalidArgument, err.Error())
	}

	_, err = query.WithContext(ctx).
		Where(query.UserID.Eq(in.UserId)).
		Update(query.Password, hashedPassword)
	if err != nil {
		errUpdate := fmt.Errorf("failed to update password for user with user id %v", in.UserId)
		u.Error(errUpdate.Error(), zap.Error(err))
		return nil, status.Error(codes.Aborted, errUpdate.Error())
	}

    return nil, nil
}
func (u *UsersServer) UpdateSlogan(
    ctx context.Context, in *rpc.UpdateSloganReq,
) (*rpc.UpdateSloganResp, error) {
	query := dao.Use(u.WriteDB).User

    _, err := query.WithContext(ctx).
		Where(query.UserID.Eq(in.UserId)).
		Update(query.Slogan, in.Slogan)
	if err != nil {
		errUpdate := fmt.Errorf("failed to update slogan for user with user id %v", in.UserId)
		u.Error(errUpdate.Error(), zap.Error(err))
		return nil, status.Error(codes.Aborted, errUpdate.Error())
	}

    return nil, nil
}
func (u *UsersServer) UpdateRole(
    ctx context.Context, in *rpc.UpdateRoleReq,
) (*rpc.UpdateRoleResp, error) {
	query := dao.Use(u.WriteDB).User

    _, err := query.WithContext(ctx).
		Where(query.UserID.Eq(in.UserId)).
		Update(query.Role, in.Role)
	if err != nil {
		errUpdate := fmt.Errorf("failed to update role for user with user id %v", in.UserId)
		u.Error(errUpdate.Error(), zap.Error(err))
		return nil, status.Error(codes.Aborted, errUpdate.Error())
	}

    return nil, nil
}

func (u *UsersServer) BatchUpdateRole(
    ctx context.Context, in *rpc.BatchUpdateRoleReq,
) (*rpc.BatchUpdateRoleResp, error) {
	query := dao.Use(u.WriteDB).User

    _, err := query.WithContext(ctx).
		Where(query.UserID.In(in.UsersId...)).
		Update(query.Role, in.Role)
	if err != nil {
		errUpdate := fmt.Errorf("failed to update role for user with user id %v", in.UsersId)
		u.Error(errUpdate.Error(), zap.Error(err))
		return nil, status.Error(codes.Aborted, errUpdate.Error())
	}

    return nil, nil
}

func (u *UsersServer) CreateFollowing(
    ctx context.Context, in *rpc.CreateFollowingReq,
) (*rpc.CreateFollowingResp, error) {
    query := dao.Use(u.WriteDB).UsersRelation

    userRelation := &models.UsersRelation{
        UserID: in.UserId,
        Relation: FOLLOW,
        RelatedUserID: in.FollowingUserId,
    }

    err := query.WithContext(ctx).Create(userRelation)
    if err != nil {
        errCreate := fmt.Errorf("[USERS]: failed to create relation %v following %v", in.UserId, in.FollowingUserId)
        u.Error(errCreate.Error(), zap.Error(err))
        return nil, status.Error(codes.Aborted, errCreate.Error())
    }

    return nil, nil
}

func (u *UsersServer) ListFollowing(
    ctx context.Context, in *rpc.ListFollowingReq,
) (*rpc.ListFollowingResp, error) {
	userQuery := dao.Use(u.ReadDB).User
	usersRelationQuery := dao.Use(u.ReadDB).UsersRelation

	// get following
    var followings []struct {
		UserID    uint64
		Name      string
		Slogan    string
		CreatedAt time.Time
	}

	err := usersRelationQuery.WithContext(ctx).
		Select(
			userQuery.UserID,
			userQuery.Name,
			userQuery.Slogan,
			usersRelationQuery.CreatedAt,
		).
		Where(
			usersRelationQuery.UserID.Eq(in.UserId),
			usersRelationQuery.Relation.Eq(FOLLOW),
		).
		LeftJoin(
			userQuery,
			userQuery.UserID.EqCol(usersRelationQuery.RelatedUserID),
		).
		Order(usersRelationQuery.CreatedAt).
		Offset(int(in.PageToken)).
		Limit(int(in.PageSize)).
		Scan(&followings)

	if err != nil {
		errFollowers := fmt.Errorf("[USERS]: failed to get followers for user with user id %v", in.UserId)
		u.Error(errFollowers.Error(), zap.Error(err))
		return nil, status.Error(codes.Aborted, errFollowers.Error())
	}

	out := &rpc.ListFollowingResp{
		Following:    make([]*rpc.ListFollowingResp_Following, 0, len(followings)),
		NextPageToke: in.PageSize + in.PageToken,
	}

	for _, following := range followings {
		out.Following = append(out.Following, &rpc.ListFollowingResp_Following{
			UserId:    following.UserID,
			Name:      following.Name,
			Slogan:    following.Slogan,
			CreatedAt: timestamppb.New(following.CreatedAt),
		})
	}

	// get total size
	if in.NeedTotalSize {
		count, err := usersRelationQuery.WithContext(ctx).
			Select(
				userQuery.UserID,
				userQuery.Name,
				userQuery.Slogan,
				usersRelationQuery.CreatedAt,
			).
			Where(
				usersRelationQuery.UserID.Eq(in.UserId),
				usersRelationQuery.Relation.Eq(FOLLOW),
			).
			LeftJoin(
				userQuery,
				userQuery.UserID.EqCol(usersRelationQuery.RelatedUserID),
			).
			Count()

		if err != nil {
			errTotalSize := fmt.Errorf("[USERS]: failed to get total size of following")
			u.Error(errTotalSize.Error(), zap.Error(err))
			return nil, status.Error(codes.Aborted, errTotalSize.Error())
		}

		*out.TotalSize = uint64(count)
	}

	return out, nil
}

func (u *UsersServer) DeleteFollowing(
    ctx context.Context, in *rpc.DeleteFollowingReq,
) (*rpc.DeleteFollowingResp, error) {
    query := dao.Use(u.WriteDB).UsersRelation

    _, err := query.WithContext(ctx).
        Where(
            query.UserID.Eq(in.UserId), 
            query.Relation.Eq(FOLLOW), 
            query.RelatedUserID.Eq(in.FollowingUserId),
        ).
        Delete()
    if err != nil {
        errDelete := fmt.Errorf("failed to delete relation %v following %v", in.UserId, in.FollowingUserId)
        u.Error(errDelete.Error(), zap.Error(err))
        return nil, status.Error(codes.Aborted, errDelete.Error())
    }
    return nil, nil
}

func (u *UsersServer) BatchDeleteFollowing(
    ctx context.Context, in *rpc.BatchDeleteFollowingReq,
) (*rpc.BatchDeleteFollowingResp, error) {
    query := dao.Use(u.WriteDB).UsersRelation

    _, err := query.WithContext(ctx).
        Where(
            query.UserID.Eq(in.UserId), 
            query.Relation.Eq(FOLLOW), 
            query.RelatedUserID.In(in.FollowingUsersId...),
        ).
        Delete()
    if err != nil {
        errDelete := fmt.Errorf("failed to delete relation %v following %v", in.UserId, in.FollowingUsersId)
        u.Error(errDelete.Error(), zap.Error(err))
        return nil, status.Error(codes.Aborted, errDelete.Error())
    }
    return nil, nil
}

func (u *UsersServer) CreateBlocked(
    ctx context.Context, in *rpc.CreateBlockedReq,
) (*rpc.CreateBlockedResp, error) {
    query := dao.Use(u.WriteDB).UsersRelation

    userRelation := &models.UsersRelation{
        UserID: in.UserId,
        Relation: BLOCK,
        RelatedUserID: in.BlockedUserId,
    }

    err := query.WithContext(ctx).Create(userRelation)
    if err != nil {
        errCreate := fmt.Errorf("[USERS]: failed to create relation %v blocking %v", in.UserId, in.BlockedUserId)
        u.Error(errCreate.Error(), zap.Error(err))
        return nil, status.Error(codes.Aborted, errCreate.Error())
    }

    return nil, nil
}

func (u *UsersServer) ListBlocked(
    ctx context.Context, in *rpc.ListBlockedReq,
) (*rpc.ListBlockedResp, error) {
	userQuery := dao.Use(u.ReadDB).User
	usersRelationQuery := dao.Use(u.ReadDB).UsersRelation

	// get following
    var blockedUsers []struct {
		UserID    uint64
		Name      string
		Slogan    string
		CreatedAt time.Time
	}

	err := usersRelationQuery.WithContext(ctx).
		Select(
			userQuery.UserID,
			userQuery.Name,
			userQuery.Slogan,
			usersRelationQuery.CreatedAt,
		).
		Where(
			usersRelationQuery.UserID.Eq(in.UserId),
			usersRelationQuery.Relation.Eq(BLOCK),
		).
		LeftJoin(
			userQuery,
			userQuery.UserID.EqCol(usersRelationQuery.RelatedUserID),
		).
		Order(usersRelationQuery.CreatedAt).
		Offset(int(in.PageToken)).
		Limit(int(in.PageSize)).
		Scan(&blockedUsers)

	if err != nil {
		errFollowers := fmt.Errorf("[USERS]: failed to get blocked users for user %v", in.UserId)
		u.Error(errFollowers.Error(), zap.Error(err))
		return nil, status.Error(codes.Aborted, errFollowers.Error())
	}

	out := &rpc.ListBlockedResp{
		Blocked:    make([]*rpc.ListBlockedResp_Blcoked, 0, len(blockedUsers)),
		NextPageToke: in.PageSize + in.PageToken,
	}

	for _, blockedUser := range blockedUsers {
		out.Blocked = append(out.Blocked, &rpc.ListBlockedResp_Blcoked{
			UserId:    blockedUser.UserID,
			Name:      blockedUser.Name,
			Slogan:    blockedUser.Slogan,
			CreatedAt: timestamppb.New(blockedUser.CreatedAt),
		})
	}

	// get total size
	if in.NeedTotalSize {
		count, err := usersRelationQuery.WithContext(ctx).
			Select(
				userQuery.UserID,
				userQuery.Name,
				userQuery.Slogan,
				usersRelationQuery.CreatedAt,
			).
			Where(
				usersRelationQuery.UserID.Eq(in.UserId),
				usersRelationQuery.Relation.Eq(BLOCK),
			).
			LeftJoin(
				userQuery,
				userQuery.UserID.EqCol(usersRelationQuery.RelatedUserID),
			).
			Count()

		if err != nil {
			errTotalSize := fmt.Errorf("[USERS]: failed to get total size of blocked users")
			u.Error(errTotalSize.Error(), zap.Error(err))
			return nil, status.Error(codes.Aborted, errTotalSize.Error())
		}

		*out.TotalSize = uint64(count)
	}

	return out, nil
}
func (u *UsersServer) DeleteBlocked(
    ctx context.Context, in *rpc.DeleteBlockedReq,
) (*rpc.DeleteBlockedResp, error) {
    query := dao.Use(u.WriteDB).UsersRelation

    _, err := query.WithContext(ctx).
        Where(
            query.UserID.Eq(in.UserId), 
            query.Relation.Eq(BLOCK), 
            query.RelatedUserID.Eq(in.BlockedUserId),
        ).
        Delete()
    if err != nil {
        errDelete := fmt.Errorf("failed to delete relation %v blocking %v", in.UserId, in.BlockedUserId)
        u.Error(errDelete.Error(), zap.Error(err))
        return nil, status.Error(codes.Aborted, errDelete.Error())
    }
    return nil, nil
}

func (u *UsersServer) BatchDeleteBlocked(
    ctx context.Context, in *rpc.BatchDeleteBlockedReq,
) (*rpc.BatchDeleteBlockedResp, error) {
    query := dao.Use(u.WriteDB).UsersRelation

    _, err := query.WithContext(ctx).
        Where(
            query.UserID.Eq(in.UserId), 
            query.Relation.Eq(FOLLOW), 
            query.RelatedUserID.In(in.BlockedUserId...),
        ).
        Delete()
    if err != nil {
        errDelete := fmt.Errorf("failed to delete relation %v blocking %v", in.UserId, in.BlockedUserId)
        u.Error(errDelete.Error(), zap.Error(err))
        return nil, status.Error(codes.Aborted, errDelete.Error())
    }
    return nil, nil
}

func HashPassword(password string) (string, error) {
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return "", err
	}
	return string(hashedPassword), nil
}

