//go:build (lang_ja || lang_all) && (country_africa || country_all || country_northern_africa || country_sd)

package country

import xlanguage "golang.org/x/text/language"

func init() {
	dataSudan.RegisterName(xlanguage.Japanese, "スーダン")
	dataSudan.RegisterOfficialName(xlanguage.Japanese, "スーダン共和国")
	dataSudan.RegisterCapital(xlanguage.Japanese, "ハルツーム")
}
