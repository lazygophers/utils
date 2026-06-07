//go:build (lang_ja || lang_all) && (country_all || country_americas || country_gl || country_northern_america)

package country

import xlanguage "golang.org/x/text/language"

func init() {
	dataGreenland.RegisterName(xlanguage.Japanese, "グリーンランド")
	dataGreenland.RegisterOfficialName(xlanguage.Japanese, "グリーンランド")
	dataGreenland.RegisterCapital(xlanguage.Japanese, "ヌーク")
}
