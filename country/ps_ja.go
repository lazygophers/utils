//go:build lang_ja || lang_all

package country

import xlanguage "golang.org/x/text/language"

func init() {
	dataPalestine.RegisterName(xlanguage.Japanese, "パレスチナ")
	dataPalestine.RegisterOfficialName(xlanguage.Japanese, "パレスチナ国")
	dataPalestine.RegisterCapital(xlanguage.Japanese, "東エルサレム")
}
