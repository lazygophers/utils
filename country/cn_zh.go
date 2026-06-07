package country

import xlanguage "golang.org/x/text/language"

func init() {
	dataChina.RegisterName(xlanguage.Chinese, "中国")
	dataChina.RegisterOfficialName(xlanguage.Chinese, "中华人民共和国")
	dataChina.RegisterCapital(xlanguage.Chinese, "北京")
}
