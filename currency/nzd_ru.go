//go:build (lang_ru || lang_all) && (country_all || country_australia_and_new_zealand || country_ck || country_nu || country_nz || country_oceania || country_pn || country_polynesia || country_tk || currency_all || currency_nzd)

package currency

import xlanguage "golang.org/x/text/language"

func init() {
	Nzd.RegisterName(xlanguage.Russian, "Новозеландский доллар")
}
