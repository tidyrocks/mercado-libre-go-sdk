package pictures

import (
	"context"
	"fmt"

	"github.com/tidyrocks/mercado-libre-go-sdk/internal/httpx"
)

const (
	baseEndpoint = "https://api.mercadolibre.com"
)

// Upload sube una imagen a los servidores de Mercado Libre.
func Upload(fileContent []byte, filename, accessToken string) (*PictureUploadResponse, error) {
	url := fmt.Sprintf("%s/pictures/items/upload", baseEndpoint)

	var response PictureUploadResponse
	err := httpx.DoMultipartUpload(context.Background(), url, accessToken, fileContent, filename, &response)
	if err != nil {
		return nil, err
	}

	return &response, nil
}

// LinkToItem vincula una imagen existente a un ítem.
func LinkToItem(itemID, pictureID, accessToken string) error {
	url := fmt.Sprintf("%s/items/%s/pictures", baseEndpoint, itemID)

	request := ItemPictureLink{
		ID: pictureID,
	}

	var response interface{} // API no retorna contenido específico
	return httpx.DoPostJSON(context.Background(), url, accessToken, request, &response)
}

// GetErrors obtiene los errores de procesamiento de una imagen.
func GetErrors(pictureID, accessToken string) (*PictureError, error) {
	url := fmt.Sprintf("%s/pictures/%s/errors", baseEndpoint, pictureID)

	var response struct {
		ID     string `json:"id"`
		Source string `json:"source"`
		Error  struct {
			Message string `json:"message"`
		} `json:"error"`
	}

	err := httpx.DoGetJSON(context.Background(), url, accessToken, &response)
	if err != nil {
		return nil, err
	}

	return &PictureError{
		ID:      response.ID,
		Source:  response.Source,
		Message: response.Error.Message,
	}, nil
}

// UpdateItemPictures reemplaza todas las imágenes existentes del ítem.
func UpdateItemPictures(itemID, accessToken string, pictures []PictureRequest) error {
	url := fmt.Sprintf("%s/items/%s", baseEndpoint, itemID)

	request := map[string]interface{}{
		"pictures": pictures,
	}

	var response interface{}
	return httpx.DoPutJSON(context.Background(), url, accessToken, request, &response)
}
