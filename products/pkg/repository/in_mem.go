package repository

import "agarwalconsulting.io/rvstore/entities"

type inmemProductsRepository struct {
	products []entities.Product
}

func (pr *inmemProductsRepository) FetchAll() ([]entities.Product, error) {
	return pr.products, nil
}

// NewProductsRepository creates an instance of an in-memory products repository
func NewProductsRepository() ProductsRepository {
	products := []entities.Product{
		{
			ID:          "sewer-hose",
			Name:        "Stinky Slinky Sewer Hose",
			Price:       19.99,
			Description: "Oh what fun times you'll have with this! Be sure to stock up on gloves and disinfectant soap before checkout!",
			Image:       "https://campaddict.com/wp-content/uploads/RV-sewer-hose-in-use.jpg",
		},
		{
			ID:          "electric-jacks",
			Name:        "Six-Point Auto-Leveling Jacks",
			Price:       3499.99,
			Description: "Six-point electric auto-leveling jacks for your travel trailer or fifth wheel. No more yelling at your spouse for getting the number of wood blocks wrong! Just push a button and walk away!",
			Image:       "https://cdn3.volusion.com/dxylq.nruds/v/vspfiles/photos/8735-2.jpg?1536058347",
		},
		{
			ID:          "heated-hose",
			Name:        "Heated Fresh Water Hose",
			Price:       89.99,
			Description: "Heated hose for those freezing nights",
			Image:       "https://i5.walmartimages.com/asr/9b926f3b-cdd5-4bcc-b09b-ef44d461c023_2.e4c509076a08564a97b13b61ea0c45da.jpeg",
		},
		{
			ID:          "50-amp-extension-cord",
			Name:        "25' 50 amp Extension Cord",
			Price:       89.99,
			Description: "Quit parking your RV based on where the post is!",
			Image:       "https://images-na.ssl-images-amazon.com/images/I/71t4SHO3EML._SX425_.jpg",
		},
		{
			ID:          "5th-wheel-tripod",
			Name:        "5th Wheel Tripod Stabilizer",
			Price:       127.49,
			Description: "This rig won't be a rockin' when you have this tripod stabilizer on your 5th wheel hitch.",
			Image:       "https://www.etrailer.com/static/images/faq/review-ultrafab-5th-wheel-king-pin-tripod-stabilizer-uf19-950001_644.jpg",
		},
	}

	return &inmemProductsRepository{products}
}
