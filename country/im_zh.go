package country

import xlanguage "golang.org/x/text/language"

func init() {
	dataIsleOfMan.RegisterName(xlanguage.Chinese, "马恩岛")
	dataIsleOfMan.RegisterOfficialName(xlanguage.Chinese, "马恩岛")
	dataIsleOfMan.RegisterCapital(xlanguage.Chinese, "道格拉斯")
}
