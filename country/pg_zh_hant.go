//go:build lang_zh_hant || lang_all

package country

import xlanguage "golang.org/x/text/language"

func init() {
	dataPapuaNewGuinea.RegisterName(xlanguage.MustParse("zh-Hant"), "巴布亞紐幾內亞")
	dataPapuaNewGuinea.RegisterOfficialName(xlanguage.MustParse("zh-Hant"), "巴布亞紐幾內亞獨立國")
	dataPapuaNewGuinea.RegisterCapital(xlanguage.MustParse("zh-Hant"), "摩斯比港")
}
