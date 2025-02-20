package entities

type Currency string

const (
	USD = Currency("USD")
)

type BasePriceIndex string

const (
	BandwidthIndex = BasePriceIndex("bandwidth")
	DurationIndex  = BasePriceIndex("duration")
	IPCountIndex   = BasePriceIndex("ip_count")
	RegionIndex    = BasePriceIndex("region")
	ThreadsIndex   = BasePriceIndex("threads")
)

type Region string

const (
	RegionMixed = Region("mixed")
	RegionUSA   = Region("USA")
)

type TransactionStatus string

const (
	TransactionStatusPaid   = TransactionStatus("Paid")
	TransactionStatusUnpaid = TransactionStatus("Unpaid")
	TransactionStatusFailed = TransactionStatus("Failed")
)

type PaymentMethod string

const (
	Crypto = PaymentMethod("Crypto")
	Paddle = PaymentMethod("Paddle")
)

type CryptomusPaymentStatus string

const (
	CryptomusPaidStatus               = CryptomusPaymentStatus("paid")
	CryptomusPaidOverStatus           = CryptomusPaymentStatus("paid_over")
	CryptomusWrongAmountStatus        = CryptomusPaymentStatus("wrong_amount")
	CryptomusProcessStatus            = CryptomusPaymentStatus("process")
	CryptomusConfirmCheckStatus       = CryptomusPaymentStatus("confirm_check")
	CryptomusWrongAmountWaitingStatus = CryptomusPaymentStatus("wrong_amount_waiting")
	CryptomusCheckStatus              = CryptomusPaymentStatus("check")
	CryptomusFailStatus               = CryptomusPaymentStatus("fail")
	CryptomusCancelStatus             = CryptomusPaymentStatus("cancel")
	CryptomusSystemFailStatus         = CryptomusPaymentStatus("system_fail")
	CryptomusRefundProcessStatus      = CryptomusPaymentStatus("refund_process")
	CryptomusRefundFailStatus         = CryptomusPaymentStatus("refund_fail")
	CryptomusRefundPaidStatus         = CryptomusPaymentStatus("refund_paid")
	CryptomusLockedStatus             = CryptomusPaymentStatus("locked")
)

type CryptomusInvoiceType string

const (
	WalletInvoice  = CryptomusInvoiceType("wallet")
	PaymentInvoice = CryptomusInvoiceType("payment")
)

type ProxyServiceType string

const (
	ProxyStatic      = ProxyServiceType("static")
	ProxyBackconnect = ProxyServiceType("backconnect")
	ProxyProvider    = ProxyServiceType("provider")
	ProxySubnet      = ProxyServiceType("subnet")
	ProxyISPPool     = ProxyServiceType("isp_pool")
)

type ProductDisplayFeature struct {
	Supported   bool   `json:"supported"`
	Description string `json:"description"`
}

type IPVersion string

const (
	IPVersion4 = IPVersion("ipv4")
	IPVersion6 = IPVersion("ipv6")
)

type Protocol string

const (
	ProtocolHTTP    = Protocol("http")
	ProtocolHTTPS   = Protocol("https")
	ProtocolSOCKS5  = Protocol("socks5")
	ProtocolSOCKS5h = Protocol("socks5h")
)

// When Interval is nil, NumericSettingRange has range of [Min, InBetween, Max]
// When Interval is not nil, NumericSettingRange has range of [Min, Min + Interval, ..., Max - Interval, Max]
type NumericSettingRange struct {
	Min       int   `json:"min"`
	Max       int   `json:"max"`
	InBetween []int `json:"in_between,omitempty" gorm:"serializer:json"`
	Interval  *int  `json:"interval,omitempty"`
}

type PurchaseStopReason string

const (
	PurchaseStopReasonByAdmin      = PurchaseStopReason("by_admin")
	PurchaseStopReasonOutBandwidth = PurchaseStopReason("out_bandwidth")
	PurchaseStopReasonOutIPCount   = PurchaseStopReason("out_ip_count")
)

type Role string

const (
	RoleAdmin    = "admin"
	RoleUser     = "user"
	RoleCustomer = "customer"
)

type InvoiceStatus string

const (
	InvoiceStatusPaid   = InvoiceStatus("Paid")
	InvoiceStatusUnpaid = InvoiceStatus("Unpaid")
)

type PrizeRarity string

const (
	PrizeRarityCommon   = PrizeRarity("common")
	PrizeRarityUncommon = PrizeRarity("uncommon")
	PrizeRarityRare     = PrizeRarity("rare")
	PrizeRarityEpic     = PrizeRarity("epic")
)

type PrizeKind string

const (
	PrizeKindProduct       = PrizeKind("product")
	PrizeKindLoyaltyPoints = PrizeKind("loyalty_points")
	PrizeKindCredit        = PrizeKind("credit")
	PrizeKindDiscount      = PrizeKind("discount")
)

type SupportTicketStatus string

const (
	TicketOpened         = SupportTicketStatus("opened")
	TicketNeedFeedback = SupportTicketStatus("need_feedback")
	TicketClosed       = SupportTicketStatus("closed")
)

type SupportMessageSenderType string

const (
	SenderCustomer = SupportMessageSenderType("customer")
	SenderUser     = SupportMessageSenderType("user")
)
