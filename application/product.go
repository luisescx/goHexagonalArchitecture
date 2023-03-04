package application

import (
	"errors"

	"github.com/asaskevich/govalidator"
	uuid "github.com/satori/go.uuid"
)

func init() {
	govalidator.SetFieldsRequiredByDefault(true)
}

type ProductInterface interface {
	IsValid() (bool, error)
	Enable() error
	Disable() error
	GetID() string
	GetName() string
	GetStatus() string
	GetPrice() float64
}

type ProductServiceInterface interface {
	Get(id string) (ProductInterface, error)
	Create(name string, price float64) (ProductInterface, error)
	Enable(product ProductInterface) (ProductInterface, error)
	Disable(product ProductInterface) (ProductInterface, error)
}

type ProductReader interface {
	Get(id string) (ProductInterface, error)
}

type ProductWriter interface {
	Save(product ProductInterface) (ProductInterface, error)
}

type ProductPersistenceInterface interface {
	ProductReader
	ProductWriter
}

const (
	DISABLED = "disabled"
	ENABLED = "enabled"
)

type Product struct {
	ID string `valid:"uuidv4"`
	Name string `valid:"required"`
	Price float64  `valid:"float,optional"`
	Status string  `valid:"required"`
}

func NewProduct() *Product {
	product := Product{
		ID: uuid.NewV4().String(),
		Status: DISABLED,
	}
	return &product
}

func (p *Product) IsValid() (bool, error) {
	if p.Status == "" {
		p.Status = DISABLED
	}

	if p.Status != DISABLED && p.Status != ENABLED {
		return false, errors.New("THE STATUS MUST BE ENABLED OR DISABLED")
	}

	if p.Price < 0 {
		return false, errors.New("THE PRICE MUST GREATER OR EQUAL TO ZERO")
	}

	value, err := govalidator.ValidateStruct(p)

	return value, err

}

func (p *Product) Enable() error {
	if p.Price > 0 {
		p.Status = ENABLED
		return nil
	}

	return errors.New("THE PRICE MUST BE GREATER THAN ZERO TO ENABLE THE PRODUCT")
}

func (p *Product) Disable() error {
	if p.Price == 0 {
		p.Status = DISABLED
		return nil
	}
	
	
	return errors.New("THE PRICE MUST BE ZERO IN ORDER TO HAVE THE PRODUCT DISABLED")
}

func (p *Product) GetID() string {
	return p.ID
}

func (p *Product) GetName() string {
	return p.Name
}

func (p *Product) GetStatus() string {
	return p.Status
}

func (p *Product) GetPrice() float64 {
	return p.Price
}