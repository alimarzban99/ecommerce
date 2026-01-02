package enums

type OrderStatus string

const (
	OrderPending   OrderStatus = "pending"   // در انتظار پرداخت / تایید
	OrderPaid      OrderStatus = "paid"      // پرداخت شده
	OrderShipped   OrderStatus = "shipped"   // ارسال شده
	OrderCancelled OrderStatus = "cancelled" // لغو شده
	OrderRefunded  OrderStatus = "refunded"  // بازگشت وجه
)
