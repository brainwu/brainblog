package tests

import (
	"testing"
	. "github.com/smartystreets/goconvey/convey"
	_ "github.com/go-sql-driver/mysql"
	"github.com/brainwu/brainblog/models"
)

func Test_Insert(t *testing.T) {
	Convey("user insert", t, func() {
			var user *models.User = &models.User{Name:"testinsert", Account:"brainwu1215@gmail.com"}
			qs := user.Query().Filter("Name", "testinsert")
			if !qs.Exist() {
				So(user.Insert(), ShouldEqual, nil)
			}
		})
}


