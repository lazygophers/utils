//go:build lang_ja || lang_all

package country

import xlanguage "golang.org/x/text/language"

func init() {
	dataIceland.RegisterName(xlanguage.Japanese, "アイスランド")
	dataIceland.RegisterOfficialName(xlanguage.Japanese, "アイスランド")
	dataIceland.RegisterCapital(xlanguage.Japanese, "レイキャヴィーク")
}
