//go:build lang_ar || lang_all

package country

import xlanguage "golang.org/x/text/language"

func init() {
	dataSolomonIslands.RegisterName(xlanguage.Arabic, "جزر سليمان")
	dataSolomonIslands.RegisterOfficialName(xlanguage.Arabic, "جزر سليمان")
	dataSolomonIslands.RegisterCapital(xlanguage.Arabic, "هونيارا")
}
