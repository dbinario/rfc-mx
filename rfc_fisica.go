package rfc

// letrasPersonaFisica arma las primeras 4 letras del RFC.
func letrasPersonaFisica(p personaFisicaNormalizada) string {
	// TODO:
	// - aplicar regla base (1a letra AP + 1a vocal interna AP + 1a letra AM + 1a letra nombre)
	// - manejar apellidos cortos (1–2 letras)
	// - manejar falta de AM
	// - aplicar regla de palabras inconvenientes y sustituir 4a por X si aplica.
	return "XXXX" // placeholder
}

// Sugerido en: rfc_moral.go

// letrasPersonaMoral arma las 3 primeras letras del RFC de moral.
func letrasPersonaMoral(p personaMoralNormalizada) string {
	// TODO:
	// - contar palabras válidas
	// - aplicar reglas de 3 palabras (1+1+1), 2 palabras (1+2), 1 palabra (3 letras)
	// - considerar números convertidos a texto si te quieres rifar hardcore.
	return "XXX" // placeholder
}
