//go:build (lang_fr || lang_all) && (country_all || country_asia || country_np || country_southern_asia)

package country

import xlanguage "golang.org/x/text/language"

func init() {
	dataNepal.RegisterName(xlanguage.French, "Népal")
	dataNepal.RegisterOfficialName(xlanguage.French, "République fédérale démocratique du Népal")
	dataNepal.RegisterCapital(xlanguage.French, "Katmandou")
}
