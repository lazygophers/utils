//go:build (lang_ja || lang_all) && (country_africa || country_all || country_eastern_africa || country_er)

package country

import xlanguage "golang.org/x/text/language"

func init() {
	dataEritrea.RegisterName(xlanguage.Japanese, "エリトリア")
	dataEritrea.RegisterOfficialName(xlanguage.Japanese, "エリトリア国")
	dataEritrea.RegisterCapital(xlanguage.Japanese, "アスマラ")
}
