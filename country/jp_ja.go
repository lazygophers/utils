package country

import xlanguage "golang.org/x/text/language"

func init() {
	dataJapan.RegisterName(xlanguage.Japanese, "日本")
	dataJapan.RegisterOfficialName(xlanguage.Japanese, "日本国")
	dataJapan.RegisterCapital(xlanguage.Japanese, "東京")
}
