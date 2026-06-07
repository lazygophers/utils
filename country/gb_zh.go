package country

import xlanguage "golang.org/x/text/language"

func init() {
	dataUnitedKingdom.RegisterName(xlanguage.Chinese, "英国")
	dataUnitedKingdom.RegisterOfficialName(xlanguage.Chinese, "大不列颠及北爱尔兰联合王国")
	dataUnitedKingdom.RegisterCapital(xlanguage.Chinese, "伦敦")
}
