//go:build (lang_ja || lang_all) && (country_all || country_au || country_australia_and_new_zealand || country_oceania)

package country

import xlanguage "golang.org/x/text/language"

func init() {
	dataAustralia.RegisterName(xlanguage.Japanese, "オーストラリア")
	dataAustralia.RegisterOfficialName(xlanguage.Japanese, "オーストラリア連邦")
	dataAustralia.RegisterCapital(xlanguage.Japanese, "キャンベラ")
}
