//go:build lang_ja || lang_all

package country

import xlanguage "golang.org/x/text/language"

func init() {
	dataPhilippines.RegisterName(xlanguage.Japanese, "フィリピン")
	dataPhilippines.RegisterOfficialName(xlanguage.Japanese, "フィリピン共和国")
	dataPhilippines.RegisterCapital(xlanguage.Japanese, "マニラ")
}
