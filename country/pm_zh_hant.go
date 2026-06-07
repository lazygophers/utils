//go:build lang_zh_hant || lang_all

package country

import xlanguage "golang.org/x/text/language"

func init() {
	dataSaintPierreAndMiquelon.RegisterName(xlanguage.MustParse("zh-Hant"), "聖皮埃與密克隆群島")
	dataSaintPierreAndMiquelon.RegisterOfficialName(xlanguage.MustParse("zh-Hant"), "聖皮埃與密克隆海外集體")
	dataSaintPierreAndMiquelon.RegisterCapital(xlanguage.MustParse("zh-Hant"), "聖皮埃")
}
