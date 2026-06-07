package country

import xlanguage "golang.org/x/text/language"

func init() {
	dataRussia.RegisterName(xlanguage.Chinese, "俄罗斯")
	dataRussia.RegisterOfficialName(xlanguage.Chinese, "俄罗斯联邦")
	dataRussia.RegisterCapital(xlanguage.Chinese, "莫斯科")
}
