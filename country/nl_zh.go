package country

import xlanguage "golang.org/x/text/language"

func init() {
	dataNetherlands.RegisterName(xlanguage.Chinese, "荷兰")
	dataNetherlands.RegisterOfficialName(xlanguage.Chinese, "荷兰王国")
	dataNetherlands.RegisterCapital(xlanguage.Chinese, "阿姆斯特丹")
}
