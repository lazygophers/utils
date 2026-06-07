//go:build (lang_ja || lang_all) && (country_all || country_melanesia || country_nc || country_oceania)

package country

import xlanguage "golang.org/x/text/language"

func init() {
	dataNewCaledonia.RegisterName(xlanguage.Japanese, "ニューカレドニア")
	dataNewCaledonia.RegisterOfficialName(xlanguage.Japanese, "ニューカレドニア")
	dataNewCaledonia.RegisterCapital(xlanguage.Japanese, "ヌメア")
}
