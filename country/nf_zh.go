package country

import xlanguage "golang.org/x/text/language"

func init() {
	dataNorfolkIsland.RegisterName(xlanguage.Chinese, "诺福克岛")
	dataNorfolkIsland.RegisterOfficialName(xlanguage.Chinese, "诺福克岛领地")
	dataNorfolkIsland.RegisterCapital(xlanguage.Chinese, "金斯敦")
}
