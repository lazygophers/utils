//go:build lang_ja || lang_all

package country

import xlanguage "golang.org/x/text/language"

func init() {
	dataAntarctica.RegisterName(xlanguage.Japanese, "南極大陸")
	dataAntarctica.RegisterOfficialName(xlanguage.Japanese, "南極大陸")
}
