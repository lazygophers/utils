//go:build (lang_ja || lang_all) && (country_all || country_asia || country_eastern_asia || country_kp)

package country

import xlanguage "golang.org/x/text/language"

func init() {
	dataNorthKorea.RegisterName(xlanguage.Japanese, "朝鮮民主主義人民共和国")
	dataNorthKorea.RegisterOfficialName(xlanguage.Japanese, "朝鮮民主主義人民共和国")
	dataNorthKorea.RegisterCapital(xlanguage.Japanese, "平壌")
}
