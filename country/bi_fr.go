//go:build country_africa || country_all || country_bi || country_eastern_africa

package country

import xlanguage "golang.org/x/text/language"

func init() {
	dataBurundi.RegisterName(xlanguage.French, "Burundi")
	dataBurundi.RegisterOfficialName(xlanguage.French, "République du Burundi")
	dataBurundi.RegisterCapital(xlanguage.French, "Gitega")
}
