//go:build lang_zh_hant || lang_all

package country

import xlanguage "golang.org/x/text/language"

func init() {
	dataUkraine.RegisterName(xlanguage.MustParse("zh-Hant"), "烏克蘭")
	dataUkraine.RegisterOfficialName(xlanguage.MustParse("zh-Hant"), "烏克蘭")
	dataUkraine.RegisterCapital(xlanguage.MustParse("zh-Hant"), "基輔")
}
