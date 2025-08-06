package pictures

// Picture representa una imagen subida a Mercado Libre.
type Picture struct {
	ID         string             `json:"id"`
	Variations []PictureVariation `json:"variations"`
	Source     *string            `json:"source,omitempty"`
}

// PictureVariation representa una variación de tamaño de imagen.
type PictureVariation struct {
	Size      string `json:"size"`
	URL       string `json:"url"`
	SecureURL string `json:"secure_url"`
	Quality   string `json:"quality"`
}

// PictureRequest representa una imagen en requests.
type PictureRequest struct {
	ID     *string `json:"id,omitempty"`
	Source *string `json:"source,omitempty"`
}

// PictureError representa un error al procesar imagen.
type PictureError struct {
	ID      string `json:"id"`
	Source  string `json:"source"`
	Message string `json:"message"`
}

// PictureUploadResponse representa la respuesta al subir una imagen.
type PictureUploadResponse struct {
	ID         string             `json:"id"`
	Variations []PictureVariation `json:"variations"`
}

// ItemPictureLink representa el request para vincular imagen a ítem.
type ItemPictureLink struct {
	ID string `json:"id"`
}
