package country

import xlanguage "golang.org/x/text/language"

func init() {
	dataFrance.RegisterName(xlanguage.Chinese, "法国")
	dataFrance.RegisterOfficialName(xlanguage.Chinese, "法兰西共和国")
	dataFrance.RegisterCapital(xlanguage.Chinese, "巴黎")
}
