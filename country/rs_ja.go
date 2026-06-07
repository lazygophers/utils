//go:build lang_ja || lang_all

package country

import xlanguage "golang.org/x/text/language"

func init() {
	dataSerbia.RegisterName(xlanguage.Japanese, "セルビア")
	dataSerbia.RegisterOfficialName(xlanguage.Japanese, "セルビア共和国")
	dataSerbia.RegisterCapital(xlanguage.Japanese, "ベオグラード")
}
