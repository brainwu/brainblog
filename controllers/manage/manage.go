package manage

type ManageController struct {
	baseController
}

func (mc *ManageController) Manage() {
	mc.Layout = "manage/manage.html"
	mc.TplNames="manage/send_article.html"
}
