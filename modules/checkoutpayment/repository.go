package checkoutpayment

import (
	"AltaStore/business/checkoutpayment"
	"time"

	"gorm.io/gorm"
)

type CheckoutPayment struct {
	CheckOutID        string                        `gorm:"checkout_id;type:uuid"`
	MerchantId        string                        `gorm:"merchant_id"`
	StatusCode        string                        `gorm:"status_code"`
	TransactionStatus checkoutpayment.PaymentStatus `gorm:"transaction_status"`
	FraudStatus       checkoutpayment.FraudStatus   `gorm:"fraud_status"`
	CreatedAt         time.Time                     `gorm:"created_at"`
	CreatedBy         string                        `gorm:"created_by;type:varchar(50)"`
	UpdatedAt         time.Time                     `gorm:"updated_at"`
	UpdatedBy         string                        `gorm:"updated_by;type:varchar(50)"`
	DeletedAt         time.Time                     `gorm:"deleted_at"`
	DeletedBy         string                        `gorm:"deleted_by;type:varchar(50)"`
}

func (p *CheckoutPayment) ToPayment() *checkoutpayment.CheckoutPayment {
	return &checkoutpayment.CheckoutPayment{
		CheckOutID:        p.CheckOutID,
		MerchantId:        p.MerchantId,
		StatusCode:        p.StatusCode,
		TransactionStatus: p.TransactionStatus,
		FraudStatus:       p.FraudStatus,
		CreatedAt:         p.CreatedAt,
		CreatedBy:         p.CreatedBy,
		UpdatedAt:         p.UpdatedAt,
		UpdatedBy:         p.UpdatedBy,
		DeletedAt:         p.DeletedAt,
		DeletedBy:         p.DeletedBy,
	}
}

func newPayment(p *checkoutpayment.CheckoutPayment) *CheckoutPayment {
	return &CheckoutPayment{
		CheckOutID:        p.CheckOutID,
		MerchantId:        p.MerchantId,
		StatusCode:        p.StatusCode,
		TransactionStatus: p.TransactionStatus,
		FraudStatus:       p.FraudStatus,
		CreatedAt:         p.CreatedAt,
		CreatedBy:         p.CreatedBy,
		UpdatedAt:         p.UpdatedAt,
		UpdatedBy:         p.UpdatedBy,
		DeletedAt:         p.DeletedAt,
		DeletedBy:         p.DeletedBy,
	}
}

type Repository struct {
	DB *gorm.DB
}

func NewRepository(db *gorm.DB) *Repository {
	return &Repository{db}
}
func (r *Repository) InsertPayment(p *checkoutpayment.CheckoutPayment) (*checkoutpayment.CheckoutPayment, error) {
	data := newPayment(p)
	if err := r.DB.Create(data).Error; err != nil {
		return nil, err
	}

	return data.ToPayment(), nil
}

// func (r *Repository) UpdatePayment(p *checkoutpayment.CheckoutPayment) error {
// 	data := newPayment(p)
// 	err := r.DB.Model(&data).Updates(checkoutpayment.CheckoutPayment{
// 		MerchantId:        data.MerchantId,
// 		TransactionStatus: data.TransactionStatus,
// 		FraudStatus:       data.FraudStatus,
// 		UpdatedAt:         data.UpdatedAt,
// 		UpdatedBy:         data.UpdatedBy,
// 	}).Error
// 	if err != nil {
// 		return err
// 	}
// 	return nil
// }
