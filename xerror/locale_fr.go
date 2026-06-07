//go:build lang_fr || lang_all

package xerror

// 法文内置错误码翻译；按 build tag 选择启用。
func init() {
	registerBuiltinLocale("fr", map[int]string{
		CodeSystem:        "erreur système",
		CodeInvalidParam:  "paramètre invalide",
		CodeNoAuth:        "non autorisé",
		CodeNoData:        "introuvable",
		CodeConflict:      "conflit",
		CodeNotLogin:      "non connecté",
		CodeTimeout:       "délai dépassé",
		CodeRateLimited:   "trop de requêtes",
		CodeForbidden:     "interdit",
		CodeUnavailable:   "service indisponible",
		CodeDataCorrupted: "données corrompues",
	})
}
