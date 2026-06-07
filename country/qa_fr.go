//go:build (lang_fr || lang_all) && (country_all || country_asia || country_qa || country_western_asia)

package country

import xlanguage "golang.org/x/text/language"

func init() {
	dataQatar.RegisterName(xlanguage.French, "Qatar")
	dataQatar.RegisterOfficialName(xlanguage.French, "État du Qatar")
	dataQatar.RegisterCapital(xlanguage.French, "Doha")
}
