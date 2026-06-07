//go:build lang_ja || lang_all

package country

import xlanguage "golang.org/x/text/language"

func init() {
	dataAmericanSamoa.RegisterName(xlanguage.Japanese, "アメリカ領サモア")
	dataAmericanSamoa.RegisterOfficialName(xlanguage.Japanese, "アメリカ領サモア")
	dataAmericanSamoa.RegisterCapital(xlanguage.Japanese, "パゴパゴ")
}
