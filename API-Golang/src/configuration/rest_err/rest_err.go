package rest_err

import "net/http"

// RestErr representa a estrutura padronizada para erros da API.
type RestErr struct {
	Message string   `json:"message"` // Mensagem principal do erro.
	Err     string   `json:"error"`   // Tipo do erro (ex: "bad_request", "internal_server_error").
	Code    int      `json:"code"`    // Código de status HTTP.
	Causes  []Causes `json:"causes"`  // Lista de causas detalhadas do erro
}

// Causes representa a causa específica de um erro, geralmente relacionada a um campo inválido.
type Causes struct {
	Field   string `json:"field"`   // Campo que originou o erro.
	Message string `json:"message"` // Mensagem específica para o campo.
}

func (r *RestErr) Error() string {
	return r.Message
}

// NewRestErr é um construtor para criar um novo RestErr com todos os campos.
func NewRestErr(message, err string, code int, causes []Causes) *RestErr {
	return &RestErr{
		Message: message,
		Err:     err,
		Code:    code,
		Causes:  causes,
	}
}

// NewBadRequestError cria um erro padronizado para "Bad Request" (Requisição Inválida).
func NewBadRequestError(message string) *RestErr {
	return &RestErr{
		Message: message,
		Err:     "bad request",
		Code:    http.StatusBadRequest,
	}
}

func NewUnauthorizedRequestError(message string) *RestErr {
	return &RestErr{
		Message: message,
		Err:     "unauthorized",
		Code:    http.StatusUnauthorized,
	}
}

// NewBadRequestValidationError cria um erro padronizado para falhas de validação.
func NewBadRequestValidationError(message string, causes []Causes) *RestErr {
	return &RestErr{
		Message: message,
		Err:     "bad_request",
		Code:    http.StatusBadRequest,
		Causes:  causes,
	}
}

// NewInternalServerError cria um erro padronizado para "Internal Server Error" (Erro Interno do Servidor).
func NewInternalServerError(message string) *RestErr {
	return &RestErr{
		Message: message,
		Err:     "internal_server_error",
		Code:    http.StatusInternalServerError,
	}
}

// NewNotFoundError cria um erro padronizado para "Not Found" (Não Encontrado).
func NewNotFoundError(message string) *RestErr {
	return &RestErr{
		Message: message,
		Err:     "not_found",
		Code:    http.StatusNotFound,
	}
}

// NewForbiddenError cria um erro padronizado para "Forbidden" (Proibido).
func NewForbiddenError(message string) *RestErr {
	return &RestErr{
		Message: message,
		Err:     "forbidden",
		Code:    http.StatusForbidden,
	}
}
