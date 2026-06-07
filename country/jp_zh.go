package country

import xlanguage "golang.org/x/text/language"

func init() {
	dataJapan.RegisterName(xlanguage.Chinese, "日本")
	dataJapan.RegisterOfficialName(xlanguage.Chinese, "日本国")
	dataJapan.RegisterCapital(xlanguage.Chinese, "东京")
}
