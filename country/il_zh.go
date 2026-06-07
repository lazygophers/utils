package country

import xlanguage "golang.org/x/text/language"

func init() {
	dataIsrael.RegisterName(xlanguage.Chinese, "以色列")
	dataIsrael.RegisterOfficialName(xlanguage.Chinese, "以色列国")
	dataIsrael.RegisterCapital(xlanguage.Chinese, "耶路撒冷")
}
