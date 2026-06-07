//go:build lang_ja || lang_all

package country

import xlanguage "golang.org/x/text/language"

func init() {
	dataNamibia.RegisterName(xlanguage.Japanese, "ナミビア")
	dataNamibia.RegisterOfficialName(xlanguage.Japanese, "ナミビア共和国")
	dataNamibia.RegisterCapital(xlanguage.Japanese, "ウィントフック")
}
