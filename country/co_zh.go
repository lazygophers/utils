package country

import xlanguage "golang.org/x/text/language"

func init() {
	dataColombia.RegisterName(xlanguage.Chinese, "哥伦比亚")
	dataColombia.RegisterOfficialName(xlanguage.Chinese, "哥伦比亚共和国")
	dataColombia.RegisterCapital(xlanguage.Chinese, "波哥大")
}
