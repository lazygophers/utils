//go:build lang_fr || lang_all

package country

import xlanguage "golang.org/x/text/language"

func init() {
	dataVietnam.RegisterName(xlanguage.French, "Viêt Nam")
	dataVietnam.RegisterOfficialName(xlanguage.French, "République socialiste du Viêt Nam")
	dataVietnam.RegisterCapital(xlanguage.French, "Hanoï")
}
