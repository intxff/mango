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

	"github.com/intxff/mango/posts/internal/models"
)

func newPostsReaction(db *gorm.DB, opts ...gen.DOOption) postsReaction {
	_postsReaction := postsReaction{}

	_postsReaction.postsReactionDo.UseDB(db, opts...)
	_postsReaction.postsReactionDo.UseModel(&models.PostsReaction{})

	tableName := _postsReaction.postsReactionDo.TableName()
	_postsReaction.ALL = field.NewAsterisk(tableName)
	_postsReaction.UserID = field.NewUint64(tableName, "user_id")
	_postsReaction.Reaction = field.NewString(tableName, "reaction")
	_postsReaction.PostID = field.NewUint64(tableName, "post_id")
	_postsReaction.CreatedAt = field.NewTime(tableName, "created_at")
	_postsReaction.DeletedAt = field.NewField(tableName, "deleted_at")
	_postsReaction.UpdatedAt = field.NewTime(tableName, "updated_at")

	_postsReaction.fillFieldMap()

	return _postsReaction
}

type postsReaction struct {
	postsReactionDo postsReactionDo

	ALL       field.Asterisk
	UserID    field.Uint64
	Reaction  field.String
	PostID    field.Uint64
	CreatedAt field.Time
	DeletedAt field.Field
	UpdatedAt field.Time

	fieldMap map[string]field.Expr
}

func (p postsReaction) Table(newTableName string) *postsReaction {
	p.postsReactionDo.UseTable(newTableName)
	return p.updateTableName(newTableName)
}

func (p postsReaction) As(alias string) *postsReaction {
	p.postsReactionDo.DO = *(p.postsReactionDo.As(alias).(*gen.DO))
	return p.updateTableName(alias)
}

func (p *postsReaction) updateTableName(table string) *postsReaction {
	p.ALL = field.NewAsterisk(table)
	p.UserID = field.NewUint64(table, "user_id")
	p.Reaction = field.NewString(table, "reaction")
	p.PostID = field.NewUint64(table, "post_id")
	p.CreatedAt = field.NewTime(table, "created_at")
	p.DeletedAt = field.NewField(table, "deleted_at")
	p.UpdatedAt = field.NewTime(table, "updated_at")

	p.fillFieldMap()

	return p
}

func (p *postsReaction) WithContext(ctx context.Context) IPostsReactionDo {
	return p.postsReactionDo.WithContext(ctx)
}

func (p postsReaction) TableName() string { return p.postsReactionDo.TableName() }

func (p postsReaction) Alias() string { return p.postsReactionDo.Alias() }

func (p *postsReaction) GetFieldByName(fieldName string) (field.OrderExpr, bool) {
	_f, ok := p.fieldMap[fieldName]
	if !ok || _f == nil {
		return nil, false
	}
	_oe, ok := _f.(field.OrderExpr)
	return _oe, ok
}

func (p *postsReaction) fillFieldMap() {
	p.fieldMap = make(map[string]field.Expr, 6)
	p.fieldMap["user_id"] = p.UserID
	p.fieldMap["reaction"] = p.Reaction
	p.fieldMap["post_id"] = p.PostID
	p.fieldMap["created_at"] = p.CreatedAt
	p.fieldMap["deleted_at"] = p.DeletedAt
	p.fieldMap["updated_at"] = p.UpdatedAt
}

func (p postsReaction) clone(db *gorm.DB) postsReaction {
	p.postsReactionDo.ReplaceConnPool(db.Statement.ConnPool)
	return p
}

func (p postsReaction) replaceDB(db *gorm.DB) postsReaction {
	p.postsReactionDo.ReplaceDB(db)
	return p
}

type postsReactionDo struct{ gen.DO }

