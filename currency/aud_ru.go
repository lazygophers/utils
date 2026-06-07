//go:build (lang_ru || lang_all) && (country_all || country_antarctic || country_au || country_australia_and_new_zealand || country_cc || country_cx || country_hm || country_ki || country_micronesia || country_nf || country_nr || country_oceania || country_polynesia || country_tv || currency_all || currency_aud)

package currency

import xlanguage "golang.org/x/text/language"

func init() {
	Aud.RegisterName(xlanguage.Russian, "Австралийский доллар")
}
