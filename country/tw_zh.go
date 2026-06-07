package country

import xlanguage "golang.org/x/text/language"

func init() {
	dataTaiwan.RegisterName(xlanguage.Chinese, "台湾")
	dataTaiwan.RegisterOfficialName(xlanguage.Chinese, "中华民国（台湾）")
	dataTaiwan.RegisterCapital(xlanguage.Chinese, "台北")
}
