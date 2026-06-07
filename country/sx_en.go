//go:build country_all || country_americas || country_caribbean || country_sx

package country

import xlanguage "golang.org/x/text/language"

func init() {
	dataSintMaarten.RegisterName(xlanguage.English, "Sint Maarten")
	dataSintMaarten.RegisterOfficialName(xlanguage.English, "Sint Maarten")
	dataSintMaarten.RegisterCapital(xlanguage.English, "Philipsburg")
}
