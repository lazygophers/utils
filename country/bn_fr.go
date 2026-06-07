//go:build lang_fr || lang_all

package country

import xlanguage "golang.org/x/text/language"

func init() {
	dataBrunei.RegisterName(xlanguage.French, "Brunei")
	dataBrunei.RegisterOfficialName(xlanguage.French, "Brunei Darussalam")
	dataBrunei.RegisterCapital(xlanguage.French, "Bandar Seri Begawan")
}