type IPostsReactionDo interface {
	gen.SubQuery
	Debug() IPostsReactionDo
	WithContext(ctx context.Context) IPostsReactionDo
	WithResult(fc func(tx gen.Dao)) gen.ResultInfo
	ReplaceDB(db *gorm.DB)
	ReadDB() IPostsReactionDo
	WriteDB() IPostsReactionDo
	As(alias string) gen.Dao
	Session(config *gorm.Session) IPostsReactionDo
	Columns(cols ...field.Expr) gen.Columns
	Clauses(conds ...clause.Expression) IPostsReactionDo
	Not(conds ...gen.Condition) IPostsReactionDo
	Or(conds ...gen.Condition) IPostsReactionDo
	Select(conds ...field.Expr) IPostsReactionDo
	Where(conds ...gen.Condition) IPostsReactionDo
	Order(conds ...field.Expr) IPostsReactionDo
	Distinct(cols ...field.Expr) IPostsReactionDo
	Omit(cols ...field.Expr) IPostsReactionDo
	Join(table schema.Tabler, on ...field.Expr) IPostsReactionDo
	LeftJoin(table schema.Tabler, on ...field.Expr) IPostsReactionDo
	RightJoin(table schema.Tabler, on ...field.Expr) IPostsReactionDo
	Group(cols ...field.Expr) IPostsReactionDo
	Having(conds ...gen.Condition) IPostsReactionDo
	Limit(limit int) IPostsReactionDo
	Offset(offset int) IPostsReactionDo
	Count() (count int64, err error)
	Scopes(funcs ...func(gen.Dao) gen.Dao) IPostsReactionDo
	Unscoped() IPostsReactionDo
	Create(values ...*models.PostsReaction) error
	CreateInBatches(values []*models.PostsReaction, batchSize int) error
	Save(values ...*models.PostsReaction) error
	First() (*models.PostsReaction, error)
	Take() (*models.PostsReaction, error)
	Last() (*models.PostsReaction, error)
	Find() ([]*models.PostsReaction, error)
	FindInBatch(batchSize int, fc func(tx gen.Dao, batch int) error) (results []*models.PostsReaction, err error)
	FindInBatches(result *[]*models.PostsReaction, batchSize int, fc func(tx gen.Dao, batch int) error) error
	Pluck(column field.Expr, dest interface{}) error
	Delete(...*models.PostsReaction) (info gen.ResultInfo, err error)
	Update(column field.Expr, value interface{}) (info gen.ResultInfo, err error)
	UpdateSimple(columns ...field.AssignExpr) (info gen.ResultInfo, err error)
	Updates(value interface{}) (info gen.ResultInfo, err error)
	UpdateColumn(column field.Expr, value interface{}) (info gen.ResultInfo, err error)
	UpdateColumnSimple(columns ...field.AssignExpr) (info gen.ResultInfo, err error)
	UpdateColumns(value interface{}) (info gen.ResultInfo, err error)
	UpdateFrom(q gen.SubQuery) gen.Dao
	Attrs(attrs ...field.AssignExpr) IPostsReactionDo
	Assign(attrs ...field.AssignExpr) IPostsReactionDo
	Joins(fields ...field.RelationField) IPostsReactionDo
	Preload(fields ...field.RelationField) IPostsReactionDo
	FirstOrInit() (*models.PostsReaction, error)
	FirstOrCreate() (*models.PostsReaction, error)
	FindByPage(offset int, limit int) (result []*models.PostsReaction, count int64, err error)
	ScanByPage(result interface{}, offset int, limit int) (count int64, err error)
	Scan(result interface{}) (err error)
	Returning(value interface{}, columns ...string) IPostsReactionDo
	UnderlyingDB() *gorm.DB
	schema.Tabler
}

func (p postsReactionDo) Debug() IPostsReactionDo {
	return p.withDO(p.DO.Debug())
}

func (p postsReactionDo) WithContext(ctx context.Context) IPostsReactionDo {
	return p.withDO(p.DO.WithContext(ctx))
}

