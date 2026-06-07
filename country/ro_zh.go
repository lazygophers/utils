package country

import xlanguage "golang.org/x/text/language"

func init() {
	dataRomania.RegisterName(xlanguage.Chinese, "罗马尼亚")
	dataRomania.RegisterOfficialName(xlanguage.Chinese, "罗马尼亚")
	dataRomania.RegisterCapital(xlanguage.Chinese, "布加勒斯特")
}
