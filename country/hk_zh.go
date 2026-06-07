package country

import xlanguage "golang.org/x/text/language"

func init() {
	dataHongKong.RegisterName(xlanguage.Chinese, "香港")
	dataHongKong.RegisterOfficialName(xlanguage.Chinese, "中华人民共和国香港特别行政区")
	dataHongKong.RegisterCapital(xlanguage.Chinese, "香港")
}
