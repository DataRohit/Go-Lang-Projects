package product

import (
	"errors"
	"net/http"
	"time"

	"github.com/datarohit/go-dynamodb-crud-api/internal/controllers/product"
	EntityProduct "github.com/datarohit/go-dynamodb-crud-api/internal/entities/product"
	"github.com/datarohit/go-dynamodb-crud-api/internal/handlers"
	"github.com/datarohit/go-dynamodb-crud-api/internal/repository/adapter"
	Rules "github.com/datarohit/go-dynamodb-crud-api/internal/rules"
	RulesProduct "github.com/datarohit/go-dynamodb-crud-api/internal/rules/product"
	HttpStatus "github.com/datarohit/go-dynamodb-crud-api/utils/http"
	"github.com/datarohit/go-dynamodb-crud-api/utils/logger"
	"github.com/go-chi/chi"
	"github.com/google/uuid"
	"go.uber.org/zap"
)

type Handler struct {
	handlers.Interface
	Controller product.Interface
	Rules      Rules.Interface
}

func NewHandler(repository adapter.Interface) handlers.Interface {
	return &Handler{
		Controller: product.NewController(repository),
		Rules:      RulesProduct.NewRules(),
	}
}

func (h *Handler) Get(w http.ResponseWriter, r *http.Request) {
	if chi.URLParam(r, "ID") != "" {
		h.getOne(w, r)
	} else {
		h.getAll(w, r)
	}
}

func (h *Handler) getOne(w http.ResponseWriter, r *http.Request) {
	ID, err := uuid.Parse(chi.URLParam(r, "ID"))
	if err != nil {
		logger.GetLogger().Error("Invalid UUID format",
			zap.String("method", r.Method),
			zap.String("url", r.URL.Path),
			zap.String("remote_addr", r.RemoteAddr),
			zap.String("id", chi.URLParam(r, "ID")),
			zap.String("message", "ID is not uuid valid"),
			zap.Time("timestamp", time.Now()),
		)
		HttpStatus.StatusBadRequest(w, errors.New("ID is not uuid valid"))
		return
	}

	response, err := h.Controller.ListOne(ID)
	if err != nil {
		logger.GetLogger().Error("Failed to get product",
			zap.String("method", r.Method),
			zap.String("url", r.URL.Path),
			zap.String("remote_addr", r.RemoteAddr),
			zap.String("id", ID.String()),
			zap.Error(err),
			zap.Time("timestamp", time.Now()),
		)
		HttpStatus.StatusInternalServerError(w, err)
		return
	}

	logger.GetLogger().Info("Successfully retrieved product",
		zap.String("method", r.Method),
		zap.String("url", r.URL.Path),
		zap.String("remote_addr", r.RemoteAddr),
		zap.String("id", ID.String()),
		zap.Any("response", response),
		zap.Time("timestamp", time.Now()),
	)
	HttpStatus.StatusOK(w, response)
}

func (h *Handler) getAll(w http.ResponseWriter, r *http.Request) {
	response, err := h.Controller.ListAll()
	if err != nil {
		logger.GetLogger().Error("Failed to retrieve all products",
			zap.String("method", r.Method),
			zap.String("url", r.URL.Path),
			zap.String("remote_addr", r.RemoteAddr),
			zap.Error(err),
			zap.Time("timestamp", time.Now()),
		)
		HttpStatus.StatusInternalServerError(w, err)
		return
	}

	logger.GetLogger().Info("Successfully retrieved all products",
		zap.String("method", r.Method),
		zap.String("url", r.URL.Path),
		zap.String("remote_addr", r.RemoteAddr),
		zap.Any("response", response),
		zap.Time("timestamp", time.Now()),
	)
	HttpStatus.StatusOK(w, response)
}

func (h *Handler) Post(w http.ResponseWriter, r *http.Request) {
	productBody, err := h.getBodyAndValidate(r, uuid.Nil)
	if err != nil {
		logger.GetLogger().Error("Failed to parse and validate product body",
			zap.String("method", r.Method),
			zap.String("url", r.URL.Path),
			zap.String("remote_addr", r.RemoteAddr),
			zap.Error(err),
			zap.Time("timestamp", time.Now()),
		)
		HttpStatus.StatusBadRequest(w, err)
		return
	}

	ID, err := h.Controller.Create(productBody)
	if err != nil {
		logger.GetLogger().Error("Failed to create product",
			zap.String("method", r.Method),
			zap.String("url", r.URL.Path),
			zap.String("remote_addr", r.RemoteAddr),
			zap.Any("product", productBody),
			zap.Error(err),
			zap.Time("timestamp", time.Now()),
		)
		HttpStatus.StatusInternalServerError(w, err)
		return
	}

	logger.GetLogger().Info("Successfully created product",
		zap.String("method", r.Method),
		zap.String("url", r.URL.Path),
		zap.String("remote_addr", r.RemoteAddr),
		zap.String("id", ID.String()),
		zap.Time("timestamp", time.Now()),
	)
	HttpStatus.StatusOK(w, map[string]interface{}{"id": ID.String()})
}

