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

func newCommentsReaction(db *gorm.DB, opts ...gen.DOOption) commentsReaction {
	_commentsReaction := commentsReaction{}

	_commentsReaction.commentsReactionDo.UseDB(db, opts...)
	_commentsReaction.commentsReactionDo.UseModel(&models.CommentsReaction{})

	tableName := _commentsReaction.commentsReactionDo.TableName()
	_commentsReaction.ALL = field.NewAsterisk(tableName)
	_commentsReaction.CommentID = field.NewUint64(tableName, "comment_id")
	_commentsReaction.Reaction = field.NewString(tableName, "reaction")
	_commentsReaction.UserID = field.NewUint64(tableName, "user_id")
	_commentsReaction.CreatedAt = field.NewTime(tableName, "created_at")
	_commentsReaction.DeletedAt = field.NewField(tableName, "deleted_at")
	_commentsReaction.UpdatedAt = field.NewTime(tableName, "updated_at")

	_commentsReaction.fillFieldMap()

	return _commentsReaction
}

type commentsReaction struct {
	commentsReactionDo commentsReactionDo

	ALL       field.Asterisk
	CommentID field.Uint64
	Reaction  field.String
	UserID    field.Uint64
	CreatedAt field.Time
	DeletedAt field.Field
	UpdatedAt field.Time

	fieldMap map[string]field.Expr
}

func (c commentsReaction) Table(newTableName string) *commentsReaction {
	c.commentsReactionDo.UseTable(newTableName)
	return c.updateTableName(newTableName)
}

func (c commentsReaction) As(alias string) *commentsReaction {
	c.commentsReactionDo.DO = *(c.commentsReactionDo.As(alias).(*gen.DO))
	return c.updateTableName(alias)
}

func (c *commentsReaction) updateTableName(table string) *commentsReaction {
	c.ALL = field.NewAsterisk(table)
	c.CommentID = field.NewUint64(table, "comment_id")
	c.Reaction = field.NewString(table, "reaction")
	c.UserID = field.NewUint64(table, "user_id")
	c.CreatedAt = field.NewTime(table, "created_at")
	c.DeletedAt = field.NewField(table, "deleted_at")
	c.UpdatedAt = field.NewTime(table, "updated_at")

	c.fillFieldMap()

	return c
}

func (c *commentsReaction) WithContext(ctx context.Context) ICommentsReactionDo {
	return c.commentsReactionDo.WithContext(ctx)
}

func (c commentsReaction) TableName() string { return c.commentsReactionDo.TableName() }

func (c commentsReaction) Alias() string { return c.commentsReactionDo.Alias() }

func (c *commentsReaction) GetFieldByName(fieldName string) (field.OrderExpr, bool) {
	_f, ok := c.fieldMap[fieldName]
	if !ok || _f == nil {
		return nil, false
	}
	_oe, ok := _f.(field.OrderExpr)
	return _oe, ok
}

func (c *commentsReaction) fillFieldMap() {
	c.fieldMap = make(map[string]field.Expr, 6)
	c.fieldMap["comment_id"] = c.CommentID
	c.fieldMap["reaction"] = c.Reaction
	c.fieldMap["user_id"] = c.UserID
	c.fieldMap["created_at"] = c.CreatedAt
	c.fieldMap["deleted_at"] = c.DeletedAt
	c.fieldMap["updated_at"] = c.UpdatedAt
}

func (c commentsReaction) clone(db *gorm.DB) commentsReaction {
	c.commentsReactionDo.ReplaceConnPool(db.Statement.ConnPool)
	return c
}

func (c commentsReaction) replaceDB(db *gorm.DB) commentsReaction {
	c.commentsReactionDo.ReplaceDB(db)
	return c
}

type commentsReactionDo struct{ gen.DO }

