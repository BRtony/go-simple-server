package handlers

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"os"
	"path/filepath"
	"time"

	"go.uber.org/zap"
)

type ImageHandler struct {
	logger *zap.Logger
}

func NewImageHandler(logger *zap.Logger) *ImageHandler {
	return &ImageHandler{
		logger: logger,
	}
}

// @Summary      Upload de imagem
// @Description  Faz upload de uma imagem para o servidor
// @Tags         images
// @Accept       multipart/form-data
// @Produce      json
// @Param        file formData file true "Arquivo de imagem"
// @Success      200  {object}  map[string]string
// @Failure      400  {object}  string
// @Failure      500  {object}  string
// @Router       /images/upload [post]
func (h *ImageHandler) UploadImage(w http.ResponseWriter, r *http.Request) {
	// Limitar o tamanho do upload para 10MB
	r.ParseMultipartForm(10 << 20)

	file, handler, err := r.FormFile("file")
	if err != nil {
		h.logger.Error("failed to get file from form",
			zap.Error(err),
		)
		http.Error(w, "Failed to get file from form", http.StatusBadRequest)
		return
	}
	defer file.Close()

	// Verificar se é uma imagem
	contentType := handler.Header.Get("Content-Type")
	if contentType != "image/jpeg" && contentType != "image/png" && contentType != "image/gif" {
		http.Error(w, "Invalid file type. Only JPEG, PNG and GIF are allowed", http.StatusBadRequest)
		return
	}

	// Gerar nome único para o arquivo
	ext := filepath.Ext(handler.Filename)
	filename := fmt.Sprintf("%d%s", time.Now().UnixNano(), ext)
	filepath := filepath.Join("static", "images", filename)

	// Criar arquivo no servidor
	out, err := os.Create(filepath)
	if err != nil {
		h.logger.Error("failed to create file",
			zap.Error(err),
		)
		http.Error(w, "Failed to save file", http.StatusInternalServerError)
		return
	}
	defer out.Close()

	// Copiar conteúdo do arquivo
	_, err = io.Copy(out, file)
	if err != nil {
		h.logger.Error("failed to copy file",
			zap.Error(err),
		)
		http.Error(w, "Failed to save file", http.StatusInternalServerError)
		return
	}

	// Retornar URL da imagem
	response := map[string]string{
		"url": fmt.Sprintf("/images/%s", filename),
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(response)
}
