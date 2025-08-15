package http

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"mime/multipart"
	"net/http"
	"net/url"
)

// DoGetJSON agrega Bearer token si se provee y decodifica respuesta.
func DoGetJSON[T any](ctx context.Context, url, token string, target *T) error {
	req, err := http.NewRequestWithContext(ctx, http.MethodGet, url, nil)
	if err != nil {
		return err
	}
	if token != "" {
		req.Header.Set("Authorization", "Bearer "+token)
	}
	req.Header.Set("Accept", "application/json")

	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return fmt.Errorf("status inesperado: %d", resp.StatusCode)
	}

	return json.NewDecoder(resp.Body).Decode(target)
}

// DoGetJSONWithParams hace GET con query parameters, token opcional y decodifica en target.
func DoGetJSONWithParams[T any](ctx context.Context, baseURL, token string, params url.Values, target *T) error {
	u, err := url.Parse(baseURL)
	if err != nil {
		return err
	}

	if len(params) > 0 {
		query := u.Query()
		for key, values := range params {
			for _, value := range values {
				query.Add(key, value)
			}
		}
		u.RawQuery = query.Encode()
	}

	return DoGetJSON(ctx, u.String(), token, target)
}

// DoPostJSON hace POST con JSON body, token opcional y decodifica en target.
func DoPostJSON[T any](ctx context.Context, url, token string, body interface{}, target *T) error {
	jsonBody, err := json.Marshal(body)
	if err != nil {
		return err
	}

	req, err := http.NewRequestWithContext(ctx, http.MethodPost, url, bytes.NewBuffer(jsonBody))
	if err != nil {
		return err
	}

	if token != "" {
		req.Header.Set("Authorization", "Bearer "+token)
	}
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Accept", "application/json")

	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	if resp.StatusCode < 200 || resp.StatusCode >= 300 {
		return fmt.Errorf("status inesperado: %d", resp.StatusCode)
	}

	return json.NewDecoder(resp.Body).Decode(target)
}

// DoMultipartUpload hace POST multipart/form-data para subir archivos.
func DoMultipartUpload[T any](ctx context.Context, url, token string, fileContent []byte, filename string, target *T) error {
	var buffer bytes.Buffer
	writer := multipart.NewWriter(&buffer)

	// Crear el campo file
	part, err := writer.CreateFormFile("file", filename)
	if err != nil {
		return err
	}

	// Escribir el contenido del archivo
	_, err = part.Write(fileContent)
	if err != nil {
		return err
	}

	// Cerrar el writer para finalizar el boundary
	err = writer.Close()
	if err != nil {
		return err
	}

	req, err := http.NewRequestWithContext(ctx, http.MethodPost, url, &buffer)
	if err != nil {
		return err
	}

	if token != "" {
		req.Header.Set("Authorization", "Bearer "+token)
	}
	req.Header.Set("Content-Type", writer.FormDataContentType())

	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	if resp.StatusCode < 200 || resp.StatusCode >= 300 {
		return fmt.Errorf("status inesperado: %d", resp.StatusCode)
	}

	return json.NewDecoder(resp.Body).Decode(target)
}

// DoPutJSON hace PUT con JSON body, token opcional y decodifica en target.
func DoPutJSON[T any](ctx context.Context, url, token string, body interface{}, target *T) error {
	jsonBody, err := json.Marshal(body)
	if err != nil {
		return err
	}

	req, err := http.NewRequestWithContext(ctx, http.MethodPut, url, bytes.NewBuffer(jsonBody))
	if err != nil {
		return err
	}

	if token != "" {
		req.Header.Set("Authorization", "Bearer "+token)
	}
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Accept", "application/json")

	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	if resp.StatusCode < 200 || resp.StatusCode >= 300 {
		return fmt.Errorf("status inesperado: %d", resp.StatusCode)
	}

	return json.NewDecoder(resp.Body).Decode(target)
}