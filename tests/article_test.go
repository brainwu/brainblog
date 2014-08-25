package tests

import (
	"testing"
	"brainblog/models"
	. "github.com/smartystreets/goconvey/convey"
	_ "github.com/go-sql-driver/mysql"
)

func Test_List(t *testing.T) {
	Convey("article list", t, func() {
			var article *models.Article = &models.Article{}
			So(article.List, ShouldNotEqual, nil)
		})
}
