package error

import "errors"

var (
	ErrInternalServerError = errors.New("internal server error")
	ErrBadRequest          = errors.New("bad request")
	ErrFileNotFound        = errors.New("file not found")
	ErrInvalidJSON         = errors.New("invalid json format")
	ErrUploadError         = errors.New("failed to upload csv file")
	ErrInvalidUploadFile   = errors.New("invalid upload file")
	ErrEmptyFile           = errors.New("uploaded file content is empty")
	ErrEmptyData           = errors.New("data is empty")
	ErrInvalidCSVFormat    = errors.New("file is not in valid csv format")
	ErrCSVParseError       = errors.New("failed to parse csv content")
	ErrInvalidFieldFormat  = errors.New("invalid field format")
)

var GeneralErrors = []error{
	ErrInternalServerError, ErrBadRequest,
	ErrFileNotFound, ErrInvalidJSON, ErrEmptyData,
	ErrUploadError, ErrInvalidUploadFile,
	ErrEmptyFile, ErrInvalidCSVFormat,
	ErrCSVParseError, ErrInvalidFieldFormat}
