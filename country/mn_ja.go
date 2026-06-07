//go:build lang_ja || lang_all

package country

import xlanguage "golang.org/x/text/language"

func init() {
	dataMongolia.RegisterName(xlanguage.Japanese, "モンゴル国")
	dataMongolia.RegisterOfficialName(xlanguage.Japanese, "モンゴル国")
	dataMongolia.RegisterCapital(xlanguage.Japanese, "ウランバートル")
}
