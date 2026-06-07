package country

import xlanguage "golang.org/x/text/language"

func init() {
	dataSouthKorea.RegisterName(xlanguage.Chinese, "韩国")
	dataSouthKorea.RegisterOfficialName(xlanguage.Chinese, "大韩民国")
	dataSouthKorea.RegisterCapital(xlanguage.Chinese, "首尔")
}
