package manage

type ManageController struct {
	baseController
}

func (mc *ManageController) Manage() {
	mc.TplNames="manage/manage.html"
}
