package validation

import (
	"encoding/json"
	"errors"
	"strings"

	"github.com/RafaelFleitas/API-Golang/src/configuration/rest_err"
	"github.com/gin-gonic/gin/binding"
	"github.com/go-playground/locales/en"
	ut "github.com/go-playground/universal-translator"
	"github.com/go-playground/validator/v10"
	en_translation "github.com/go-playground/validator/v10/translations/en"
)

var (
	transl ut.Translator
)

// A função init é executada automaticamente quando o pacote é importado pela primeira vez.
// Ela configura o tradutor para as mensagens de erro de validação.
func init() {
	// Verifica se o motor de validação do Gin é do tipo *validator.Validate.
	if val, ok := binding.Validator.Engine().(*validator.Validate); ok {
		// Cria um localizador para o idioma inglês.
		en := en.New()
		// Cria um tradutor universal com o inglês como fallback.
		unt := ut.New(en, en)

		// Obtém o tradutor específico para o idioma "en" (inglês).
		transl, _ = unt.GetTranslator("en")
		// Registra as traduções padrão em inglês para o validador.
		en_translation.RegisterDefaultTranslations(val, transl)

		val.RegisterValidation("notblank", func(fl validator.FieldLevel) bool {
			return strings.TrimSpace(fl.Field().String()) != ""
		})

	}
}

// ValidateUserError recebe um erro de validação e o converte em um erro de API padronizado (RestErr).
func ValidateUserError(validation_err error) *rest_err.RestErr {
	var jsonErr *json.UnmarshalTypeError
	var jsonValidationError validator.ValidationErrors

	// Verifica se o erro é do tipo json.UnmarshalTypeError (ex: tipo de dado inválido no JSON).
	if errors.As(validation_err, &jsonErr) {
		return rest_err.NewBadRequestError("Invalid field type")
		// Verifica se o erro é do tipo validator.ValidationErrors (ex: campo obrigatório faltando, formato de email inválido).
	} else if errors.As(validation_err, &jsonValidationError) {
		// Cria uma lista para armazenar as causas detalhadas do erro.
		errorsCauses := []rest_err.Causes{}

		// Itera sobre cada erro de validação encontrado.
		for _, e := range validation_err.(validator.ValidationErrors) {
			// Cria uma causa de erro com a mensagem traduzida e o nome do campo.
			cause := rest_err.Causes{
				Message: e.Translate(transl), // Traduz o erro para uma mensagem amigável (ex: "Field is required").
				Field:   e.Field(),           // Obtém o nome do campo que falhou na validação.
			}

			// Adiciona a causa à lista de causas.
			errorsCauses = append(errorsCauses, cause)
		}

		// Retorna um erro de validação de "Bad Request" com a lista detalhada de causas.
		return rest_err.NewBadRequestValidationError("Some fields are invalid", errorsCauses)
	} else {
		// Se o erro não for de nenhum dos tipos esperados, retorna um erro genérico.
		return rest_err.NewBadRequestError("Error trying to convert fields")

	}
}
