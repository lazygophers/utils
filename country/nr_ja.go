//go:build (lang_ja || lang_all) && (country_all || country_micronesia || country_nr || country_oceania)

package country

import xlanguage "golang.org/x/text/language"

func init() {
	dataNauru.RegisterName(xlanguage.Japanese, "ナウル")
	dataNauru.RegisterOfficialName(xlanguage.Japanese, "ナウル共和国")
	dataNauru.RegisterCapital(xlanguage.Japanese, "ヤレン")
}
