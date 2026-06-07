package country

import xlanguage "golang.org/x/text/language"

func init() {
	dataSingapore.RegisterName(xlanguage.Chinese, "新加坡")
	dataSingapore.RegisterOfficialName(xlanguage.Chinese, "新加坡共和国")
	dataSingapore.RegisterCapital(xlanguage.Chinese, "新加坡")
}
