//go:build (lang_ja || lang_all) && (country_all || country_am || country_asia || country_western_asia)

package country

import xlanguage "golang.org/x/text/language"

func init() {
	dataArmenia.RegisterName(xlanguage.Japanese, "アルメニア")
	dataArmenia.RegisterOfficialName(xlanguage.Japanese, "アルメニア共和国")
	dataArmenia.RegisterCapital(xlanguage.Japanese, "エレバン")
}
