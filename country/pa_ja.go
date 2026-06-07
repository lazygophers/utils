//go:build (lang_ja || lang_all) && (country_all || country_americas || country_central_america || country_pa)

package country

import xlanguage "golang.org/x/text/language"

func init() {
	dataPanama.RegisterName(xlanguage.Japanese, "パナマ")
	dataPanama.RegisterOfficialName(xlanguage.Japanese, "パナマ共和国")
	dataPanama.RegisterCapital(xlanguage.Japanese, "パナマシティ")
}
