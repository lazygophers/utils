//go:build (lang_ja || lang_all) && (country_all || country_europe || country_is || country_northern_europe || currency_all || currency_isk)

package currency

import xlanguage "golang.org/x/text/language"

func init() {
	ISK.RegisterName(xlanguage.Japanese, "アイスランド・クローナ")
}
