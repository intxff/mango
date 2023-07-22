// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.

package dao

import (
	"context"

	"gorm.io/gorm"
	"gorm.io/gorm/clause"
	"gorm.io/gorm/schema"

	"gorm.io/gen"
	"gorm.io/gen/field"

	"gorm.io/plugin/dbresolver"

	"github.com/intxff/mango/comments/internal/models"
)

func newComment(db *gorm.DB, opts ...gen.DOOption) comment {
	_comment := comment{}

	_comment.commentDo.UseDB(db, opts...)
	_comment.commentDo.UseModel(&models.Comment{})

	tableName := _comment.commentDo.TableName()
	_comment.ALL = field.NewAsterisk(tableName)
	_comment.CommentID = field.NewUint64(tableName, "comment_id")
	_comment.PostID = field.NewUint64(tableName, "post_id")
	_comment.SenderUserID = field.NewUint64(tableName, "sender_user_id")
	_comment.SenderUserIP = field.NewBytes(tableName, "sender_user_ip")
	_comment.ReceiverUserID = field.NewUint64(tableName, "receiver_user_id")
	_comment.Content = field.NewString(tableName, "content")
	_comment.LikeCount = field.NewInt32(tableName, "like_count")
	_comment.DislikeCount = field.NewInt32(tableName, "dislike_count")
	_comment.Parent = field.NewUint64(tableName, "parent")
	_comment.CreatedAt = field.NewTime(tableName, "created_at")
	_comment.DeletedAt = field.NewField(tableName, "deleted_at")
	_comment.UpdatedAt = field.NewTime(tableName, "updated_at")

	_comment.fillFieldMap()

	return _comment
}

type comment struct {
	commentDo commentDo

	ALL            field.Asterisk
	CommentID      field.Uint64
	PostID         field.Uint64
	SenderUserID   field.Uint64
	SenderUserIP   field.Bytes
	ReceiverUserID field.Uint64
	Content        field.String
	LikeCount      field.Int32
	DislikeCount   field.Int32
	Parent         field.Uint64
	CreatedAt      field.Time
	DeletedAt      field.Field
	UpdatedAt      field.Time

	fieldMap map[string]field.Expr
}

func (c comment) Table(newTableName string) *comment {
	c.commentDo.UseTable(newTableName)
	return c.updateTableName(newTableName)
}

func (c comment) As(alias string) *comment {
	c.commentDo.DO = *(c.commentDo.As(alias).(*gen.DO))
	return c.updateTableName(alias)
}

func (c *comment) updateTableName(table string) *comment {
	c.ALL = field.NewAsterisk(table)
	c.CommentID = field.NewUint64(table, "comment_id")
	c.PostID = field.NewUint64(table, "post_id")
	c.SenderUserID = field.NewUint64(table, "sender_user_id")
	c.SenderUserIP = field.NewBytes(table, "sender_user_ip")
	c.ReceiverUserID = field.NewUint64(table, "receiver_user_id")
	c.Content = field.NewString(table, "content")
	c.LikeCount = field.NewInt32(table, "like_count")
	c.DislikeCount = field.NewInt32(table, "dislike_count")
	c.Parent = field.NewUint64(table, "parent")
	c.CreatedAt = field.NewTime(table, "created_at")
	c.DeletedAt = field.NewField(table, "deleted_at")
	c.UpdatedAt = field.NewTime(table, "updated_at")

	c.fillFieldMap()

	return c
}

func (c *comment) WithContext(ctx context.Context) ICommentDo { return c.commentDo.WithContext(ctx) }

func (c comment) TableName() string { return c.commentDo.TableName() }

func (c comment) Alias() string { return c.commentDo.Alias() }

func (c *comment) GetFieldByName(fieldName string) (field.OrderExpr, bool) {
	_f, ok := c.fieldMap[fieldName]
	if !ok || _f == nil {
		return nil, false
	}
	_oe, ok := _f.(field.OrderExpr)
	return _oe, ok
}

