package use_cases

import (
	"my-app/src/input_dtos"
	"my-app/src/models"
)

type CreateProductUseCase struct {
	InputDTO input_dtos.CreateProductInputDTO
}

func NewCreateProductUseCase(dto input_dtos.CreateProductInputDTO) CreateProductUseCase {
	return CreateProductUseCase{
		InputDTO: dto,
	}
}

func (uc CreateProductUseCase) Execute() models.Product {
	product := models.Product{
		Name: uc.InputDTO.Name,
		Description: uc.InputDTO.Description,
		Price: uc.InputDTO.Price,
		Quantity: uc.InputDTO.Quantity,
		Active: uc.InputDTO.Active,
	}

	models.DB.Create(&product)

	return product
}

type UpdateProductUseCase struct {
	InputDTO input_dtos.UpdateProductInputDTO
	Product models.Product
}

func NewUpdateProductUseCase(dto input_dtos.UpdateProductInputDTO, product models.Product) UpdateProductUseCase {
	return UpdateProductUseCase{
		InputDTO: dto,
		Product: product,
	}
}

func (uc UpdateProductUseCase) Execute() models.Product {
	models.DB.Model(&uc.Product).Updates(uc.InputDTO)
	return uc.Product
}

type DeleteProductUseCase struct {
	Product models.Product
}

func NewDeleteProductUseCase(product models.Product) DeleteProductUseCase {
	return DeleteProductUseCase{
		Product: product,
	}
}

func (uc DeleteProductUseCase) Execute() models.Product {
	models.DB.Delete(&uc.Product)
	return uc.Product
}
