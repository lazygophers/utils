//go:build i18n_pt || i18n_all

package i18n

// 注册葡萄牙语配置
func init() {
	RegisterLocale(Portuguese, &Locale{
		Language:     Portuguese,
		Region:       "PT",
		Name:         "Português",
		EnglishName:  "Portuguese",
		Messages: map[string]string{
			// Mensagens gerais
			"error":   "Erro",
			"warning": "Aviso",
			"info":    "Informação",
			"success": "Sucesso",
			"failed":  "Falhou",
			"loading": "Carregando...",
			"saving":  "Salvando...",
			"done":    "Concluído",
			"cancel":  "Cancelar",
			"confirm": "Confirmar",
			"yes":     "Sim",
			"no":      "Não",
			"ok":      "OK",
			
			// Tempo relativo
			"just_now":      "Agora mesmo",
			"seconds_ago":   "Há %d segundos",
			"minutes_ago":   "Há %d minutos",
			"hours_ago":     "Há %d horas",
			"days_ago":      "Há %d dias",
			"weeks_ago":     "Há %d semanas",
			"months_ago":    "Há %d meses",
			"years_ago":     "Há %d anos",
			"seconds_later": "Em %d segundos",
			"minutes_later": "Em %d minutos",
			"hours_later":   "Em %d horas",
			"days_later":    "Em %d dias",
			"weeks_later":   "Em %d semanas",
			"months_later":  "Em %d meses",
			"years_later":   "Em %d anos",
			
			// Mensagens de erro de validação
			"required":        "%s é obrigatório",
			"email":           "%s deve ser um endereço de e-mail válido",
			"url":             "%s deve ser uma URL válida",
			"min":             "O valor mínimo de %s é %s",
			"max":             "O valor máximo de %s é %s",
			"len":             "O comprimento de %s deve ser de %s caracteres",
			"mobile":          "%s deve ser um número de telefone móvel válido",
			"idcard":          "%s deve ser um número de documento de identidade válido",
			"bankcard":        "%s deve ser um número de cartão bancário válido",
			"chinese_name":    "%s deve ser um nome chinês válido",
			"strong_password": "%s deve ser uma senha forte (pelo menos 8 caracteres, contendo letras maiúsculas, minúsculas, números e caracteres especiais)",
			
			// Rede
			"network_error":     "Erro de rede",
			"connection_failed": "Falha na conexão",
			"timeout":           "Tempo esgotado",
			"server_error":      "Erro do servidor",
			"not_found":         "Não encontrado",
			"unauthorized":      "Não autorizado",
			"forbidden":         "Acesso proibido",
			
			// Operações de arquivo
			"file_not_found":    "Arquivo não encontrado",
			"file_read_error":   "Erro de leitura do arquivo",
			"file_write_error":  "Erro de escrita do arquivo",
			"file_delete_error": "Erro de exclusão do arquivo",
			"file_create_error": "Erro de criação do arquivo",
			
			// Banco de dados
			"db_connection_error": "Erro de conexão com o banco de dados",
			"db_query_error":      "Erro de consulta do banco de dados",
			"db_insert_error":     "Erro de inserção no banco de dados",
			"db_update_error":     "Erro de atualização do banco de dados",
			"db_delete_error":     "Erro de exclusão do banco de dados",
			
			// Operações gerais
			"create": "Criar",
			"read":   "Ler",
			"update": "Atualizar",
			"delete": "Excluir",
			"list":   "Lista",
			"search": "Pesquisar",
			"filter": "Filtrar",
			"sort":   "Classificar",
		},
		Formats: &Formats{
			DateFormat:        "02/01/2006",
			TimeFormat:        "15:04:05",
			DateTimeFormat:    "02/01/2006 15:04:05",
			NumberFormat:      "%.2f",
			CurrencyFormat:    "€ %.2f",
			DecimalSeparator:  ",",
			ThousandSeparator: ".",
			Units: &Units{
				ByteUnits:     []string{"B", "KB", "MB", "GB", "TB", "PB"},
				SpeedUnits:    []string{"B/s", "KB/s", "MB/s", "GB/s", "TB/s", "PB/s"},
				BitSpeedUnits: []string{"bps", "Kbps", "Mbps", "Gbps", "Tbps", "Pbps"},
				TimeUnits: map[string]string{
					"nanosecond":  "ns",
					"microsecond": "μs",
					"millisecond": "ms",
					"second":      "segundo",
					"minute":      "minuto",
					"hour":        "hora",
					"day":         "dia",
					"week":        "semana",
					"month":       "mês",
					"year":        "ano",
					"seconds":     "segundos",
					"minutes":     "minutos",
					"hours":       "horas",
					"days":        "dias",
					"weeks":       "semanas",
					"months":      "meses",
					"years":       "anos",
				},
				DistanceUnits: []string{"mm", "cm", "m", "km"},
				WeightUnits:   []string{"g", "kg", "t"},
			},
		},
	})
}