package rfc

import (
	"errors"
	"fmt"
	"regexp"
	"strings"
	"time"
)

// TipoPersona distingue entre física y moral.
type TipoPersona int

const (
	PersonaFisica TipoPersona = iota
	PersonaMoral
)

// Datos para persona física.
type DatosFisica struct {
	Nombre          string
	ApellidoPaterno string
	ApellidoMaterno string // puede ir vacío si no tiene
	FechaNacimiento time.Time
}

// Datos para persona moral.
type DatosMoral struct {
	RazonSocial       string
	FechaConstitucion time.Time
}

// Opciones internas para generación.
type opciones struct {
	incluirHomoclave         bool
	incluirDigitoVerificador bool
}

// Option para configurar la generación.
type Option func(*opciones)

// WithHomoclave indica si se incluye la homoclave (2 caracteres).
func WithHomoclave(on bool) Option {
	return func(o *opciones) { o.incluirHomoclave = on }
}

// WithDigitoVerificador indica si se incluye el dígito verificador final.
func WithDigitoVerificador(on bool) Option {
	return func(o *opciones) { o.incluirDigitoVerificador = on }
}

// Errores comunes.
var (
	ErrNombreInvalido = errors.New("datos de persona física inválidos")
	ErrRazonInvalida  = errors.New("datos de persona moral inválidos")
)

// GeneraRFCFisica genera un RFC estimado para persona física
// siguiendo las reglas publicadas por el SAT.
// Nota: el RFC oficial solo lo asigna la autoridad fiscal.
func GeneraRFCFisica(d DatosFisica, opts ...Option) (string, error) {
	cfg := opciones{
		incluirHomoclave:         true,
		incluirDigitoVerificador: true,
	}
	for _, opt := range opts {
		opt(&cfg)
	}

	if d.Nombre == "" || d.ApellidoPaterno == "" || d.FechaNacimiento.IsZero() {
		return "", ErrNombreInvalido
	}

	norm := normalizarPersonaFisica(d)

	letras := letrasPersonaFisica(norm)
	fecha := formatoFechaRFC(d.FechaNacimiento)

	base := letras + fecha // LLLL + AAMMDD

	if cfg.incluirHomoclave {
		hc := calcularHomoclaveFisica(norm)
		base += hc
	}

	if cfg.incluirDigitoVerificador {
		dv := calcularDigitoVerificador(base)
		base += dv
	}

	return base, nil
}

// GeneraRFCMoral genera un RFC estimado para persona moral.
// Nota: el RFC oficial solo lo asigna la autoridad fiscal.
func GeneraRFCMoral(d DatosMoral, opts ...Option) (string, error) {
	cfg := opciones{
		incluirHomoclave:         true,
		incluirDigitoVerificador: true,
	}
	for _, opt := range opts {
		opt(&cfg)
	}

	if d.RazonSocial == "" || d.FechaConstitucion.IsZero() {
		return "", ErrRazonInvalida
	}

	norm := normalizarPersonaMoral(d)

	letras := letrasPersonaMoral(norm)
	fecha := formatoFechaRFC(d.FechaConstitucion)

	base := letras + fecha // LLL + AAMMDD

	if cfg.incluirHomoclave {
		hc := calcularHomoclaveMoral(norm)
		base += hc
	}

	if cfg.incluirDigitoVerificador {
		dv := calcularDigitoVerificador(base)
		base += dv
	}

	return base, nil
}

// InfoRFC contiene metadatos extraídos de un RFC existente.
type InfoRFC struct {
	RFC            string
	Tipo           TipoPersona
	Fecha          time.Time
	TieneHomoclave bool
}

// EsValidoRFC valida estructura general y dígito verificador.
// No valida contra el SAT, solo consistencia interna.
func EsValidoRFC(rfcStr string) bool {
	_, err := ParseRFC(rfcStr)
	return err == nil
}

// ParseRFC intenta inferir tipo de persona, fecha y DV a partir de un RFC.
func ParseRFC(rfcStr string) (InfoRFC, error) {
	var info InfoRFC

	rfcStr = strings.ToUpper(strings.TrimSpace(rfcStr))
	if len(rfcStr) != 12 && len(rfcStr) != 13 {
		return info, fmt.Errorf("longitud inválida")
	}

	// Regex básicos (puedes afinarlos luego).
	reFis := regexp.MustCompile(`^[A-Z&Ñ]{4}\d{6}[A-Z0-9]{3}$`)
	reMor := regexp.MustCompile(`^[A-Z&Ñ]{3}\d{6}[A-Z0-9]{3}$`)

	isFis := reFis.MatchString(rfcStr)
	isMor := reMor.MatchString(rfcStr)

	if !isFis && !isMor {
		return info, fmt.Errorf("formato no coincide con RFC físico ni moral")
	}

	base := rfcStr[:len(rfcStr)-1]
	dvEsperado := calcularDigitoVerificador(base)
	if !strings.HasSuffix(rfcStr, dvEsperado) {
		return info, fmt.Errorf("dígito verificador inválido")
	}

	info.RFC = rfcStr

	// Extraer fecha
	var fechaStr string
	if isFis {
		info.Tipo = PersonaFisica
		fechaStr = rfcStr[4:10]
	} else {
		info.Tipo = PersonaMoral
		fechaStr = rfcStr[3:9]
	}

	fecha, err := parseFechaRFC(fechaStr)
	if err != nil {
		return info, fmt.Errorf("fecha inválida en RFC: %w", err)
	}
	info.Fecha = fecha
	info.TieneHomoclave = true // en la práctica casi todos la traen

	return info, nil
}
