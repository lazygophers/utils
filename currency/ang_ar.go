//go:build (lang_ar || lang_all) && (country_all || country_americas || country_caribbean || country_cw || country_sx || currency_all || currency_ang)

package currency

import xlanguage "golang.org/x/text/language"

func init() {
	ANG.RegisterName(xlanguage.Arabic, "غيلدر أنتيلي هولندي")
}
