//go:build lang_ja || lang_all

package country

import xlanguage "golang.org/x/text/language"

func init() {
	dataNicaragua.RegisterName(xlanguage.Japanese, "ニカラグア")
	dataNicaragua.RegisterOfficialName(xlanguage.Japanese, "ニカラグア共和国")
	dataNicaragua.RegisterCapital(xlanguage.Japanese, "マナグア")
}
