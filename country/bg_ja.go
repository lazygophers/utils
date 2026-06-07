//go:build lang_ja || lang_all

package country

import xlanguage "golang.org/x/text/language"

func init() {
	dataBulgaria.RegisterName(xlanguage.Japanese, "ブルガリア")
	dataBulgaria.RegisterOfficialName(xlanguage.Japanese, "ブルガリア共和国")
	dataBulgaria.RegisterCapital(xlanguage.Japanese, "ソフィア")
}
