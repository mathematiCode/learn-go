package other

import "github.com/lucasb-eyer/go-colorful"

func MaxContrast(color string) (string, error) {
	c, err := colorful.Hex(color)
	if err != nil {
		return "", err
	}
	// Calculate luminance (per ITU-R BT.709)
	luminance := 0.2126*c.R + 0.7152*c.G + 0.0722*c.B
	if luminance > 0.5 {
		return "#000000", nil // dark text on light background
	}
	return "#ffffff", nil // light text on dark background
}
