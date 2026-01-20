package utils

import "net/http"

// ErrorCode representa un código único de error
type ErrorCode string

// Códigos de error de autenticación (100-199)
const (
	ErrCodePasswordMismatch ErrorCode = "ERR_101" // Las contraseñas no coinciden
	ErrCodeInvalidPassword  ErrorCode = "ERR_102" // Contraseña incorrecta
	ErrCodeWeakPassword     ErrorCode = "ERR_103" // Contraseña muy débil
	ErrCodeEmailExists      ErrorCode = "ERR_104" // Email ya registrado
	ErrCodeUserNotFound     ErrorCode = "ERR_105" // Usuario no encontrado
	ErrCodeInvalidEmail     ErrorCode = "ERR_106" // Email inválido
	ErrCodeTokenExpired     ErrorCode = "ERR_107" // Token expirado
	ErrCodeTokenInvalid     ErrorCode = "ERR_108" // Token inválido
	ErrCodeUnauthorized     ErrorCode = "ERR_109" // No autorizado
)

// Códigos de error de validación (200-299)
const (
	ErrCodeInvalidInput    ErrorCode = "ERR_201" // Datos de entrada inválidos
	ErrCodeMissingField    ErrorCode = "ERR_202" // Campo requerido faltante
	ErrCodeInvalidAmount   ErrorCode = "ERR_203" // Monto inválido
	ErrCodeInvalidDate     ErrorCode = "ERR_204" // Fecha inválida
	ErrCodeInvalidCategory ErrorCode = "ERR_205" // Categoría inválida
	ErrCodeInvalidCurrency ErrorCode = "ERR_206" // Moneda inválida
)

// Códigos de error de negocio (300-399)
const (
	ErrCodeInsufficientFunds ErrorCode = "ERR_301" // Fondos insuficientes
	ErrCodeBudgetExceeded    ErrorCode = "ERR_302" // Presupuesto excedido
	ErrCodeDuplicateEntry    ErrorCode = "ERR_303" // Entrada duplicada
	ErrCodeNotFound          ErrorCode = "ERR_304" // Recurso no encontrado
	ErrCodeAlreadyExists     ErrorCode = "ERR_305" // Recurso ya existe
	ErrCodeCannotDelete      ErrorCode = "ERR_306" // No se puede eliminar
)

// Códigos de error del servidor (500-599)
const (
	ErrCodeDatabaseError  ErrorCode = "ERR_501" // Error de base de datos
	ErrCodeInternalError  ErrorCode = "ERR_502" // Error interno del servidor
	ErrCodeServiceUnavail ErrorCode = "ERR_503" // Servicio no disponible
)

// AppError representa un error de aplicación con código
type AppError struct {
	Code       ErrorCode `json:"code"`
	Message    string    `json:"message"`
	HTTPStatus int       `json:"-"`
}

// Error implementa la interfaz error
func (e *AppError) Error() string {
	return e.Message
}

// NewAppError crea un nuevo error de aplicación
func NewAppError(code ErrorCode, message string, httpStatus int) *AppError {
	return &AppError{
		Code:       code,
		Message:    message,
		HTTPStatus: httpStatus,
	}
}

// Errores predefinidos comunes
var (
	ErrPasswordMismatch = NewAppError(
		ErrCodePasswordMismatch,
		"Passwords do not match",
		http.StatusBadRequest,
	)

	ErrInvalidPassword = NewAppError(
		ErrCodeInvalidPassword,
		"Invalid password",
		http.StatusUnauthorized,
	)

	ErrWeakPassword = NewAppError(
		ErrCodeWeakPassword,
		"Password is too weak",
		http.StatusBadRequest,
	)

	ErrEmailExists = NewAppError(
		ErrCodeEmailExists,
		"Email already registered",
		http.StatusConflict,
	)

	ErrUserNotFound = NewAppError(
		ErrCodeUserNotFound,
		"User not found",
		http.StatusNotFound,
	)

	ErrInvalidEmail = NewAppError(
		ErrCodeInvalidEmail,
		"Invalid email format",
		http.StatusBadRequest,
	)

	ErrUnauthorized = NewAppError(
		ErrCodeUnauthorized,
		"Unauthorized access",
		http.StatusUnauthorized,
	)

	ErrInvalidInput = NewAppError(
		ErrCodeInvalidInput,
		"Invalid input data",
		http.StatusBadRequest,
	)

	ErrNotFound = NewAppError(
		ErrCodeNotFound,
		"Resource not found",
		http.StatusNotFound,
	)

	ErrInternalError = NewAppError(
		ErrCodeInternalError,
		"Internal server error",
		http.StatusInternalServerError,
	)
)
