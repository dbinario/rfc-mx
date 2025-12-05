package rfc

// Aquí defines:
// - mapa de caracteres → valor numérico para homoclave.
// - mapa de valores → caracteres para homoclave.
// - mapa de caracteres → valor para dígito verificador.
// - lista de palabras prohibidas.

var (
	// TODO: llena con la tabla oficial (Anexo I y II) para homoclave.
	valorCharHomoclave = map[rune]int{
		// ' ' : 00,
		// '&' : 10,
		// 'A' : 11,
		// ...
	}

	charDesdeValorHomoclave = map[int]rune{
		// 0: '1', 1: '2', ..., 33: 'Z'
	}

	// Tabla para dígito verificador (Anexo III).
	valorCharDV = map[rune]int{
		// '0': 0,
		// '1': 1,
		// ...
		// 'A': 10, ...
	}

	// Lista de palabras inconvenientes (ALTISONANTES) para la regla 9.
	palabrasInconvenientes = map[string]struct{}{
		// "BUEI": {},
		// "BUEY": {},
		// "CACA": {},
		// ...
	}
)
