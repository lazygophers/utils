//go:build lang_ja || lang_all

package country

import xlanguage "golang.org/x/text/language"

func init() {
	dataBahrain.RegisterName(xlanguage.Japanese, "バーレーン")
	dataBahrain.RegisterOfficialName(xlanguage.Japanese, "バーレーン王国")
	dataBahrain.RegisterCapital(xlanguage.Japanese, "マナーマ")
}
