package country

import xlanguage "golang.org/x/text/language"

func init() {
	dataChile.RegisterName(xlanguage.Chinese, "智利")
	dataChile.RegisterOfficialName(xlanguage.Chinese, "智利共和国")
	dataChile.RegisterCapital(xlanguage.Chinese, "圣地亚哥")
}
