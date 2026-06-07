//go:build (lang_ja || lang_all) && (country_africa || country_all || country_eastern_africa || country_ss)

package country

import xlanguage "golang.org/x/text/language"

func init() {
	dataSouthSudan.RegisterName(xlanguage.Japanese, "南スーダン")
	dataSouthSudan.RegisterOfficialName(xlanguage.Japanese, "南スーダン共和国")
	dataSouthSudan.RegisterCapital(xlanguage.Japanese, "ジュバ")
}
