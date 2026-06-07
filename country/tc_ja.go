//go:build (lang_ja || lang_all) && (country_all || country_americas || country_caribbean || country_tc)

package country

import xlanguage "golang.org/x/text/language"

func init() {
	dataTurksAndCaicosIslands.RegisterName(xlanguage.Japanese, "タークス・カイコス諸島")
	dataTurksAndCaicosIslands.RegisterOfficialName(xlanguage.Japanese, "タークス・カイコス諸島")
	dataTurksAndCaicosIslands.RegisterCapital(xlanguage.Japanese, "コックバーン・タウン")
}
