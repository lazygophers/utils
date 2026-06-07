//go:build (lang_ja || lang_all) && (country_all || country_americas || country_caribbean || country_do)

package country

import xlanguage "golang.org/x/text/language"

func init() {
	dataDominicanRepublic.RegisterName(xlanguage.Japanese, "ドミニカ共和国")
	dataDominicanRepublic.RegisterOfficialName(xlanguage.Japanese, "ドミニカ共和国")
	dataDominicanRepublic.RegisterCapital(xlanguage.Japanese, "サントドミンゴ")
}
