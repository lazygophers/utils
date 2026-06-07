//go:build (lang_fr || lang_all) && (country_all || country_americas || country_caribbean || country_sx)

package country

import xlanguage "golang.org/x/text/language"

func init() {
	dataSintMaarten.RegisterName(xlanguage.French, "Saint-Martin")
	dataSintMaarten.RegisterOfficialName(xlanguage.French, "Saint-Martin")
	dataSintMaarten.RegisterCapital(xlanguage.French, "Philipsburg")
}
