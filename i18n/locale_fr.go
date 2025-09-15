//go:build i18n_fr || i18n_all

package i18n

// 注册法语配置
func init() {
	RegisterLocale(French, &Locale{
		Language:     French,
		Region:       "FR",
		Name:         "Français",
		EnglishName:  "French",
		Messages: map[string]string{
			// Messages généraux
			"error":   "Erreur",
			"warning": "Avertissement",
			"info":    "Information",
			"success": "Succès",
			"failed":  "Échec",
			"loading": "Chargement...",
			"saving":  "Enregistrement...",
			"done":    "Terminé",
			"cancel":  "Annuler",
			"confirm": "Confirmer",
			"yes":     "Oui",
			"no":      "Non",
			"ok":      "OK",
			
			// Temps relatif
			"just_now":      "À l'instant",
			"seconds_ago":   "Il y a %d secondes",
			"minutes_ago":   "Il y a %d minutes",
			"hours_ago":     "Il y a %d heures",
			"days_ago":      "Il y a %d jours",
			"weeks_ago":     "Il y a %d semaines",
			"months_ago":    "Il y a %d mois",
			"years_ago":     "Il y a %d ans",
			"seconds_later": "Dans %d secondes",
			"minutes_later": "Dans %d minutes",
			"hours_later":   "Dans %d heures",
			"days_later":    "Dans %d jours",
			"weeks_later":   "Dans %d semaines",
			"months_later":  "Dans %d mois",
			"years_later":   "Dans %d ans",
			
			// Messages d'erreur de validation
			"required":        "%s est requis",
			"email":           "%s doit être une adresse e-mail valide",
			"url":             "%s doit être une URL valide",
			"min":             "La valeur minimale de %s est %s",
			"max":             "La valeur maximale de %s est %s",
			"len":             "La longueur de %s doit être de %s caractères",
			"mobile":          "%s doit être un numéro de téléphone mobile valide",
			"idcard":          "%s doit être un numéro de carte d'identité valide",
			"bankcard":        "%s doit être un numéro de carte bancaire valide",
			"chinese_name":    "%s doit être un nom chinois valide",
			"strong_password": "%s doit être un mot de passe fort (au moins 8 caractères, contenant des majuscules, des minuscules, des chiffres et des caractères spéciaux)",
			
			// Réseau
			"network_error":     "Erreur réseau",
			"connection_failed": "Échec de connexion",
			"timeout":           "Délai d'attente dépassé",
			"server_error":      "Erreur serveur",
			"not_found":         "Non trouvé",
			"unauthorized":      "Non autorisé",
			"forbidden":         "Accès interdit",
			
			// Opérations de fichiers
			"file_not_found":    "Fichier non trouvé",
			"file_read_error":   "Erreur de lecture de fichier",
			"file_write_error":  "Erreur d'écriture de fichier",
			"file_delete_error": "Erreur de suppression de fichier",
			"file_create_error": "Erreur de création de fichier",
			
			// Base de données
			"db_connection_error": "Erreur de connexion à la base de données",
			"db_query_error":      "Erreur de requête de base de données",
			"db_insert_error":     "Erreur d'insertion dans la base de données",
			"db_update_error":     "Erreur de mise à jour de la base de données",
			"db_delete_error":     "Erreur de suppression de la base de données",
			
			// Opérations générales
			"create": "Créer",
			"read":   "Lire",
			"update": "Mettre à jour",
			"delete": "Supprimer",
			"list":   "Liste",
			"search": "Rechercher",
			"filter": "Filtrer",
			"sort":   "Trier",
		},
		Formats: &Formats{
			DateFormat:        "02/01/2006",
			TimeFormat:        "15:04:05",
			DateTimeFormat:    "02/01/2006 15:04:05",
			NumberFormat:      "%.2f",
			CurrencyFormat:    "%.2f €",
			DecimalSeparator:  ",",
			ThousandSeparator: " ",
			Units: &Units{
				ByteUnits:     []string{"o", "Ko", "Mo", "Go", "To", "Po"},
				SpeedUnits:    []string{"o/s", "Ko/s", "Mo/s", "Go/s", "To/s", "Po/s"},
				BitSpeedUnits: []string{"bps", "Kbps", "Mbps", "Gbps", "Tbps", "Pbps"},
				TimeUnits: map[string]string{
					"nanosecond":  "ns",
					"microsecond": "μs",
					"millisecond": "ms",
					"second":      "seconde",
					"minute":      "minute",
					"hour":        "heure",
					"day":         "jour",
					"week":        "semaine",
					"month":       "mois",
					"year":        "année",
					"seconds":     "secondes",
					"minutes":     "minutes",
					"hours":       "heures",
					"days":        "jours",
					"weeks":       "semaines",
					"months":      "mois",
					"years":       "années",
				},
				DistanceUnits: []string{"mm", "cm", "m", "km"},
				WeightUnits:   []string{"g", "kg", "t"},
			},
		},
	})
}