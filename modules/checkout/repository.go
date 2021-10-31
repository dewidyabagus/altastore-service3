package checkout

import (
	"AltaStore/business/checkout"
	"errors"
	"time"

	"gorm.io/gorm"
)

type Repository struct {
	DB *gorm.DB
}

type Checkout struct {
	ID             string    `gorm:"id;type:uuid;primaryKey"`
	ShoppingCartId string    `gorm:"shopping_cart_id;type:uuid;unique"`
	Description    string    `gorm:"description;type:varchar(100)"`
	CreatedBy      string    `gorm:"created_by;type:varchar(50)"`
	CreatedAt      time.Time `gorm:"created_at;type:timestamp"`
	UpdatedAt      time.Time `gorm:"updated_at;type:timestamp"`
	DeletedAt      time.Time `gorm:"deleted_at;type:timestamp"`
}

func (c *Checkout) toBusinessCheckout() checkout.Checkout {
	return checkout.Checkout{
		ID:             c.ID,
		ShoppingCardId: c.ShoppingCartId,
		Description:    c.Description,
		CreatedBy:      c.CreatedBy,
		CreatedAt:      c.CreatedAt,
		UpdatedAt:      c.UpdatedAt,
	}
}

func toListCheckout(c *[]Checkout) *[]checkout.Checkout {
	var listCheckout []checkout.Checkout

	for _, checkout := range *c {
		listCheckout = append(listCheckout, checkout.toBusinessCheckout())
	}

	if listCheckout == nil {
		listCheckout = []checkout.Checkout{}
	}

	return &listCheckout
}

func insertCheckout(data *checkout.Checkout) *Checkout {
	return &Checkout{
		ID:             data.ID,
		ShoppingCartId: data.ShoppingCardId,
		Description:    data.Description,
		CreatedBy:      data.CreatedBy,
		CreatedAt:      data.CreatedAt,
		UpdatedAt:      data.UpdatedAt,
	}
}

func NewRepository(db *gorm.DB) *Repository {
	return &Repository{db}
}

func (r *Repository) NewCheckoutShoppingCart(checkout *checkout.Checkout) error {
	checkoutShopCart := insertCheckout(checkout)

	return r.DB.Create(&checkoutShopCart).Error
}

func (r *Repository) GetAllCheckout() (*[]checkout.Checkout, error) {
	var checkoutList []checkout.Checkout

	err := r.DB.Raw(
		"select c.*, cp.transaction_status payment_status " +
			"from checkouts c " +
			"inner join ( " +
			"select cp.check_out_id id, cp.transaction_status, " +
			"row_number() OVER (PARTITION BY  cp.check_out_id " +
			"order by cp.created_at desc) as rnum, " +
			"cp.created_at " +
			"from checkout_payments cp " +
			"order by cp.created_at desc " +
			")cp on cp.id  = CAST (c.id AS text) " +
			"where  cp.rnum = 1").Scan(&checkoutList).Error
	if err != nil {
		return nil, err
	}

	return &checkoutList, nil
}

func (r *Repository) GetCheckoutById(id string) (*checkout.Checkout, error) {
	checkout := new(checkout.Checkout)

	err := r.DB.Raw(
		"select c.*, cp.transaction_status payment_status "+
			"from checkouts c "+
			"inner join ( "+
			"select cp.check_out_id id, cp.transaction_status, "+
			"row_number() OVER (PARTITION BY  cp.check_out_id "+
			"order by cp.created_at desc) as rnum, "+
			"cp.created_at "+
			"from checkout_payments cp "+
			"order by cp.created_at desc "+
			")cp on cp.id  = CAST (c.id AS text) "+
			"where  cp.rnum = 1 and  CAST (c.id AS text) = ? ", id).Scan(&checkout).Error
	if err != nil {
		return nil, err
	}

	result := checkout

	return result, nil
}

func (r *Repository) GetCheckoutByShoppingCartId(cartId string) (bool, error) {
	var checkout = new(Checkout)

	err := r.DB.First(checkout, " shopping_cart_id = ? ", cartId).Error
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return false, nil
		}
		return false, err
	}

	return true, nil
}