type ICommentsReactionDo interface {
	gen.SubQuery
	Debug() ICommentsReactionDo
	WithContext(ctx context.Context) ICommentsReactionDo
	WithResult(fc func(tx gen.Dao)) gen.ResultInfo
	ReplaceDB(db *gorm.DB)
	ReadDB() ICommentsReactionDo
	WriteDB() ICommentsReactionDo
	As(alias string) gen.Dao
	Session(config *gorm.Session) ICommentsReactionDo
	Columns(cols ...field.Expr) gen.Columns
	Clauses(conds ...clause.Expression) ICommentsReactionDo
	Not(conds ...gen.Condition) ICommentsReactionDo
	Or(conds ...gen.Condition) ICommentsReactionDo
	Select(conds ...field.Expr) ICommentsReactionDo
	Where(conds ...gen.Condition) ICommentsReactionDo
	Order(conds ...field.Expr) ICommentsReactionDo
	Distinct(cols ...field.Expr) ICommentsReactionDo
	Omit(cols ...field.Expr) ICommentsReactionDo
	Join(table schema.Tabler, on ...field.Expr) ICommentsReactionDo
	LeftJoin(table schema.Tabler, on ...field.Expr) ICommentsReactionDo
	RightJoin(table schema.Tabler, on ...field.Expr) ICommentsReactionDo
	Group(cols ...field.Expr) ICommentsReactionDo
	Having(conds ...gen.Condition) ICommentsReactionDo
	Limit(limit int) ICommentsReactionDo
	Offset(offset int) ICommentsReactionDo
	Count() (count int64, err error)
	Scopes(funcs ...func(gen.Dao) gen.Dao) ICommentsReactionDo
	Unscoped() ICommentsReactionDo
	Create(values ...*models.CommentsReaction) error
	CreateInBatches(values []*models.CommentsReaction, batchSize int) error
	Save(values ...*models.CommentsReaction) error
	First() (*models.CommentsReaction, error)
	Take() (*models.CommentsReaction, error)
	Last() (*models.CommentsReaction, error)
	Find() ([]*models.CommentsReaction, error)
	FindInBatch(batchSize int, fc func(tx gen.Dao, batch int) error) (results []*models.CommentsReaction, err error)
	FindInBatches(result *[]*models.CommentsReaction, batchSize int, fc func(tx gen.Dao, batch int) error) error
	Pluck(column field.Expr, dest interface{}) error
	Delete(...*models.CommentsReaction) (info gen.ResultInfo, err error)
	Update(column field.Expr, value interface{}) (info gen.ResultInfo, err error)
	UpdateSimple(columns ...field.AssignExpr) (info gen.ResultInfo, err error)
	Updates(value interface{}) (info gen.ResultInfo, err error)
	UpdateColumn(column field.Expr, value interface{}) (info gen.ResultInfo, err error)
	UpdateColumnSimple(columns ...field.AssignExpr) (info gen.ResultInfo, err error)
	UpdateColumns(value interface{}) (info gen.ResultInfo, err error)
	UpdateFrom(q gen.SubQuery) gen.Dao
	Attrs(attrs ...field.AssignExpr) ICommentsReactionDo
	Assign(attrs ...field.AssignExpr) ICommentsReactionDo
	Joins(fields ...field.RelationField) ICommentsReactionDo
	Preload(fields ...field.RelationField) ICommentsReactionDo
	FirstOrInit() (*models.CommentsReaction, error)
	FirstOrCreate() (*models.CommentsReaction, error)
	FindByPage(offset int, limit int) (result []*models.CommentsReaction, count int64, err error)
	ScanByPage(result interface{}, offset int, limit int) (count int64, err error)
	Scan(result interface{}) (err error)
	Returning(value interface{}, columns ...string) ICommentsReactionDo
	UnderlyingDB() *gorm.DB
	schema.Tabler
}

func (c commentsReactionDo) Debug() ICommentsReactionDo {
	return c.withDO(c.DO.Debug())
}

func (c commentsReactionDo) WithContext(ctx context.Context) ICommentsReactionDo {
	return c.withDO(c.DO.WithContext(ctx))
}

