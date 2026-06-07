//go:build lang_ar || lang_all

package country

import xlanguage "golang.org/x/text/language"

func init() {
	dataFalklandIslands.RegisterName(xlanguage.Arabic, "جزر فوكلاند")
	dataFalklandIslands.RegisterOfficialName(xlanguage.Arabic, "إقليم جزر فوكلاند البريطاني فيما وراء البحار")
	dataFalklandIslands.RegisterCapital(xlanguage.Arabic, "ستانلي")
}
