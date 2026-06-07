package country

import xlanguage "golang.org/x/text/language"

func init() {
	dataNewZealand.RegisterName(xlanguage.Chinese, "新西兰")
	dataNewZealand.RegisterOfficialName(xlanguage.Chinese, "新西兰")
	dataNewZealand.RegisterCapital(xlanguage.Chinese, "惠灵顿")
}