func (c commentsReactionDo) ReadDB() ICommentsReactionDo {
	return c.Clauses(dbresolver.Read)
}

func (c commentsReactionDo) WriteDB() ICommentsReactionDo {
	return c.Clauses(dbresolver.Write)
}

func (c commentsReactionDo) Session(config *gorm.Session) ICommentsReactionDo {
	return c.withDO(c.DO.Session(config))
}

func (c commentsReactionDo) Clauses(conds ...clause.Expression) ICommentsReactionDo {
	return c.withDO(c.DO.Clauses(conds...))
}

func (c commentsReactionDo) Returning(value interface{}, columns ...string) ICommentsReactionDo {
	return c.withDO(c.DO.Returning(value, columns...))
}

func (c commentsReactionDo) Not(conds ...gen.Condition) ICommentsReactionDo {
	return c.withDO(c.DO.Not(conds...))
}

func (c commentsReactionDo) Or(conds ...gen.Condition) ICommentsReactionDo {
	return c.withDO(c.DO.Or(conds...))
}

func (c commentsReactionDo) Select(conds ...field.Expr) ICommentsReactionDo {
	return c.withDO(c.DO.Select(conds...))
}

func (c commentsReactionDo) Where(conds ...gen.Condition) ICommentsReactionDo {
	return c.withDO(c.DO.Where(conds...))
}

func (c commentsReactionDo) Exists(subquery interface{ UnderlyingDB() *gorm.DB }) ICommentsReactionDo {
	return c.Where(field.CompareSubQuery(field.ExistsOp, nil, subquery.UnderlyingDB()))
}

func (c commentsReactionDo) Order(conds ...field.Expr) ICommentsReactionDo {
	return c.withDO(c.DO.Order(conds...))
}

func (c commentsReactionDo) Distinct(cols ...field.Expr) ICommentsReactionDo {
	return c.withDO(c.DO.Distinct(cols...))
}

func (c commentsReactionDo) Omit(cols ...field.Expr) ICommentsReactionDo {
	return c.withDO(c.DO.Omit(cols...))
}

func (c commentsReactionDo) Join(table schema.Tabler, on ...field.Expr) ICommentsReactionDo {
	return c.withDO(c.DO.Join(table, on...))
}

func (c commentsReactionDo) LeftJoin(table schema.Tabler, on ...field.Expr) ICommentsReactionDo {
	return c.withDO(c.DO.LeftJoin(table, on...))
}

func (c commentsReactionDo) RightJoin(table schema.Tabler, on ...field.Expr) ICommentsReactionDo {
	return c.withDO(c.DO.RightJoin(table, on...))
}

func (c commentsReactionDo) Group(cols ...field.Expr) ICommentsReactionDo {
	return c.withDO(c.DO.Group(cols...))
}

func (c commentsReactionDo) Having(conds ...gen.Condition) ICommentsReactionDo {
	return c.withDO(c.DO.Having(conds...))
}

func (c commentsReactionDo) Limit(limit int) ICommentsReactionDo {
	return c.withDO(c.DO.Limit(limit))
}

func (c commentsReactionDo) Offset(offset int) ICommentsReactionDo {
	return c.withDO(c.DO.Offset(offset))
}

func (c commentsReactionDo) Scopes(funcs ...func(gen.Dao) gen.Dao) ICommentsReactionDo {
	return c.withDO(c.DO.Scopes(funcs...))
}

func (c commentsReactionDo) Unscoped() ICommentsReactionDo {
	return c.withDO(c.DO.Unscoped())
}

func (c commentsReactionDo) Create(values ...*models.CommentsReaction) error {
	if len(values) == 0 {
		return nil
	}
	return c.DO.Create(values)
}

func (c commentsReactionDo) CreateInBatches(values []*models.CommentsReaction, batchSize int) error {
	return c.DO.CreateInBatches(values, batchSize)
}