func (c *comment) fillFieldMap() {
	c.fieldMap = make(map[string]field.Expr, 12)
	c.fieldMap["comment_id"] = c.CommentID
	c.fieldMap["post_id"] = c.PostID
	c.fieldMap["sender_user_id"] = c.SenderUserID
	c.fieldMap["sender_user_ip"] = c.SenderUserIP
	c.fieldMap["receiver_user_id"] = c.ReceiverUserID
	c.fieldMap["content"] = c.Content
	c.fieldMap["like_count"] = c.LikeCount
	c.fieldMap["dislike_count"] = c.DislikeCount
	c.fieldMap["parent"] = c.Parent
	c.fieldMap["created_at"] = c.CreatedAt
	c.fieldMap["deleted_at"] = c.DeletedAt
	c.fieldMap["updated_at"] = c.UpdatedAt
}

func (c comment) clone(db *gorm.DB) comment {
	c.commentDo.ReplaceConnPool(db.Statement.ConnPool)
	return c
}

func (c comment) replaceDB(db *gorm.DB) comment {
	c.commentDo.ReplaceDB(db)
	return c
}

type commentDo struct{ gen.DO }

type ICommentDo interface {
	gen.SubQuery
	Debug() ICommentDo
	WithContext(ctx context.Context) ICommentDo
	WithResult(fc func(tx gen.Dao)) gen.ResultInfo
	ReplaceDB(db *gorm.DB)
	ReadDB() ICommentDo
	WriteDB() ICommentDo
	As(alias string) gen.Dao
	Session(config *gorm.Session) ICommentDo
	Columns(cols ...field.Expr) gen.Columns
	Clauses(conds ...clause.Expression) ICommentDo
	Not(conds ...gen.Condition) ICommentDo
	Or(conds ...gen.Condition) ICommentDo
	Select(conds ...field.Expr) ICommentDo
	Where(conds ...gen.Condition) ICommentDo
	Order(conds ...field.Expr) ICommentDo
	Distinct(cols ...field.Expr) ICommentDo
	Omit(cols ...field.Expr) ICommentDo
	Join(table schema.Tabler, on ...field.Expr) ICommentDo
	LeftJoin(table schema.Tabler, on ...field.Expr) ICommentDo
	RightJoin(table schema.Tabler, on ...field.Expr) ICommentDo
	Group(cols ...field.Expr) ICommentDo
	Having(conds ...gen.Condition) ICommentDo
	Limit(limit int) ICommentDo
	Offset(offset int) ICommentDo
	Count() (count int64, err error)
	Scopes(funcs ...func(gen.Dao) gen.Dao) ICommentDo
	Unscoped() ICommentDo
	Create(values ...*models.Comment) error
	CreateInBatches(values []*models.Comment, batchSize int) error
	Save(values ...*models.Comment) error
	First() (*models.Comment, error)
	Take() (*models.Comment, error)
	Last() (*models.Comment, error)
	Find() ([]*models.Comment, error)
	FindInBatch(batchSize int, fc func(tx gen.Dao, batch int) error) (results []*models.Comment, err error)
	FindInBatches(result *[]*models.Comment, batchSize int, fc func(tx gen.Dao, batch int) error) error
	Pluck(column field.Expr, dest interface{}) error
	Delete(...*models.Comment) (info gen.ResultInfo, err error)
	Update(column field.Expr, value interface{}) (info gen.ResultInfo, err error)
	UpdateSimple(columns ...field.AssignExpr) (info gen.ResultInfo, err error)
	Updates(value interface{}) (info gen.ResultInfo, err error)
	UpdateColumn(column field.Expr, value interface{}) (info gen.ResultInfo, err error)
	UpdateColumnSimple(columns ...field.AssignExpr) (info gen.ResultInfo, err error)
	UpdateColumns(value interface{}) (info gen.ResultInfo, err error)
	UpdateFrom(q gen.SubQuery) gen.Dao
	Attrs(attrs ...field.AssignExpr) ICommentDo
	Assign(attrs ...field.AssignExpr) ICommentDo
	Joins(fields ...field.RelationField) ICommentDo
	Preload(fields ...field.RelationField) ICommentDo
	FirstOrInit() (*models.Comment, error)
	FirstOrCreate() (*models.Comment, error)
	FindByPage(offset int, limit int) (result []*models.Comment, count int64, err error)
	ScanByPage(result interface{}, offset int, limit int) (count int64, err error)
	Scan(result interface{}) (err error)
	Returning(value interface{}, columns ...string) ICommentDo
	UnderlyingDB() *gorm.DB
	schema.Tabler
}

