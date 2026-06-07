package country

import xlanguage "golang.org/x/text/language"

func init() {
	dataMalaysia.RegisterName(xlanguage.Chinese, "马来西亚")
	dataMalaysia.RegisterOfficialName(xlanguage.Chinese, "马来西亚")
	dataMalaysia.RegisterCapital(xlanguage.Chinese, "吉隆坡")
}
