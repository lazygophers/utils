//go:build (lang_ja || lang_all) && (country_all || country_australia_and_new_zealand || country_nf || country_oceania)

package country

import xlanguage "golang.org/x/text/language"

func init() {
	dataNorfolkIsland.RegisterName(xlanguage.Japanese, "ノーフォーク島")
	dataNorfolkIsland.RegisterOfficialName(xlanguage.Japanese, "ノーフォーク島")
	dataNorfolkIsland.RegisterCapital(xlanguage.Japanese, "キングストン")
}
