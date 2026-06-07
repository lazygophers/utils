//go:build lang_ja || lang_all

package country

import xlanguage "golang.org/x/text/language"

func init() {
	dataDominica.RegisterName(xlanguage.Japanese, "ドミニカ国")
	dataDominica.RegisterOfficialName(xlanguage.Japanese, "ドミニカ国")
	dataDominica.RegisterCapital(xlanguage.Japanese, "ロゾー")
}
