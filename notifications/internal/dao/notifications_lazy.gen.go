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

	"github.com/intxff/mango/notifications/internal/models"
)

func newNotificationsLazy(db *gorm.DB, opts ...gen.DOOption) notificationsLazy {
	_notificationsLazy := notificationsLazy{}

	_notificationsLazy.notificationsLazyDo.UseDB(db, opts...)
	_notificationsLazy.notificationsLazyDo.UseModel(&models.NotificationsLazy{})

	tableName := _notificationsLazy.notificationsLazyDo.TableName()
	_notificationsLazy.ALL = field.NewAsterisk(tableName)
	_notificationsLazy.NotificationType = field.NewString(tableName, "notification_type")
	_notificationsLazy.ObjectID = field.NewUint64(tableName, "object_id")
	_notificationsLazy.LastScanTime = field.NewTime(tableName, "last_scan_time")
	_notificationsLazy.CreatedAt = field.NewTime(tableName, "created_at")
	_notificationsLazy.DeletedAt = field.NewField(tableName, "deleted_at")
	_notificationsLazy.UpdatedAt = field.NewTime(tableName, "updated_at")

	_notificationsLazy.fillFieldMap()

	return _notificationsLazy
}

type notificationsLazy struct {
	notificationsLazyDo notificationsLazyDo

	ALL              field.Asterisk
	NotificationType field.String
	ObjectID         field.Uint64
	LastScanTime     field.Time
	CreatedAt        field.Time
	DeletedAt        field.Field
	UpdatedAt        field.Time

	fieldMap map[string]field.Expr
}

func (n notificationsLazy) Table(newTableName string) *notificationsLazy {
	n.notificationsLazyDo.UseTable(newTableName)
	return n.updateTableName(newTableName)
}

func (n notificationsLazy) As(alias string) *notificationsLazy {
	n.notificationsLazyDo.DO = *(n.notificationsLazyDo.As(alias).(*gen.DO))
	return n.updateTableName(alias)
}

func (n *notificationsLazy) updateTableName(table string) *notificationsLazy {
	n.ALL = field.NewAsterisk(table)
	n.NotificationType = field.NewString(table, "notification_type")
	n.ObjectID = field.NewUint64(table, "object_id")
	n.LastScanTime = field.NewTime(table, "last_scan_time")
	n.CreatedAt = field.NewTime(table, "created_at")
	n.DeletedAt = field.NewField(table, "deleted_at")
	n.UpdatedAt = field.NewTime(table, "updated_at")

	n.fillFieldMap()

	return n
}

func (n *notificationsLazy) WithContext(ctx context.Context) INotificationsLazyDo {
	return n.notificationsLazyDo.WithContext(ctx)
}

func (n notificationsLazy) TableName() string { return n.notificationsLazyDo.TableName() }

func (n notificationsLazy) Alias() string { return n.notificationsLazyDo.Alias() }

func (n *notificationsLazy) GetFieldByName(fieldName string) (field.OrderExpr, bool) {
	_f, ok := n.fieldMap[fieldName]
	if !ok || _f == nil {
		return nil, false
	}
	_oe, ok := _f.(field.OrderExpr)
	return _oe, ok
}

func (n *notificationsLazy) fillFieldMap() {
	n.fieldMap = make(map[string]field.Expr, 6)
	n.fieldMap["notification_type"] = n.NotificationType
	n.fieldMap["object_id"] = n.ObjectID
	n.fieldMap["last_scan_time"] = n.LastScanTime
	n.fieldMap["created_at"] = n.CreatedAt
	n.fieldMap["deleted_at"] = n.DeletedAt
	n.fieldMap["updated_at"] = n.UpdatedAt
}

func (n notificationsLazy) clone(db *gorm.DB) notificationsLazy {
	n.notificationsLazyDo.ReplaceConnPool(db.Statement.ConnPool)
	return n
}

func (n notificationsLazy) replaceDB(db *gorm.DB) notificationsLazy {
	n.notificationsLazyDo.ReplaceDB(db)
	return n
}

type notificationsLazyDo struct{ gen.DO }

