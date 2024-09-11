package product

import (
	"time"

	"github.com/aws/aws-sdk-go/service/dynamodb/expression"
	"github.com/datarohit/go-dynamodb-crud-api/internal/entities"
	"github.com/datarohit/go-dynamodb-crud-api/internal/entities/product"
	"github.com/datarohit/go-dynamodb-crud-api/internal/repository/adapter"
	"github.com/datarohit/go-dynamodb-crud-api/utils/logger"
	"github.com/google/uuid"
	"go.uber.org/zap"
)

type Controller struct {
	repository adapter.Interface
}

type Interface interface {
	ListOne(ID uuid.UUID) (product.Product, error)
	ListAll() ([]product.Product, error)
	Create(entity *product.Product) (uuid.UUID, error)
	Update(ID uuid.UUID, entity *product.Product) error
	Remove(ID uuid.UUID) error
}

func NewController(repository adapter.Interface) Interface {
	return &Controller{repository: repository}
}

func (c *Controller) ListOne(id uuid.UUID) (product.Product, error) {
	entity := product.Product{Base: entities.Base{ID: id}}
	response, err := c.repository.FindOne(entity.GetFilterId(), entity.TableName())
	if err != nil {
		logger.GetLogger().Error("Failed to find product by ID",
			zap.String("id", id.String()),
			zap.Error(err),
		)
		return entity, err
	}

	parsedEntity, err := product.ParseDynamoAttributeToStruct(response.Item)
	if err != nil {
		logger.GetLogger().Error("Failed to parse DynamoDB item",
			zap.String("id", id.String()),
			zap.Error(err),
		)
		return entity, err
	}

	return parsedEntity, nil
}

func (c *Controller) ListAll() ([]product.Product, error) {
	var entities []product.Product

	filter := expression.Name("name").NotEqual(expression.Value(""))
	condition, err := expression.NewBuilder().WithFilter(filter).Build()
	if err != nil {
		logger.GetLogger().Error("Failed to build DynamoDB query condition",
			zap.Error(err),
		)
		return entities, err
	}

	productInstance := &product.Product{}
	response, err := c.repository.FindAll(condition, productInstance.TableName())
	if err != nil {
		logger.GetLogger().Error("Failed to retrieve all products",
			zap.Error(err),
		)
		return entities, err
	}

	for _, item := range response.Items {
		entity, err := product.ParseDynamoAttributeToStruct(item)
		if err != nil {
			logger.GetLogger().Error("Failed to parse DynamoDB item",
				zap.Error(err),
			)
			return entities, err
		}
		entities = append(entities, entity)
	}

	return entities, nil
}

func (c *Controller) Create(entity *product.Product) (uuid.UUID, error) {
	entity.CreatedAt = time.Now()
	entity.ID = uuid.New()

	_, err := c.repository.CreateOrUpdate(entity.GetMap(), entity.TableName())
	if err != nil {
		logger.GetLogger().Error("Failed to create product",
			zap.String("name", entity.Name),
			zap.Error(err),
		)
		return uuid.Nil, err
	}

	logger.GetLogger().Info("Product created successfully",
		zap.String("id", entity.ID.String()),
		zap.String("name", entity.Name),
	)

	return entity.ID, nil
}

func (c *Controller) Update(id uuid.UUID, entity *product.Product) error {
	existingEntity, err := c.ListOne(id)
	if err != nil {
		logger.GetLogger().Error("Failed to find product for update",
			zap.String("id", id.String()),
			zap.Error(err),
		)
		return err
	}

	existingEntity.Name = entity.Name
	existingEntity.UpdatedAt = time.Now()

	_, err = c.repository.CreateOrUpdate(existingEntity.GetMap(), existingEntity.TableName())
	if err != nil {
		logger.GetLogger().Error("Failed to update product",
			zap.String("id", id.String()),
			zap.String("name", entity.Name),
			zap.Error(err),
		)
		return err
	}

	logger.GetLogger().Info("Product updated successfully",
		zap.String("id", id.String()),
		zap.String("name", entity.Name),
	)

	return nil
}

func (c *Controller) Remove(id uuid.UUID) error {
	entity, err := c.ListOne(id)
	if err != nil {
		logger.GetLogger().Error("Failed to find product for removal",
			zap.String("id", id.String()),
			zap.Error(err),
		)
		return err
	}

	_, err = c.repository.Delete(entity.GetFilterId(), entity.TableName())
	if err != nil {
		logger.GetLogger().Error("Failed to delete product",
			zap.String("id", id.String()),
			zap.Error(err),
		)
		return err
	}

	logger.GetLogger().Info("Product removed successfully",
		zap.String("id", id.String()),
	)

	return nil
}
