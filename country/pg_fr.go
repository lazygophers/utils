//go:build (lang_fr || lang_all) && (country_all || country_melanesia || country_oceania || country_pg)

package country

import xlanguage "golang.org/x/text/language"

func init() {
	dataPapuaNewGuinea.RegisterName(xlanguage.French, "Papouasie-Nouvelle-Guinée")
	dataPapuaNewGuinea.RegisterOfficialName(xlanguage.French, "État indépendant de Papouasie-Nouvelle-Guinée")
	dataPapuaNewGuinea.RegisterCapital(xlanguage.French, "Port Moresby")
}
