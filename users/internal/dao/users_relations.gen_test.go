// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.

package dao

import (
	"context"
	"fmt"
	"testing"

	"github.com/intxff/mango/users/internal/models"
	"gorm.io/gen"
	"gorm.io/gen/field"
	"gorm.io/gorm/clause"
)

func init() {
	InitializeDB()
	err := db.AutoMigrate(&models.UsersRelation{})
	if err != nil {
		fmt.Printf("Error: AutoMigrate(&models.UsersRelation{}) fail: %s", err)
	}
}

func Test_usersRelationQuery(t *testing.T) {
	usersRelation := newUsersRelation(db)
	usersRelation = *usersRelation.As(usersRelation.TableName())
	_do := usersRelation.WithContext(context.Background()).Debug()

	primaryKey := field.NewString(usersRelation.TableName(), clause.PrimaryKey)
	_, err := _do.Unscoped().Where(primaryKey.IsNotNull()).Delete()
	if err != nil {
		t.Error("clean table <users_relations> fail:", err)
		return
	}

	_, ok := usersRelation.GetFieldByName("")
	if ok {
		t.Error("GetFieldByName(\"\") from usersRelation success")
	}

	err = _do.Create(&models.UsersRelation{})
	if err != nil {
		t.Error("create item in table <users_relations> fail:", err)
	}

	err = _do.Save(&models.UsersRelation{})
	if err != nil {
		t.Error("create item in table <users_relations> fail:", err)
	}

	err = _do.CreateInBatches([]*models.UsersRelation{{}, {}}, 10)
	if err != nil {
		t.Error("create item in table <users_relations> fail:", err)
	}

	_, err = _do.Select(usersRelation.ALL).Take()
	if err != nil {
		t.Error("Take() on table <users_relations> fail:", err)
	}

	_, err = _do.First()
	if err != nil {
		t.Error("First() on table <users_relations> fail:", err)
	}

	_, err = _do.Last()
	if err != nil {
		t.Error("First() on table <users_relations> fail:", err)
	}

	_, err = _do.Where(primaryKey.IsNotNull()).FindInBatch(10, func(tx gen.Dao, batch int) error { return nil })
	if err != nil {
		t.Error("FindInBatch() on table <users_relations> fail:", err)
	}

	err = _do.Where(primaryKey.IsNotNull()).FindInBatches(&[]*models.UsersRelation{}, 10, func(tx gen.Dao, batch int) error { return nil })
	if err != nil {
		t.Error("FindInBatches() on table <users_relations> fail:", err)
	}

	_, err = _do.Select(usersRelation.ALL).Where(primaryKey.IsNotNull()).Order(primaryKey.Desc()).Find()
	if err != nil {
		t.Error("Find() on table <users_relations> fail:", err)
	}

	_, err = _do.Distinct(primaryKey).Take()
	if err != nil {
		t.Error("select Distinct() on table <users_relations> fail:", err)
	}

	_, err = _do.Select(usersRelation.ALL).Omit(primaryKey).Take()
	if err != nil {
		t.Error("Omit() on table <users_relations> fail:", err)
	}

	_, err = _do.Group(primaryKey).Find()
	if err != nil {
		t.Error("Group() on table <users_relations> fail:", err)
	}

	_, err = _do.Scopes(func(dao gen.Dao) gen.Dao { return dao.Where(primaryKey.IsNotNull()) }).Find()
	if err != nil {
		t.Error("Scopes() on table <users_relations> fail:", err)
	}

	_, _, err = _do.FindByPage(0, 1)
	if err != nil {
		t.Error("FindByPage() on table <users_relations> fail:", err)
	}

	_, err = _do.ScanByPage(&models.UsersRelation{}, 0, 1)
	if err != nil {
		t.Error("ScanByPage() on table <users_relations> fail:", err)
	}

	_, err = _do.Attrs(primaryKey).Assign(primaryKey).FirstOrInit()
	if err != nil {
		t.Error("FirstOrInit() on table <users_relations> fail:", err)
	}

	_, err = _do.Attrs(primaryKey).Assign(primaryKey).FirstOrCreate()
	if err != nil {
		t.Error("FirstOrCreate() on table <users_relations> fail:", err)
	}

	var _a _another
	var _aPK = field.NewString(_a.TableName(), clause.PrimaryKey)

	err = _do.Join(&_a, primaryKey.EqCol(_aPK)).Scan(map[string]interface{}{})
	if err != nil {
		t.Error("Join() on table <users_relations> fail:", err)
	}

	err = _do.LeftJoin(&_a, primaryKey.EqCol(_aPK)).Scan(map[string]interface{}{})
	if err != nil {
		t.Error("LeftJoin() on table <users_relations> fail:", err)
	}

	_, err = _do.Not().Or().Clauses().Take()
	if err != nil {
		t.Error("Not/Or/Clauses on table <users_relations> fail:", err)
	}
}