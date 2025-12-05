package rfc

import (
	"fmt"
	"time"
)

// formatoFechaRFC convierte time.Time → "AAMMDD".
func formatoFechaRFC(t time.Time) string {
	year := t.Year() % 100
	month := int(t.Month())
	day := t.Day()
	return fmt.Sprintf("%02d%02d%02d", year, month, day)
}

// parseFechaRFC interpreta "AAMMDD" → time.Time (siglo heurístico).
func parseFechaRFC(s string) (time.Time, error) {
	if len(s) != 6 {
		return time.Time{}, fmt.Errorf("longitud fecha != 6")
	}
	yy := s[0:2]
	mm := s[2:4]
	dd := s[4:6]

	var yearBase int
	// Heurística: si yy >= 50 → 1900+, si no → 2000+
	// Puedes ajustar esto si quieres algo más estricto.
	if yy[0] >= '5' {
		yearBase = 1900
	} else {
		yearBase = 2000
	}

	var y, m, d int
	fmt.Sscanf(yy, "%02d", &y)
	fmt.Sscanf(mm, "%02d", &m)
	fmt.Sscanf(dd, "%02d", &d)

	return time.Date(yearBase+y, time.Month(m), d, 0, 0, 0, 0, time.UTC), nil
}
