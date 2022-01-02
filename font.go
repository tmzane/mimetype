package mimetype

// CODE GENERATED AUTOMATICALLY
// THIS FILE MUST NOT BE EDITED BY HAND

const (
	// FontCollection collection mime type.
	FontCollection = "font/collection"
	// FontOtf otf mime type.
	FontOtf = "font/otf"
	// FontSfnt sfnt mime type.
	FontSfnt = "font/sfnt"
	// FontTtf ttf mime type.
	FontTtf = "font/ttf"
	// FontWoff woff mime type.
	FontWoff = "font/woff"
	// FontWoff2 woff2 mime type.
	FontWoff2 = "font/woff2"
)

// IsFont checks if the mime types is font.
func IsFont(mt string) bool {
	switch mt {
	case "font/collection":
		return true
	case "font/otf":
		return true
	case "font/sfnt":
		return true
	case "font/ttf":
		return true
	case "font/woff":
		return true
	case "font/woff2":
		return true
	default:
		return false
	}
}
