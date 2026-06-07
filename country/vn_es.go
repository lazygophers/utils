//go:build lang_es || lang_all

package country

import xlanguage "golang.org/x/text/language"

func init() {
	dataVietnam.RegisterName(xlanguage.Spanish, "Vietnam")
	dataVietnam.RegisterOfficialName(xlanguage.Spanish, "República Socialista de Vietnam")
	dataVietnam.RegisterCapital(xlanguage.Spanish, "Hanói")
}