type INotificationsLazyDo interface {
	gen.SubQuery
	Debug() INotificationsLazyDo
	WithContext(ctx context.Context) INotificationsLazyDo
	WithResult(fc func(tx gen.Dao)) gen.ResultInfo
	ReplaceDB(db *gorm.DB)
	ReadDB() INotificationsLazyDo
	WriteDB() INotificationsLazyDo
	As(alias string) gen.Dao
	Session(config *gorm.Session) INotificationsLazyDo
	Columns(cols ...field.Expr) gen.Columns
	Clauses(conds ...clause.Expression) INotificationsLazyDo
	Not(conds ...gen.Condition) INotificationsLazyDo
	Or(conds ...gen.Condition) INotificationsLazyDo
	Select(conds ...field.Expr) INotificationsLazyDo
	Where(conds ...gen.Condition) INotificationsLazyDo
	Order(conds ...field.Expr) INotificationsLazyDo
	Distinct(cols ...field.Expr) INotificationsLazyDo
	Omit(cols ...field.Expr) INotificationsLazyDo
	Join(table schema.Tabler, on ...field.Expr) INotificationsLazyDo
	LeftJoin(table schema.Tabler, on ...field.Expr) INotificationsLazyDo
	RightJoin(table schema.Tabler, on ...field.Expr) INotificationsLazyDo
	Group(cols ...field.Expr) INotificationsLazyDo
	Having(conds ...gen.Condition) INotificationsLazyDo
	Limit(limit int) INotificationsLazyDo
	Offset(offset int) INotificationsLazyDo
	Count() (count int64, err error)
	Scopes(funcs ...func(gen.Dao) gen.Dao) INotificationsLazyDo
	Unscoped() INotificationsLazyDo
	Create(values ...*models.NotificationsLazy) error
	CreateInBatches(values []*models.NotificationsLazy, batchSize int) error
	Save(values ...*models.NotificationsLazy) error
	First() (*models.NotificationsLazy, error)
	Take() (*models.NotificationsLazy, error)
	Last() (*models.NotificationsLazy, error)
	Find() ([]*models.NotificationsLazy, error)
	FindInBatch(batchSize int, fc func(tx gen.Dao, batch int) error) (results []*models.NotificationsLazy, err error)
	FindInBatches(result *[]*models.NotificationsLazy, batchSize int, fc func(tx gen.Dao, batch int) error) error
	Pluck(column field.Expr, dest interface{}) error
	Delete(...*models.NotificationsLazy) (info gen.ResultInfo, err error)
	Update(column field.Expr, value interface{}) (info gen.ResultInfo, err error)
	UpdateSimple(columns ...field.AssignExpr) (info gen.ResultInfo, err error)
	Updates(value interface{}) (info gen.ResultInfo, err error)
	UpdateColumn(column field.Expr, value interface{}) (info gen.ResultInfo, err error)
	UpdateColumnSimple(columns ...field.AssignExpr) (info gen.ResultInfo, err error)
	UpdateColumns(value interface{}) (info gen.ResultInfo, err error)
	UpdateFrom(q gen.SubQuery) gen.Dao
	Attrs(attrs ...field.AssignExpr) INotificationsLazyDo
	Assign(attrs ...field.AssignExpr) INotificationsLazyDo
	Joins(fields ...field.RelationField) INotificationsLazyDo
	Preload(fields ...field.RelationField) INotificationsLazyDo
	FirstOrInit() (*models.NotificationsLazy, error)
	FirstOrCreate() (*models.NotificationsLazy, error)
	FindByPage(offset int, limit int) (result []*models.NotificationsLazy, count int64, err error)
	ScanByPage(result interface{}, offset int, limit int) (count int64, err error)
	Scan(result interface{}) (err error)
	Returning(value interface{}, columns ...string) INotificationsLazyDo
	UnderlyingDB() *gorm.DB
	schema.Tabler
}

func (n notificationsLazyDo) Debug() INotificationsLazyDo {
	return n.withDO(n.DO.Debug())
}

func (n notificationsLazyDo) WithContext(ctx context.Context) INotificationsLazyDo {
	return n.withDO(n.DO.WithContext(ctx))
}

