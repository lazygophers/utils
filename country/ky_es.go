//go:build lang_es || lang_all

package country

import xlanguage "golang.org/x/text/language"

func init() {
	dataCaymanIslands.RegisterName(xlanguage.Spanish, "Islas Caimán")
	dataCaymanIslands.RegisterOfficialName(xlanguage.Spanish, "Islas Caimán")
	dataCaymanIslands.RegisterCapital(xlanguage.Spanish, "George Town")
}
