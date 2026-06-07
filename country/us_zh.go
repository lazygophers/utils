package country

import xlanguage "golang.org/x/text/language"

func init() {
	dataUnitedStates.RegisterName(xlanguage.Chinese, "美国")
	dataUnitedStates.RegisterOfficialName(xlanguage.Chinese, "美利坚合众国")
	dataUnitedStates.RegisterCapital(xlanguage.Chinese, "华盛顿哥伦比亚特区")
}
