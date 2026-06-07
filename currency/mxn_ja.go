//go:build (lang_ja || lang_all) && (country_all || country_americas || country_central_america || country_mx || currency_all || currency_mxn)

package currency

import xlanguage "golang.org/x/text/language"

func init() {
	Mxn.RegisterName(xlanguage.Japanese, "メキシコ・ペソ")
}
