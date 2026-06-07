//go:build (lang_ja || lang_all) && (country_all || country_asia || country_np || country_southern_asia || currency_all || currency_npr)

package currency

import xlanguage "golang.org/x/text/language"

func init() {
	NPR.RegisterName(xlanguage.Japanese, "ネパール・ルピー")
}