func (n notificationsLazyDo) ReadDB() INotificationsLazyDo {
	return n.Clauses(dbresolver.Read)
}

func (n notificationsLazyDo) WriteDB() INotificationsLazyDo {
	return n.Clauses(dbresolver.Write)
}

func (n notificationsLazyDo) Session(config *gorm.Session) INotificationsLazyDo {
	return n.withDO(n.DO.Session(config))
}

func (n notificationsLazyDo) Clauses(conds ...clause.Expression) INotificationsLazyDo {
	return n.withDO(n.DO.Clauses(conds...))
}

func (n notificationsLazyDo) Returning(value interface{}, columns ...string) INotificationsLazyDo {
	return n.withDO(n.DO.Returning(value, columns...))
}

func (n notificationsLazyDo) Not(conds ...gen.Condition) INotificationsLazyDo {
	return n.withDO(n.DO.Not(conds...))
}

func (n notificationsLazyDo) Or(conds ...gen.Condition) INotificationsLazyDo {
	return n.withDO(n.DO.Or(conds...))
}

func (n notificationsLazyDo) Select(conds ...field.Expr) INotificationsLazyDo {
	return n.withDO(n.DO.Select(conds...))
}

func (n notificationsLazyDo) Where(conds ...gen.Condition) INotificationsLazyDo {
	return n.withDO(n.DO.Where(conds...))
}

func (n notificationsLazyDo) Exists(subquery interface{ UnderlyingDB() *gorm.DB }) INotificationsLazyDo {
	return n.Where(field.CompareSubQuery(field.ExistsOp, nil, subquery.UnderlyingDB()))
}

func (n notificationsLazyDo) Order(conds ...field.Expr) INotificationsLazyDo {
	return n.withDO(n.DO.Order(conds...))
}

func (n notificationsLazyDo) Distinct(cols ...field.Expr) INotificationsLazyDo {
	return n.withDO(n.DO.Distinct(cols...))
}

func (n notificationsLazyDo) Omit(cols ...field.Expr) INotificationsLazyDo {
	return n.withDO(n.DO.Omit(cols...))
}

func (n notificationsLazyDo) Join(table schema.Tabler, on ...field.Expr) INotificationsLazyDo {
	return n.withDO(n.DO.Join(table, on...))
}

func (n notificationsLazyDo) LeftJoin(table schema.Tabler, on ...field.Expr) INotificationsLazyDo {
	return n.withDO(n.DO.LeftJoin(table, on...))
}

func (n notificationsLazyDo) RightJoin(table schema.Tabler, on ...field.Expr) INotificationsLazyDo {
	return n.withDO(n.DO.RightJoin(table, on...))
}

func (n notificationsLazyDo) Group(cols ...field.Expr) INotificationsLazyDo {
	return n.withDO(n.DO.Group(cols...))
}

func (n notificationsLazyDo) Having(conds ...gen.Condition) INotificationsLazyDo {
	return n.withDO(n.DO.Having(conds...))
}

func (n notificationsLazyDo) Limit(limit int) INotificationsLazyDo {
	return n.withDO(n.DO.Limit(limit))
}

func (n notificationsLazyDo) Offset(offset int) INotificationsLazyDo {
	return n.withDO(n.DO.Offset(offset))
}

func (n notificationsLazyDo) Scopes(funcs ...func(gen.Dao) gen.Dao) INotificationsLazyDo {
	return n.withDO(n.DO.Scopes(funcs...))
}

func (n notificationsLazyDo) Unscoped() INotificationsLazyDo {
	return n.withDO(n.DO.Unscoped())
}

func (n notificationsLazyDo) Create(values ...*models.NotificationsLazy) error {
	if len(values) == 0 {
		return nil
	}
	return n.DO.Create(values)
}

func (n notificationsLazyDo) CreateInBatches(values []*models.NotificationsLazy, batchSize int) error {
	return n.DO.CreateInBatches(values, batchSize)
}

