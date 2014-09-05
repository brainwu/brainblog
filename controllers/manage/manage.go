package manage

type ManageController struct {
	baseController
}

func (mc *ManageController) Manage() {
	mc.Data["name"] = "Brain Wu"
	mc.Layout = "manage/manage.html"
	mc.TplNames="manage/send_article.html"
}
