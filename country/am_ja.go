//go:build lang_ja || lang_all

package country

import xlanguage "golang.org/x/text/language"

func init() {
	dataArmenia.RegisterName(xlanguage.Japanese, "アルメニア")
	dataArmenia.RegisterOfficialName(xlanguage.Japanese, "アルメニア共和国")
	dataArmenia.RegisterCapital(xlanguage.Japanese, "エレバン")
}
