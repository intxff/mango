package main

import (
	"os"
	"strings"

	"gorm.io/driver/mysql"
	"gorm.io/gen"
	"gorm.io/gorm"
)

var (
	dbs map[string][]string = map[string][]string{
		"mango_users": {
			"users",
			"users_relations",
		},
		"mango_categories": {
			"categories",
		},
		"mango_tags": {
			"tags",
		},
		"mango_posts": {
			"posts",
			"posts_tags",
			"posts_reactions",
			"posts_priority",
		},
		"mango_comments": {
			"comments",
			"comments_reactions",
		},
		"mango_notifications": {
			"notifications",
			"notifications_lazy",
		},
		"mango_validations": {
			"validations",
			"validations_email",
		},
	}
)

func main() {
	for database, tables := range dbs {
        t := strings.Split(database, "_")
        // clean previous built files
        modelPath := "../" + t[1] + "/internal/models/"
        daoPath := "../" + t[1] + "/internal/dao/"
        os.RemoveAll(modelPath)
        os.RemoveAll(daoPath)

		g := gen.NewGenerator(gen.Config{
			ModelPkgPath:      modelPath,
			OutPath:           daoPath,
			Mode:              gen.WithDefaultQuery | gen.WithQueryInterface,
			FieldNullable:     false,
			FieldCoverable:    false,
			FieldSignable:     true,
			FieldWithIndexTag: true,
			FieldWithTypeTag:  true,
			WithUnitTest:      true,
		})
		link := "root:root321@(127.0.0.1:10000)/" + database + "?charset=utf8mb4&parseTime=True&loc=Local"
		gormdb, _ := gorm.Open(mysql.Open(link))
		g.UseDB(gormdb)
		// generate models
		models := make([]any, 0, len(tables))
		for _, table := range tables {
			models = append(models, g.GenerateModel(table))
		}
		g.ApplyBasic(
			models...,
		)
		g.Execute()
	}
}
