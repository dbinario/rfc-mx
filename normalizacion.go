package rfc

import "strings"

// personaFisicaNormalizada contiene nombres ya limpios para el algoritmo.
type personaFisicaNormalizada struct {
	Nombre          string
	ApellidoPaterno string
	ApellidoMaterno string
	PalabrasNombre  []string
	PalabrasAP      []string
	PalabrasAM      []string
}

// personaMoralNormalizada para razón social.
type personaMoralNormalizada struct {
	RazonSocial     string
	PalabrasValidas []string
}

// TODO: rellena esto después con:
// - quitar acentos
// - upper
// - eliminar partículas (DE, DEL, LA, LOS, Y...)
// - manejar MARIA / JOSE
func normalizarPersonaFisica(d DatosFisica) personaFisicaNormalizada {

	n := personaFisicaNormalizada{
		Nombre:          strings.ToUpper(strings.TrimSpace(d.Nombre)),
		ApellidoPaterno: strings.ToUpper(strings.TrimSpace(d.ApellidoPaterno)),
		ApellidoMaterno: strings.ToUpper(strings.TrimSpace(d.ApellidoMaterno)),
	}

	// TODO: quitar acentos, particulas y aplicar reglas especiales.
	// TODO: llenar PalabrasNombre, PalabrasAP, PalabrasAM.

	return n
}

// TODO: aquí igual limpias razón social, quitas "S.A.", "DE C.V.", "SOCIEDAD", etc.
func normalizarPersonaMoral(d DatosMoral) personaMoralNormalizada {
	rs := strings.ToUpper(strings.TrimSpace(d.RazonSocial))

	n := personaMoralNormalizada{
		RazonSocial: rs,
		// TODO: split + quitar partículas + palabras ignoradas
	}
	return n
}