func (p postsReactionDo) ReadDB() IPostsReactionDo {
	return p.Clauses(dbresolver.Read)
}

func (p postsReactionDo) WriteDB() IPostsReactionDo {
	return p.Clauses(dbresolver.Write)
}

func (p postsReactionDo) Session(config *gorm.Session) IPostsReactionDo {
	return p.withDO(p.DO.Session(config))
}

func (p postsReactionDo) Clauses(conds ...clause.Expression) IPostsReactionDo {
	return p.withDO(p.DO.Clauses(conds...))
}

func (p postsReactionDo) Returning(value interface{}, columns ...string) IPostsReactionDo {
	return p.withDO(p.DO.Returning(value, columns...))
}

func (p postsReactionDo) Not(conds ...gen.Condition) IPostsReactionDo {
	return p.withDO(p.DO.Not(conds...))
}

func (p postsReactionDo) Or(conds ...gen.Condition) IPostsReactionDo {
	return p.withDO(p.DO.Or(conds...))
}

func (p postsReactionDo) Select(conds ...field.Expr) IPostsReactionDo {
	return p.withDO(p.DO.Select(conds...))
}

func (p postsReactionDo) Where(conds ...gen.Condition) IPostsReactionDo {
	return p.withDO(p.DO.Where(conds...))
}

func (p postsReactionDo) Exists(subquery interface{ UnderlyingDB() *gorm.DB }) IPostsReactionDo {
	return p.Where(field.CompareSubQuery(field.ExistsOp, nil, subquery.UnderlyingDB()))
}

func (p postsReactionDo) Order(conds ...field.Expr) IPostsReactionDo {
	return p.withDO(p.DO.Order(conds...))
}

func (p postsReactionDo) Distinct(cols ...field.Expr) IPostsReactionDo {
	return p.withDO(p.DO.Distinct(cols...))
}

func (p postsReactionDo) Omit(cols ...field.Expr) IPostsReactionDo {
	return p.withDO(p.DO.Omit(cols...))
}

func (p postsReactionDo) Join(table schema.Tabler, on ...field.Expr) IPostsReactionDo {
	return p.withDO(p.DO.Join(table, on...))
}

func (p postsReactionDo) LeftJoin(table schema.Tabler, on ...field.Expr) IPostsReactionDo {
	return p.withDO(p.DO.LeftJoin(table, on...))
}

func (p postsReactionDo) RightJoin(table schema.Tabler, on ...field.Expr) IPostsReactionDo {
	return p.withDO(p.DO.RightJoin(table, on...))
}

func (p postsReactionDo) Group(cols ...field.Expr) IPostsReactionDo {
	return p.withDO(p.DO.Group(cols...))
}

func (p postsReactionDo) Having(conds ...gen.Condition) IPostsReactionDo {
	return p.withDO(p.DO.Having(conds...))
}

func (p postsReactionDo) Limit(limit int) IPostsReactionDo {
	return p.withDO(p.DO.Limit(limit))
}

func (p postsReactionDo) Offset(offset int) IPostsReactionDo {
	return p.withDO(p.DO.Offset(offset))
}

func (p postsReactionDo) Scopes(funcs ...func(gen.Dao) gen.Dao) IPostsReactionDo {
	return p.withDO(p.DO.Scopes(funcs...))
}

func (p postsReactionDo) Unscoped() IPostsReactionDo {
	return p.withDO(p.DO.Unscoped())
}

func (p postsReactionDo) Create(values ...*models.PostsReaction) error {
	if len(values) == 0 {
		return nil
	}
	return p.DO.Create(values)
}

func (p postsReactionDo) CreateInBatches(values []*models.PostsReaction, batchSize int) error {
	return p.DO.CreateInBatches(values, batchSize)
}

// Save : !!! underlying implementation is different with GORM
// The method is equivalent to executing the statement: db.Clauses(clause.OnConflict{UpdateAll: true}).Create(values)
func (p postsReactionDo) Save(values ...*models.PostsReaction) error {
	if len(values) == 0 {
		return nil
	}
	return p.DO.Save(values)
}