func (c commentDo) Debug() ICommentDo {
	return c.withDO(c.DO.Debug())
}

func (c commentDo) WithContext(ctx context.Context) ICommentDo {
	return c.withDO(c.DO.WithContext(ctx))
}

func (c commentDo) ReadDB() ICommentDo {
	return c.Clauses(dbresolver.Read)
}

func (c commentDo) WriteDB() ICommentDo {
	return c.Clauses(dbresolver.Write)
}

func (c commentDo) Session(config *gorm.Session) ICommentDo {
	return c.withDO(c.DO.Session(config))
}

func (c commentDo) Clauses(conds ...clause.Expression) ICommentDo {
	return c.withDO(c.DO.Clauses(conds...))
}

func (c commentDo) Returning(value interface{}, columns ...string) ICommentDo {
	return c.withDO(c.DO.Returning(value, columns...))
}

func (c commentDo) Not(conds ...gen.Condition) ICommentDo {
	return c.withDO(c.DO.Not(conds...))
}

func (c commentDo) Or(conds ...gen.Condition) ICommentDo {
	return c.withDO(c.DO.Or(conds...))
}

func (c commentDo) Select(conds ...field.Expr) ICommentDo {
	return c.withDO(c.DO.Select(conds...))
}

func (c commentDo) Where(conds ...gen.Condition) ICommentDo {
	return c.withDO(c.DO.Where(conds...))
}

func (c commentDo) Exists(subquery interface{ UnderlyingDB() *gorm.DB }) ICommentDo {
	return c.Where(field.CompareSubQuery(field.ExistsOp, nil, subquery.UnderlyingDB()))
}

func (c commentDo) Order(conds ...field.Expr) ICommentDo {
	return c.withDO(c.DO.Order(conds...))
}

func (c commentDo) Distinct(cols ...field.Expr) ICommentDo {
	return c.withDO(c.DO.Distinct(cols...))
}

func (c commentDo) Omit(cols ...field.Expr) ICommentDo {
	return c.withDO(c.DO.Omit(cols...))
}

func (c commentDo) Join(table schema.Tabler, on ...field.Expr) ICommentDo {
	return c.withDO(c.DO.Join(table, on...))
}

func (c commentDo) LeftJoin(table schema.Tabler, on ...field.Expr) ICommentDo {
	return c.withDO(c.DO.LeftJoin(table, on...))
}

func (c commentDo) RightJoin(table schema.Tabler, on ...field.Expr) ICommentDo {
	return c.withDO(c.DO.RightJoin(table, on...))
}

func (c commentDo) Group(cols ...field.Expr) ICommentDo {
	return c.withDO(c.DO.Group(cols...))
}

func (c commentDo) Having(conds ...gen.Condition) ICommentDo {
	return c.withDO(c.DO.Having(conds...))
}

func (c commentDo) Limit(limit int) ICommentDo {
	return c.withDO(c.DO.Limit(limit))
}

func (c commentDo) Offset(offset int) ICommentDo {
	return c.withDO(c.DO.Offset(offset))
}

func (c commentDo) Scopes(funcs ...func(gen.Dao) gen.Dao) ICommentDo {
	return c.withDO(c.DO.Scopes(funcs...))
}

func (c commentDo) Unscoped() ICommentDo {
	return c.withDO(c.DO.Unscoped())
}

func (c commentDo) Create(values ...*models.Comment) error {
	if len(values) == 0 {
		return nil
	}
	return c.DO.Create(values)
}