// Save : !!! underlying implementation is different with GORM
// The method is equivalent to executing the statement: db.Clauses(clause.OnConflict{UpdateAll: true}).Create(values)
func (c commentsReactionDo) Save(values ...*models.CommentsReaction) error {
	if len(values) == 0 {
		return nil
	}
	return c.DO.Save(values)
}

func (c commentsReactionDo) First() (*models.CommentsReaction, error) {
	if result, err := c.DO.First(); err != nil {
		return nil, err
	} else {
		return result.(*models.CommentsReaction), nil
	}
}

func (c commentsReactionDo) Take() (*models.CommentsReaction, error) {
	if result, err := c.DO.Take(); err != nil {
		return nil, err
	} else {
		return result.(*models.CommentsReaction), nil
	}
}

func (c commentsReactionDo) Last() (*models.CommentsReaction, error) {
	if result, err := c.DO.Last(); err != nil {
		return nil, err
	} else {
		return result.(*models.CommentsReaction), nil
	}
}

func (c commentsReactionDo) Find() ([]*models.CommentsReaction, error) {
	result, err := c.DO.Find()
	return result.([]*models.CommentsReaction), err
}

func (c commentsReactionDo) FindInBatch(batchSize int, fc func(tx gen.Dao, batch int) error) (results []*models.CommentsReaction, err error) {
	buf := make([]*models.CommentsReaction, 0, batchSize)
	err = c.DO.FindInBatches(&buf, batchSize, func(tx gen.Dao, batch int) error {
		defer func() { results = append(results, buf...) }()
		return fc(tx, batch)
	})
	return results, err
}

func (c commentsReactionDo) FindInBatches(result *[]*models.CommentsReaction, batchSize int, fc func(tx gen.Dao, batch int) error) error {
	return c.DO.FindInBatches(result, batchSize, fc)
}

func (c commentsReactionDo) Attrs(attrs ...field.AssignExpr) ICommentsReactionDo {
	return c.withDO(c.DO.Attrs(attrs...))
}

func (c commentsReactionDo) Assign(attrs ...field.AssignExpr) ICommentsReactionDo {
	return c.withDO(c.DO.Assign(attrs...))
}

func (c commentsReactionDo) Joins(fields ...field.RelationField) ICommentsReactionDo {
	for _, _f := range fields {
		c = *c.withDO(c.DO.Joins(_f))
	}
	return &c
}

func (c commentsReactionDo) Preload(fields ...field.RelationField) ICommentsReactionDo {
	for _, _f := range fields {
		c = *c.withDO(c.DO.Preload(_f))
	}
	return &c
}

func (c commentsReactionDo) FirstOrInit() (*models.CommentsReaction, error) {
	if result, err := c.DO.FirstOrInit(); err != nil {
		return nil, err
	} else {
		return result.(*models.CommentsReaction), nil
	}
}

func (c commentsReactionDo) FirstOrCreate() (*models.CommentsReaction, error) {
	if result, err := c.DO.FirstOrCreate(); err != nil {
		return nil, err
	} else {
		return result.(*models.CommentsReaction), nil
	}
}

func (c commentsReactionDo) FindByPage(offset int, limit int) (result []*models.CommentsReaction, count int64, err error) {
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

func (c commentsReactionDo) ScanByPage(result interface{}, offset int, limit int) (count int64, err error) {
	count, err = c.Count()
	if err != nil {
		return
	}

	err = c.Offset(offset).Limit(limit).Scan(result)
	return
}

func (c commentsReactionDo) Scan(result interface{}) (err error) {
	return c.DO.Scan(result)
}

func (c commentsReactionDo) Delete(models ...*models.CommentsReaction) (result gen.ResultInfo, err error) {
	return c.DO.Delete(models)
}

func (c *commentsReactionDo) withDO(do gen.Dao) *commentsReactionDo {
	c.DO = *do.(*gen.DO)
	return c
}
