package country

import xlanguage "golang.org/x/text/language"

func init() {
	dataSaintMartin.RegisterName(xlanguage.English, "Saint Martin")
	dataSaintMartin.RegisterOfficialName(xlanguage.English, "Collectivity of Saint Martin")
	dataSaintMartin.RegisterCapital(xlanguage.English, "Marigot")
}
