//go:build (lang_ja || lang_all) && (country_all || country_dk || country_europe || country_northern_europe)

package country

import xlanguage "golang.org/x/text/language"

func init() {
	dataDenmark.RegisterName(xlanguage.Japanese, "デンマーク")
	dataDenmark.RegisterOfficialName(xlanguage.Japanese, "デンマーク王国")
	dataDenmark.RegisterCapital(xlanguage.Japanese, "コペンハーゲン")
}
