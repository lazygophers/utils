//go:build (lang_ja || lang_all) && (country_africa || country_all || country_eastern_africa || country_ug || currency_all || currency_ugx)

package currency

import xlanguage "golang.org/x/text/language"

func init() {
	UGX.RegisterName(xlanguage.Japanese, "ウガンダ・シリング")
}