func (c commentDo) CreateInBatches(values []*models.Comment, batchSize int) error {
	return c.DO.CreateInBatches(values, batchSize)
}

// Save : !!! underlying implementation is different with GORM
// The method is equivalent to executing the statement: db.Clauses(clause.OnConflict{UpdateAll: true}).Create(values)
func (c commentDo) Save(values ...*models.Comment) error {
	if len(values) == 0 {
		return nil
	}
	return c.DO.Save(values)
}

func (c commentDo) First() (*models.Comment, error) {
	if result, err := c.DO.First(); err != nil {
		return nil, err
	} else {
		return result.(*models.Comment), nil
	}
}

func (c commentDo) Take() (*models.Comment, error) {
	if result, err := c.DO.Take(); err != nil {
		return nil, err
	} else {
		return result.(*models.Comment), nil
	}
}

func (c commentDo) Last() (*models.Comment, error) {
	if result, err := c.DO.Last(); err != nil {
		return nil, err
	} else {
		return result.(*models.Comment), nil
	}
}

func (c commentDo) Find() ([]*models.Comment, error) {
	result, err := c.DO.Find()
	return result.([]*models.Comment), err
}

func (c commentDo) FindInBatch(batchSize int, fc func(tx gen.Dao, batch int) error) (results []*models.Comment, err error) {
	buf := make([]*models.Comment, 0, batchSize)
	err = c.DO.FindInBatches(&buf, batchSize, func(tx gen.Dao, batch int) error {
		defer func() { results = append(results, buf...) }()
		return fc(tx, batch)
	})
	return results, err
}

func (c commentDo) FindInBatches(result *[]*models.Comment, batchSize int, fc func(tx gen.Dao, batch int) error) error {
	return c.DO.FindInBatches(result, batchSize, fc)
}

func (c commentDo) Attrs(attrs ...field.AssignExpr) ICommentDo {
	return c.withDO(c.DO.Attrs(attrs...))
}

func (c commentDo) Assign(attrs ...field.AssignExpr) ICommentDo {
	return c.withDO(c.DO.Assign(attrs...))
}

func (c commentDo) Joins(fields ...field.RelationField) ICommentDo {
	for _, _f := range fields {
		c = *c.withDO(c.DO.Joins(_f))
	}
	return &c
}

func (c commentDo) Preload(fields ...field.RelationField) ICommentDo {
	for _, _f := range fields {
		c = *c.withDO(c.DO.Preload(_f))
	}
	return &c
}

func (c commentDo) FirstOrInit() (*models.Comment, error) {
	if result, err := c.DO.FirstOrInit(); err != nil {
		return nil, err
	} else {
		return result.(*models.Comment), nil
	}
}

func (c commentDo) FirstOrCreate() (*models.Comment, error) {
	if result, err := c.DO.FirstOrCreate(); err != nil {
		return nil, err
	} else {
		return result.(*models.Comment), nil
	}
}

func (c commentDo) FindByPage(offset int, limit int) (result []*models.Comment, count int64, err error) {
	result, err = c.Offset(offset).Limit(limit).Find()
	if err != nil {
		return
	}

	if size := len(result); 0 < limit && 0 < size && size < limit {
		count = int64(size + offset)
		return
	}

	count, err = c.Offset(-1).Limit(-1).Count()
	return
}

func (c commentDo) ScanByPage(result interface{}, offset int, limit int) (count int64, err error) {
	count, err = c.Count()
	if err != nil {
		return
	}

	err = c.Offset(offset).Limit(limit).Scan(result)
	return
}

func (c commentDo) Scan(result interface{}) (err error) {
	return c.DO.Scan(result)
}

func (c commentDo) Delete(models ...*models.Comment) (result gen.ResultInfo, err error) {
	return c.DO.Delete(models)
}

func (c *commentDo) withDO(do gen.Dao) *commentDo {
	c.DO = *do.(*gen.DO)
	return c
}