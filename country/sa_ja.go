//go:build (lang_ja || lang_all) && (country_all || country_asia || country_sa || country_western_asia)

package country

import xlanguage "golang.org/x/text/language"

func init() {
	dataSaudiArabia.RegisterName(xlanguage.Japanese, "サウジアラビア")
	dataSaudiArabia.RegisterOfficialName(xlanguage.Japanese, "サウジアラビア王国")
	dataSaudiArabia.RegisterCapital(xlanguage.Japanese, "リヤド")
}
