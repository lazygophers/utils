package country

import xlanguage "golang.org/x/text/language"

func init() {
	dataPoland.RegisterName(xlanguage.Chinese, "波兰")
	dataPoland.RegisterOfficialName(xlanguage.Chinese, "波兰共和国")
	dataPoland.RegisterCapital(xlanguage.Chinese, "华沙")
}
