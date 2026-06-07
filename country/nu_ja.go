//go:build (lang_ja || lang_all) && (country_all || country_nu || country_oceania || country_polynesia)

package country

import xlanguage "golang.org/x/text/language"

func init() {
	dataNiue.RegisterName(xlanguage.Japanese, "ニウエ")
	dataNiue.RegisterOfficialName(xlanguage.Japanese, "ニウエ")
	dataNiue.RegisterCapital(xlanguage.Japanese, "アロフィ")
}
