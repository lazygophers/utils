//go:build (lang_ja || lang_all) && (country_all || country_asia || country_bt || country_southern_asia || currency_all || currency_btn)

package currency

import xlanguage "golang.org/x/text/language"

func init() {
	Btn.RegisterName(xlanguage.Japanese, "ニュルタム")
}
