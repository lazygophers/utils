//go:build country_africa || country_all || country_tg || country_western_africa

package country

import xlanguage "golang.org/x/text/language"

func init() {
	dataTogo.RegisterName(xlanguage.French, "Togo")
	dataTogo.RegisterOfficialName(xlanguage.French, "République togolaise")
	dataTogo.RegisterCapital(xlanguage.French, "Lomé")
}
