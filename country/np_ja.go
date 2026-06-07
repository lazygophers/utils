//go:build (lang_ja || lang_all) && (country_all || country_asia || country_np || country_southern_asia)

package country

import xlanguage "golang.org/x/text/language"

func init() {
	dataNepal.RegisterName(xlanguage.Japanese, "ネパール")
	dataNepal.RegisterOfficialName(xlanguage.Japanese, "ネパール連邦民主共和国")
	dataNepal.RegisterCapital(xlanguage.Japanese, "カトマンズ")
}
