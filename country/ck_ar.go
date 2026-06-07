//go:build lang_ar || lang_all

package country

import xlanguage "golang.org/x/text/language"

func init() {
	dataCookIslands.RegisterName(xlanguage.Arabic, "جزر كوك")
	dataCookIslands.RegisterOfficialName(xlanguage.Arabic, "جزر كوك")
	dataCookIslands.RegisterCapital(xlanguage.Arabic, "أفاروا")
}
