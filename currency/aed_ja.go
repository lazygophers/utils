//go:build (lang_ja || lang_all) && (country_ae || country_all || country_asia || country_western_asia || currency_aed || currency_all)

package currency

import xlanguage "golang.org/x/text/language"

func init() {
	AED.RegisterName(xlanguage.Japanese, "UAEディルハム")
}
