//go:build lang_ja || lang_all

package country

import xlanguage "golang.org/x/text/language"

func init() {
	dataBouvetIsland.RegisterName(xlanguage.Japanese, "ブーベ島")
	dataBouvetIsland.RegisterOfficialName(xlanguage.Japanese, "ブーベ島")
}
