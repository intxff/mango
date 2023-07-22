// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.

package dao

import (
	"context"
	"fmt"
	"testing"

	"github.com/intxff/mango/comments/internal/models"
	"gorm.io/gen"
	"gorm.io/gen/field"
	"gorm.io/gorm/clause"
)

func init() {
	InitializeDB()
	err := db.AutoMigrate(&models.Comment{})
	if err != nil {
		fmt.Printf("Error: AutoMigrate(&models.Comment{}) fail: %s", err)
	}
}

func Test_commentQuery(t *testing.T) {
	comment := newComment(db)
	comment = *comment.As(comment.TableName())
	_do := comment.WithContext(context.Background()).Debug()

	primaryKey := field.NewString(comment.TableName(), clause.PrimaryKey)
	_, err := _do.Unscoped().Where(primaryKey.IsNotNull()).Delete()
	if err != nil {
		t.Error("clean table <comments> fail:", err)
		return
	}

	_, ok := comment.GetFieldByName("")
	if ok {
		t.Error("GetFieldByName(\"\") from comment success")
	}

	err = _do.Create(&models.Comment{})
	if err != nil {
		t.Error("create item in table <comments> fail:", err)
	}

	err = _do.Save(&models.Comment{})
	if err != nil {
		t.Error("create item in table <comments> fail:", err)
	}

	err = _do.CreateInBatches([]*models.Comment{{}, {}}, 10)
	if err != nil {
		t.Error("create item in table <comments> fail:", err)
	}

	_, err = _do.Select(comment.ALL).Take()
	if err != nil {
		t.Error("Take() on table <comments> fail:", err)
	}

	_, err = _do.First()
	if err != nil {
		t.Error("First() on table <comments> fail:", err)
	}

	_, err = _do.Last()
	if err != nil {
		t.Error("First() on table <comments> fail:", err)
	}

	_, err = _do.Where(primaryKey.IsNotNull()).FindInBatch(10, func(tx gen.Dao, batch int) error { return nil })
	if err != nil {
		t.Error("FindInBatch() on table <comments> fail:", err)
	}

	err = _do.Where(primaryKey.IsNotNull()).FindInBatches(&[]*models.Comment{}, 10, func(tx gen.Dao, batch int) error { return nil })
	if err != nil {
		t.Error("FindInBatches() on table <comments> fail:", err)
	}

	_, err = _do.Select(comment.ALL).Where(primaryKey.IsNotNull()).Order(primaryKey.Desc()).Find()
	if err != nil {
		t.Error("Find() on table <comments> fail:", err)
	}

	_, err = _do.Distinct(primaryKey).Take()
	if err != nil {
		t.Error("select Distinct() on table <comments> fail:", err)
	}

	_, err = _do.Select(comment.ALL).Omit(primaryKey).Take()
	if err != nil {
		t.Error("Omit() on table <comments> fail:", err)
	}

	_, err = _do.Group(primaryKey).Find()
	if err != nil {
		t.Error("Group() on table <comments> fail:", err)
	}

	_, err = _do.Scopes(func(dao gen.Dao) gen.Dao { return dao.Where(primaryKey.IsNotNull()) }).Find()
	if err != nil {
		t.Error("Scopes() on table <comments> fail:", err)
	}

	_, _, err = _do.FindByPage(0, 1)
	if err != nil {
		t.Error("FindByPage() on table <comments> fail:", err)
	}

	_, err = _do.ScanByPage(&models.Comment{}, 0, 1)
	if err != nil {
		t.Error("ScanByPage() on table <comments> fail:", err)
	}

	_, err = _do.Attrs(primaryKey).Assign(primaryKey).FirstOrInit()
	if err != nil {
		t.Error("FirstOrInit() on table <comments> fail:", err)
	}

	_, err = _do.Attrs(primaryKey).Assign(primaryKey).FirstOrCreate()
	if err != nil {
		t.Error("FirstOrCreate() on table <comments> fail:", err)
	}

	var _a _another
	var _aPK = field.NewString(_a.TableName(), clause.PrimaryKey)

	err = _do.Join(&_a, primaryKey.EqCol(_aPK)).Scan(map[string]interface{}{})
	if err != nil {
		t.Error("Join() on table <comments> fail:", err)
	}

	err = _do.LeftJoin(&_a, primaryKey.EqCol(_aPK)).Scan(map[string]interface{}{})
	if err != nil {
		t.Error("LeftJoin() on table <comments> fail:", err)
	}

	_, err = _do.Not().Or().Clauses().Take()
	if err != nil {
		t.Error("Not/Or/Clauses on table <comments> fail:", err)
	}
}
