//go:build lang_ja || lang_all

package country

import xlanguage "golang.org/x/text/language"

func init() {
	dataPapuaNewGuinea.RegisterName(xlanguage.Japanese, "パプアニューギニア")
	dataPapuaNewGuinea.RegisterOfficialName(xlanguage.Japanese, "パプアニューギニア独立国")
	dataPapuaNewGuinea.RegisterCapital(xlanguage.Japanese, "ポートモレスビー")
}