func (p postsReactionDo) First() (*models.PostsReaction, error) {
	if result, err := p.DO.First(); err != nil {
		return nil, err
	} else {
		return result.(*models.PostsReaction), nil
	}
}

func (p postsReactionDo) Take() (*models.PostsReaction, error) {
	if result, err := p.DO.Take(); err != nil {
		return nil, err
	} else {
		return result.(*models.PostsReaction), nil
	}
}

func (p postsReactionDo) Last() (*models.PostsReaction, error) {
	if result, err := p.DO.Last(); err != nil {
		return nil, err
	} else {
		return result.(*models.PostsReaction), nil
	}
}

func (p postsReactionDo) Find() ([]*models.PostsReaction, error) {
	result, err := p.DO.Find()
	return result.([]*models.PostsReaction), err
}

func (p postsReactionDo) FindInBatch(batchSize int, fc func(tx gen.Dao, batch int) error) (results []*models.PostsReaction, err error) {
	buf := make([]*models.PostsReaction, 0, batchSize)
	err = p.DO.FindInBatches(&buf, batchSize, func(tx gen.Dao, batch int) error {
		defer func() { results = append(results, buf...) }()
		return fc(tx, batch)
	})
	return results, err
}

func (p postsReactionDo) FindInBatches(result *[]*models.PostsReaction, batchSize int, fc func(tx gen.Dao, batch int) error) error {
	return p.DO.FindInBatches(result, batchSize, fc)
}

func (p postsReactionDo) Attrs(attrs ...field.AssignExpr) IPostsReactionDo {
	return p.withDO(p.DO.Attrs(attrs...))
}

func (p postsReactionDo) Assign(attrs ...field.AssignExpr) IPostsReactionDo {
	return p.withDO(p.DO.Assign(attrs...))
}

func (p postsReactionDo) Joins(fields ...field.RelationField) IPostsReactionDo {
	for _, _f := range fields {
		p = *p.withDO(p.DO.Joins(_f))
	}
	return &p
}

func (p postsReactionDo) Preload(fields ...field.RelationField) IPostsReactionDo {
	for _, _f := range fields {
		p = *p.withDO(p.DO.Preload(_f))
	}
	return &p
}

func (p postsReactionDo) FirstOrInit() (*models.PostsReaction, error) {
	if result, err := p.DO.FirstOrInit(); err != nil {
		return nil, err
	} else {
		return result.(*models.PostsReaction), nil
	}
}

func (p postsReactionDo) FirstOrCreate() (*models.PostsReaction, error) {
	if result, err := p.DO.FirstOrCreate(); err != nil {
		return nil, err
	} else {
		return result.(*models.PostsReaction), nil
	}
}

func (p postsReactionDo) FindByPage(offset int, limit int) (result []*models.PostsReaction, count int64, err error) {
	result, err = p.Offset(offset).Limit(limit).Find()
	if err != nil {
		return
	}

	if size := len(result); 0 < limit && 0 < size && size < limit {
		count = int64(size + offset)
		return
	}

	count, err = p.Offset(-1).Limit(-1).Count()
	return
}

func (p postsReactionDo) ScanByPage(result interface{}, offset int, limit int) (count int64, err error) {
	count, err = p.Count()
	if err != nil {
		return
	}

	err = p.Offset(offset).Limit(limit).Scan(result)
	return
}

func (p postsReactionDo) Scan(result interface{}) (err error) {
	return p.DO.Scan(result)
}

func (p postsReactionDo) Delete(models ...*models.PostsReaction) (result gen.ResultInfo, err error) {
	return p.DO.Delete(models)
}

func (p *postsReactionDo) withDO(do gen.Dao) *postsReactionDo {
	p.DO = *do.(*gen.DO)
	return p
}
