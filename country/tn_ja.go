//go:build lang_ja || lang_all

package country

import xlanguage "golang.org/x/text/language"

func init() {
	dataTunisia.RegisterName(xlanguage.Japanese, "チュニジア")
	dataTunisia.RegisterOfficialName(xlanguage.Japanese, "チュニジア共和国")
	dataTunisia.RegisterCapital(xlanguage.Japanese, "チュニス")
}
