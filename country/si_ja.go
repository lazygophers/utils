//go:build (lang_ja || lang_all) && (country_all || country_europe || country_si || country_southern_europe)

package country

import xlanguage "golang.org/x/text/language"

func init() {
	dataSlovenia.RegisterName(xlanguage.Japanese, "スロベニア")
	dataSlovenia.RegisterOfficialName(xlanguage.Japanese, "スロベニア共和国")
	dataSlovenia.RegisterCapital(xlanguage.Japanese, "リュブリャナ")
}
