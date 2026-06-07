//go:build (lang_ja || lang_all) && (country_all || country_americas || country_caribbean || country_tt || currency_all || currency_ttd)

package currency

import xlanguage "golang.org/x/text/language"

func init() {
	TTD.RegisterName(xlanguage.Japanese, "トリニダード・トバゴ・ドル")
}
