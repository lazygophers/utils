//go:build lang_es || lang_all

package xerror

// 西班牙文内置错误码翻译；按 build tag 选择启用。
func init() {
	registerBuiltinLocale("es", map[int]string{
		CodeSystem:        "error del sistema",
		CodeInvalidParam:  "parámetro inválido",
		CodeNoAuth:        "no autorizado",
		CodeNoData:        "no encontrado",
		CodeConflict:      "conflicto",
		CodeNotLogin:      "no ha iniciado sesión",
		CodeTimeout:       "tiempo de espera agotado",
		CodeRateLimited:   "demasiadas solicitudes",
		CodeForbidden:     "prohibido",
		CodeUnavailable:   "servicio no disponible",
		CodeDataCorrupted: "datos dañados",
	})
}
