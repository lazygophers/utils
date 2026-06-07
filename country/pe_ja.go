//go:build (lang_ja || lang_all) && (country_all || country_americas || country_pe || country_south_america)

package country

import xlanguage "golang.org/x/text/language"

func init() {
	dataPeru.RegisterName(xlanguage.Japanese, "ペルー")
	dataPeru.RegisterOfficialName(xlanguage.Japanese, "ペルー共和国")
	dataPeru.RegisterCapital(xlanguage.Japanese, "リマ")
}
