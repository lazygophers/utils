//go:build (lang_ja || lang_all) && (country_all || country_europe || country_mk || country_southern_europe || currency_all || currency_mkd)

package currency

import xlanguage "golang.org/x/text/language"

func init() {
	MKD.RegisterName(xlanguage.Japanese, "マケドニア・デナール")
}