// Save : !!! underlying implementation is different with GORM
// The method is equivalent to executing the statement: db.Clauses(clause.OnConflict{UpdateAll: true}).Create(values)
func (n notificationsLazyDo) Save(values ...*models.NotificationsLazy) error {
	if len(values) == 0 {
		return nil
	}
	return n.DO.Save(values)
}

func (n notificationsLazyDo) First() (*models.NotificationsLazy, error) {
	if result, err := n.DO.First(); err != nil {
		return nil, err
	} else {
		return result.(*models.NotificationsLazy), nil
	}
}

func (n notificationsLazyDo) Take() (*models.NotificationsLazy, error) {
	if result, err := n.DO.Take(); err != nil {
		return nil, err
	} else {
		return result.(*models.NotificationsLazy), nil
	}
}

func (n notificationsLazyDo) Last() (*models.NotificationsLazy, error) {
	if result, err := n.DO.Last(); err != nil {
		return nil, err
	} else {
		return result.(*models.NotificationsLazy), nil
	}
}

func (n notificationsLazyDo) Find() ([]*models.NotificationsLazy, error) {
	result, err := n.DO.Find()
	return result.([]*models.NotificationsLazy), err
}

func (n notificationsLazyDo) FindInBatch(batchSize int, fc func(tx gen.Dao, batch int) error) (results []*models.NotificationsLazy, err error) {
	buf := make([]*models.NotificationsLazy, 0, batchSize)
	err = n.DO.FindInBatches(&buf, batchSize, func(tx gen.Dao, batch int) error {
		defer func() { results = append(results, buf...) }()
		return fc(tx, batch)
	})
	return results, err
}

func (n notificationsLazyDo) FindInBatches(result *[]*models.NotificationsLazy, batchSize int, fc func(tx gen.Dao, batch int) error) error {
	return n.DO.FindInBatches(result, batchSize, fc)
}

func (n notificationsLazyDo) Attrs(attrs ...field.AssignExpr) INotificationsLazyDo {
	return n.withDO(n.DO.Attrs(attrs...))
}

func (n notificationsLazyDo) Assign(attrs ...field.AssignExpr) INotificationsLazyDo {
	return n.withDO(n.DO.Assign(attrs...))
}

func (n notificationsLazyDo) Joins(fields ...field.RelationField) INotificationsLazyDo {
	for _, _f := range fields {
		n = *n.withDO(n.DO.Joins(_f))
	}
	return &n
}

func (n notificationsLazyDo) Preload(fields ...field.RelationField) INotificationsLazyDo {
	for _, _f := range fields {
		n = *n.withDO(n.DO.Preload(_f))
	}
	return &n
}

func (n notificationsLazyDo) FirstOrInit() (*models.NotificationsLazy, error) {
	if result, err := n.DO.FirstOrInit(); err != nil {
		return nil, err
	} else {
		return result.(*models.NotificationsLazy), nil
	}
}

func (n notificationsLazyDo) FirstOrCreate() (*models.NotificationsLazy, error) {
	if result, err := n.DO.FirstOrCreate(); err != nil {
		return nil, err
	} else {
		return result.(*models.NotificationsLazy), nil
	}
}

func (n notificationsLazyDo) FindByPage(offset int, limit int) (result []*models.NotificationsLazy, count int64, err error) {
	result, err = n.Offset(offset).Limit(limit).Find()
	if err != nil {
		return
	}

	if size := len(result); 0 < limit && 0 < size && size < limit {
		count = int64(size + offset)
		return
	}

	count, err = n.Offset(-1).Limit(-1).Count()
	return
}

func (n notificationsLazyDo) ScanByPage(result interface{}, offset int, limit int) (count int64, err error) {
	count, err = n.Count()
	if err != nil {
		return
	}

	err = n.Offset(offset).Limit(limit).Scan(result)
	return
}

func (n notificationsLazyDo) Scan(result interface{}) (err error) {
	return n.DO.Scan(result)
}

func (n notificationsLazyDo) Delete(models ...*models.NotificationsLazy) (result gen.ResultInfo, err error) {
	return n.DO.Delete(models)
}

func (n *notificationsLazyDo) withDO(do gen.Dao) *notificationsLazyDo {
	n.DO = *do.(*gen.DO)
	return n
}
