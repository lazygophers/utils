//go:build (lang_ja || lang_all) && (country_all || country_asia || country_kh || country_south_eastern_asia)

package country

import xlanguage "golang.org/x/text/language"

func init() {
	dataCambodia.RegisterName(xlanguage.Japanese, "カンボジア")
	dataCambodia.RegisterOfficialName(xlanguage.Japanese, "カンボジア王国")
	dataCambodia.RegisterCapital(xlanguage.Japanese, "プノンペン")
}
