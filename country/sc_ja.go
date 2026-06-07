//go:build (lang_ja || lang_all) && (country_africa || country_all || country_eastern_africa || country_sc)

package country

import xlanguage "golang.org/x/text/language"

func init() {
	dataSeychelles.RegisterName(xlanguage.Japanese, "セーシェル")
	dataSeychelles.RegisterOfficialName(xlanguage.Japanese, "セーシェル共和国")
	dataSeychelles.RegisterCapital(xlanguage.Japanese, "ヴィクトリア")
}
