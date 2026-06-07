//go:build (lang_ja || lang_all) && (country_all || country_ee || country_europe || country_northern_europe)

package country

import xlanguage "golang.org/x/text/language"

func init() {
	dataEstonia.RegisterName(xlanguage.Japanese, "エストニア")
	dataEstonia.RegisterOfficialName(xlanguage.Japanese, "エストニア共和国")
	dataEstonia.RegisterCapital(xlanguage.Japanese, "タリン")
}
