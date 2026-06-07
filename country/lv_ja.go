//go:build lang_ja || lang_all

package country

import xlanguage "golang.org/x/text/language"

func init() {
	dataLatvia.RegisterName(xlanguage.Japanese, "ラトビア")
	dataLatvia.RegisterOfficialName(xlanguage.Japanese, "ラトビア共和国")
	dataLatvia.RegisterCapital(xlanguage.Japanese, "リガ")
}
