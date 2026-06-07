package country

import xlanguage "golang.org/x/text/language"

func init() {
	dataSuriname.RegisterName(xlanguage.Chinese, "苏里南")
	dataSuriname.RegisterOfficialName(xlanguage.Chinese, "苏里南共和国")
	dataSuriname.RegisterCapital(xlanguage.Chinese, "帕拉马里博")
}
