//go:build country_all || country_asia || country_lb || country_western_asia

package country

import xlanguage "golang.org/x/text/language"

func init() {
	dataLebanon.RegisterName(xlanguage.French, "Liban")
	dataLebanon.RegisterOfficialName(xlanguage.French, "République libanaise")
	dataLebanon.RegisterCapital(xlanguage.French, "Beyrouth")
}
