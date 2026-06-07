//go:build (lang_ru || lang_all) && (country_all || country_asia || country_il || country_ps || country_western_asia || currency_all || currency_ils)

package currency

import xlanguage "golang.org/x/text/language"

func init() {
	Ils.RegisterName(xlanguage.Russian, "Новый израильский шекель")
}
