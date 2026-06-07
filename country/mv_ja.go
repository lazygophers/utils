//go:build (lang_ja || lang_all) && (country_all || country_asia || country_mv || country_southern_asia)

package country

import xlanguage "golang.org/x/text/language"

func init() {
	dataMaldives.RegisterName(xlanguage.Japanese, "モルディブ")
	dataMaldives.RegisterOfficialName(xlanguage.Japanese, "モルディブ共和国")
	dataMaldives.RegisterCapital(xlanguage.Japanese, "マレ")
}
