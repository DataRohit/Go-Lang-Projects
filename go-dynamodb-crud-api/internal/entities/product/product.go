package product

import (
	"encoding/json"
	"errors"
	"time"

	"github.com/aws/aws-sdk-go/service/dynamodb"
	"github.com/datarohit/go-dynamodb-crud-api/internal/entities"
	"github.com/google/uuid"
)

type Product struct {
	entities.Base
	Name string `json:"name"`
}

func InterfaceToModel(data interface{}) (product *Product, err error) {
	bytes, err := json.Marshal(data)
	if err != nil {
		return nil, err
	}

	product = &Product{}
	err = json.Unmarshal(bytes, product)
	if err != nil {
		return nil, err
	}

	return product, nil
}

func (p *Product) GetFilterId() map[string]interface{} {
	return map[string]interface{}{"_id": p.ID.String()}
}

func (p *Product) TableName() string {
	return "products"
}

func (p *Product) Bytes() ([]byte, error) {
	return json.Marshal(p)
}

func (p *Product) GetMap() map[string]interface{} {
	return map[string]interface{}{
		"_id":       p.ID.String(),
		"name":      p.Name,
		"createdAt": p.CreatedAt.Format(entities.GetTimeFormat()),
		"updatedAt": p.UpdatedAt.Format(entities.GetTimeFormat()),
	}
}

func ParseDynamoAttributeToStruct(response map[string]*dynamodb.AttributeValue) (p Product, err error) {
	if response == nil || len(response) == 0 {
		return p, errors.New("item not found")
	}

	if idAttr, ok := response["_id"]; ok {
		p.ID, err = uuid.Parse(*idAttr.S)
		if err != nil || p.ID == uuid.Nil {
			return p, errors.New("invalid or missing item ID")
		}
	}

	if nameAttr, ok := response["name"]; ok {
		p.Name = *nameAttr.S
	}

	if createdAtAttr, ok := response["createdAt"]; ok {
		p.CreatedAt, err = time.Parse(entities.GetTimeFormat(), *createdAtAttr.S)
		if err != nil {
			return p, errors.New("invalid creation timestamp")
		}
	}

	if updatedAtAttr, ok := response["updatedAt"]; ok {
		p.UpdatedAt, err = time.Parse(entities.GetTimeFormat(), *updatedAtAttr.S)
		if err != nil {
			return p, errors.New("invalid update timestamp")
		}
	}

	return p, nil
}
