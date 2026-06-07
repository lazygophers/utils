//go:build lang_fr || lang_all

package country

import xlanguage "golang.org/x/text/language"

func init() {
	dataMoldova.RegisterName(xlanguage.French, "Moldavie")
	dataMoldova.RegisterOfficialName(xlanguage.French, "République de Moldavie")
	dataMoldova.RegisterCapital(xlanguage.French, "Chișinău")
}
