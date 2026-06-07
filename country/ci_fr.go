package country

import xlanguage "golang.org/x/text/language"

func init() {
	dataIvoryCoast.RegisterName(xlanguage.French, "Côte d'Ivoire")
	dataIvoryCoast.RegisterOfficialName(xlanguage.French, "République de Côte d'Ivoire")
	dataIvoryCoast.RegisterCapital(xlanguage.French, "Yamoussoukro")
}
