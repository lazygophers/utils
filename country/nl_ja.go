//go:build (lang_ja || lang_all) && (country_all || country_europe || country_nl || country_western_europe)

package country

import xlanguage "golang.org/x/text/language"

func init() {
	dataNetherlands.RegisterName(xlanguage.Japanese, "オランダ")
	dataNetherlands.RegisterOfficialName(xlanguage.Japanese, "オランダ王国")
	dataNetherlands.RegisterCapital(xlanguage.Japanese, "アムステルダム")
}
