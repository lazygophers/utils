//go:build (lang_ar || lang_all) && (country_africa || country_all || country_gn || country_western_africa || currency_all || currency_gnf)

package currency

import xlanguage "golang.org/x/text/language"

func init() {
	GNF.RegisterName(xlanguage.Arabic, "فرنك غيني")
}
