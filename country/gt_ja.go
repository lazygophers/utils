//go:build lang_ja || lang_all

package country

import xlanguage "golang.org/x/text/language"

func init() {
	dataGuatemala.RegisterName(xlanguage.Japanese, "グアテマラ")
	dataGuatemala.RegisterOfficialName(xlanguage.Japanese, "グアテマラ共和国")
	dataGuatemala.RegisterCapital(xlanguage.Japanese, "グアテマラシティ")
}
