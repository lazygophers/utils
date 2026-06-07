//go:build (lang_ar || lang_all) && (country_all || country_asia || country_il || country_ps || country_western_asia || currency_all || currency_ils)

package currency

import xlanguage "golang.org/x/text/language"

func init() {
	ILS.RegisterName(xlanguage.Arabic, "شيكل إسرائيلي جديد")
}
