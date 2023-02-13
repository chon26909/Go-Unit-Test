package repositories

type PromotionRepository interface {
	GetPromotion() (Promotion, error)
}

type Promotion struct {
	Id              int
	PurchaseMin     int
	DiscountPercent int
}
