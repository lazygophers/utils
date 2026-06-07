package country

import xlanguage "golang.org/x/text/language"

func init() {
	dataSaintBarthelemy.RegisterName(xlanguage.French, "Saint-Barthélemy")
	dataSaintBarthelemy.RegisterOfficialName(xlanguage.French, "Collectivité de Saint-Barthélemy")
	dataSaintBarthelemy.RegisterCapital(xlanguage.French, "Gustavia")
}
