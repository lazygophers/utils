//go:build (lang_ar || lang_all) && (country_all || country_asia || country_qa || country_western_asia || currency_all || currency_qar)

package currency

import xlanguage "golang.org/x/text/language"

func init() {
	Qar.RegisterName(xlanguage.Arabic, "ريال قطري")
}
