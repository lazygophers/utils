//go:build lang_ko || lang_all

package country

import xlanguage "golang.org/x/text/language"

func init() {
	dataPalestine.RegisterName(xlanguage.Korean, "팔레스타인")
	dataPalestine.RegisterOfficialName(xlanguage.Korean, "팔레스타인국")
	dataPalestine.RegisterCapital(xlanguage.Korean, "라말라")
}
