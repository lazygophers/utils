//go:build (lang_ja || lang_all) && (country_all || country_am || country_asia || country_western_asia || currency_all || currency_amd)

package currency

import xlanguage "golang.org/x/text/language"

func init() {
	AMD.RegisterName(xlanguage.Japanese, "アルメニア・ドラム")
}
