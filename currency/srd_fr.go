//go:build (lang_fr || lang_all) && (country_all || country_americas || country_south_america || country_sr || currency_all || currency_srd)

package currency

import xlanguage "golang.org/x/text/language"

func init() {
	SRD.RegisterName(xlanguage.French, "Dollar surinamais")
}
