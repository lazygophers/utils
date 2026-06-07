//go:build (lang_ja || lang_all) && (country_africa || country_all || country_ga || country_middle_africa)

package country

import xlanguage "golang.org/x/text/language"

func init() {
	dataGabon.RegisterName(xlanguage.Japanese, "ガボン")
	dataGabon.RegisterOfficialName(xlanguage.Japanese, "ガボン共和国")
	dataGabon.RegisterCapital(xlanguage.Japanese, "リーブルヴィル")
}
