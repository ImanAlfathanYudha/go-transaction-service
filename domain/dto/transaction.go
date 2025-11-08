package dto

import "mime/multipart"

// Request DTO for POST /upload
type UploadRequest struct {
	File   multipart.File `json:"-"`
	Header *multipart.FileHeader
}

// Response DTO for POST /upload
type UploadResponse struct {
	Message string `json:"message"`
	Count   int    `json:"count"`
}