func (h *Handler) Put(w http.ResponseWriter, r *http.Request) {
	ID, err := uuid.Parse(chi.URLParam(r, "ID"))
	if err != nil {
		logger.GetLogger().Error("Invalid UUID format for update",
			zap.String("method", r.Method),
			zap.String("url", r.URL.Path),
			zap.String("remote_addr", r.RemoteAddr),
			zap.String("id", chi.URLParam(r, "ID")),
			zap.String("message", "ID is not uuid valid"),
			zap.Time("timestamp", time.Now()),
		)
		HttpStatus.StatusBadRequest(w, errors.New("ID is not uuid valid"))
		return
	}

	productBody, err := h.getBodyAndValidate(r, ID)
	if err != nil {
		logger.GetLogger().Error("Failed to parse and validate product body for update",
			zap.String("method", r.Method),
			zap.String("url", r.URL.Path),
			zap.String("remote_addr", r.RemoteAddr),
			zap.Error(err),
			zap.Time("timestamp", time.Now()),
		)
		HttpStatus.StatusBadRequest(w, err)
		return
	}

	if err := h.Controller.Update(ID, productBody); err != nil {
		logger.GetLogger().Error("Failed to update product",
			zap.String("method", r.Method),
			zap.String("url", r.URL.Path),
			zap.String("remote_addr", r.RemoteAddr),
			zap.String("id", ID.String()),
			zap.Any("product", productBody),
			zap.Error(err),
			zap.Time("timestamp", time.Now()),
		)
		HttpStatus.StatusInternalServerError(w, err)
		return
	}

	logger.GetLogger().Info("Successfully updated product",
		zap.String("method", r.Method),
		zap.String("url", r.URL.Path),
		zap.String("remote_addr", r.RemoteAddr),
		zap.String("id", ID.String()),
		zap.Time("timestamp", time.Now()),
	)
	HttpStatus.StatusNoContent(w)
}

func (h *Handler) Delete(w http.ResponseWriter, r *http.Request) {
	ID, err := uuid.Parse(chi.URLParam(r, "ID"))
	if err != nil {
		logger.GetLogger().Error("Invalid UUID format for deletion",
			zap.String("method", r.Method),
			zap.String("url", r.URL.Path),
			zap.String("remote_addr", r.RemoteAddr),
			zap.String("id", chi.URLParam(r, "ID")),
			zap.String("message", "ID is not uuid valid"),
			zap.Time("timestamp", time.Now()),
		)
		HttpStatus.StatusBadRequest(w, errors.New("ID is not uuid valid"))
		return
	}

	if err := h.Controller.Remove(ID); err != nil {
		logger.GetLogger().Error("Failed to delete product",
			zap.String("method", r.Method),
			zap.String("url", r.URL.Path),
			zap.String("remote_addr", r.RemoteAddr),
			zap.String("id", ID.String()),
			zap.Error(err),
			zap.Time("timestamp", time.Now()),
		)
		HttpStatus.StatusInternalServerError(w, err)
		return
	}

	logger.GetLogger().Info("Successfully deleted product",
		zap.String("method", r.Method),
		zap.String("url", r.URL.Path),
		zap.String("remote_addr", r.RemoteAddr),
		zap.String("id", ID.String()),
		zap.Time("timestamp", time.Now()),
	)
	HttpStatus.StatusNoContent(w)
}

func (h *Handler) Options(w http.ResponseWriter, r *http.Request) {
	logger.GetLogger().Info("OPTIONS request received",
		zap.String("method", r.Method),
		zap.String("url", r.URL.Path),
		zap.String("remote_addr", r.RemoteAddr),
		zap.Time("timestamp", time.Now()),
	)
	HttpStatus.StatusNoContent(w)
}

func (h *Handler) getBodyAndValidate(r *http.Request, ID uuid.UUID) (*EntityProduct.Product, error) {
	productBody := &EntityProduct.Product{}
	body, err := h.Rules.ConvertIoReaderToStruct(r.Body, productBody)
	if err != nil {
		logger.GetLogger().Error("Failed to convert request body to product structure",
			zap.String("method", r.Method),
			zap.String("url", r.URL.Path),
			zap.String("remote_addr", r.RemoteAddr),
			zap.Error(err),
			zap.Time("timestamp", time.Now()),
		)
		return &EntityProduct.Product{}, errors.New("body is required")
	}

	productParsed, err := EntityProduct.InterfaceToModel(body)
	if err != nil {
		logger.GetLogger().Error("Failed to convert request body to model",
			zap.String("method", r.Method),
			zap.String("url", r.URL.Path),
			zap.String("remote_addr", r.RemoteAddr),
			zap.Error(err),
			zap.Time("timestamp", time.Now()),
		)
		return &EntityProduct.Product{}, errors.New("error on convert body to model")
	}

	setDefaultValues(productParsed, ID)

	if err := h.Rules.Validate(productParsed); err != nil {
		logger.GetLogger().Error("Product validation failed",
			zap.String("method", r.Method),
			zap.String("url", r.URL.Path),
			zap.String("remote_addr", r.RemoteAddr),
			zap.Error(err),
			zap.Time("timestamp", time.Now()),
		)
		return productParsed, err
	}

	return productParsed, nil
}

func setDefaultValues(product *EntityProduct.Product, ID uuid.UUID) {
	product.UpdatedAt = time.Now()
	if ID == uuid.Nil {
		product.ID = uuid.New()
		product.CreatedAt = time.Now()
	} else {
		product.ID = ID
	}
}
