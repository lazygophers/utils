//go:build lang_ar || lang_all

package country

import xlanguage "golang.org/x/text/language"

func init() {
	dataBahamas.RegisterName(xlanguage.Arabic, "الباهاما")
	dataBahamas.RegisterOfficialName(xlanguage.Arabic, "كومنولث جزر الباهاما")
	dataBahamas.RegisterCapital(xlanguage.Arabic, "ناسو")
}
