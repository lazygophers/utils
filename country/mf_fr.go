//go:build country_all || country_americas || country_caribbean || country_mf

package country

import xlanguage "golang.org/x/text/language"

func init() {
	dataSaintMartin.RegisterName(xlanguage.French, "Saint-Martin")
	dataSaintMartin.RegisterOfficialName(xlanguage.French, "Collectivité de Saint-Martin")
	dataSaintMartin.RegisterCapital(xlanguage.French, "Marigot")
}
