package adapter

import "time"

type PaymentType string
type ApmType string
type CardProvider string
type PaymentProvider string
type PosApmPaymentProvider string
type PaymentStatus string
type TokenizedCardType string
type PaymentSource string
type PaymentGroup string
type PaymentPhase string
type PaymentMethod string
type CardType string
type CardAssociation string
type CardExpiryStatus string
type Currency string
type LoyaltyType string
type MultiPaymentStatus string
type PaymentRefundStatus string
type RefundStatus string
type RefundType string
type ApprovalStatus string
type Status string
type OnboardingStatus string
type MemberType string
type SettlementType string
type SettlementEarningsDestination string
type RefundDestinationType string
type TransactionStatus string
type TransactionPayoutStatus string
type WalletTransactionRefundTransactionType string
type FraudAction string
type FraudCheckStatus string
type FraudValueType string
type AdditionalAction string
type ApmAdditionalAction string
type ReportFileType string
type WalletTransactionType string
type WebhookEventType string
type WebhookStatus string
type PosStatus string
type PosIntegrator string
type PosUserType string
type PosOperationType string
type FileStatus string
type AccountOwner string
type PayoutAccountType string
type RecordType string
type BankAccountTrackingSource string
type BnplCartItemType string
type PaymentAuthenticationType string
type CardBrand string
type ClientType string
type MasterpassValidationType string

const (
	ApiKeyHeaderName        = "x-api-key"
	RandomHeaderName        = "x-rnd-key"
	AuthVersionHeaderName   = "x-auth-version"
	ClientVersionHeaderName = "x-client-version"
	SignatureHeaderName     = "x-signature"
	AuthVersion             = "1"
	ClientVersion           = "craftgate-go-client:1.0.22"
)

// payment type declaration
const (
	PaymentType_CARD_PAYMENT             PaymentType = "CARD_PAYMENT"
	PaymentType_WALLET_PAYMENT           PaymentType = "WALLET_PAYMENT"
	PaymentType_CARD_AND_WALLET_PAYMENT  PaymentType = "CARD_AND_WALLET_PAYMENT"
	PaymentType_DEPOSIT_PAYMENT          PaymentType = "DEPOSIT_PAYMENT"
	PaymentType_HEPSIPAY_DEPOSIT_PAYMENT PaymentType = "HEPSIPAY_DEPOSIT_PAYMENT"
	PaymentType_BANK_TRANSFER            PaymentType = "BANK_TRANSFER"
	PaymentType_APM                      PaymentType = "APM"
)

// apm type declaration
const (
	ApmType_PAPARA                 ApmType = "PAPARA"
	ApmType_PAYONEER               ApmType = "PAYONEER"
	ApmType_SODEXO                 ApmType = "SODEXO"
	ApmType_SETCARD                ApmType = "SETCARD"
	ApmType_METROPOL               ApmType = "METROPOL"
	ApmType_EDENRED                ApmType = "EDENRED"
	ApmType_EDENRED_GIFT           ApmType = "EDENRED_GIFT"
	ApmType_PAYPAL                 ApmType = "PAYPAL"
	ApmType_KLARNA                 ApmType = "KLARNA"
	ApmType_AFTERPAY               ApmType = "AFTERPAY"
	ApmType_KASPI                  ApmType = "KASPI"
	ApmType_INSTANT_TRANSFER       ApmType = "INSTANT_TRANSFER"
	ApmType_TOMPAY                 ApmType = "TOMPAY"
	ApmType_MASLAK                 ApmType = "MASLAK"
	ApmType_ALFABANK               ApmType = "ALFABANK"
	ApmType_TOM_FINANCE            ApmType = "TOM_FINANCE"
	ApmType_STRIPE                 ApmType = "STRIPE"
	ApmType_PAYCELL                ApmType = "PAYCELL"
	ApmType_HASO                   ApmType = "HASO"
	ApmType_MULTINET               ApmType = "MULTINET"
	ApmType_MULTINET_GIFT          ApmType = "MULTINET_GIFT"
	ApmType_MULTINET_NEO_GIFT      ApmType = "MULTINET_NEO_GIFT"
	ApmType_ALIPAY                 ApmType = "ALIPAY"
	ApmType_APPLEPAY               ApmType = "APPLEPAY"
	ApmType_GOOGLEPAY              ApmType = "GOOGLEPAY"
	ApmType_HEPSIPAY_WALLET        ApmType = "HEPSIPAY_WALLET"
	ApmType_HEPSIPAY_SHOPPING_LOAN ApmType = "HEPSIPAY_SHOPPING_LOAN"
	ApmType_ZIP                    ApmType = "ZIP"
	ApmType_CHIPPIN                ApmType = "CHIPPIN"
	ApmType_ISPAY                  ApmType = "ISPAY"
	ApmType_VODAFONE_DCB           ApmType = "VODAFONE_DCB"
	ApmType_PAYMOB                 ApmType = "PAYMOB"
	ApmType_BIZUM                  ApmType = "BIZUM"
	ApmType_PAYLANDS_MB_WAY        ApmType = "PAYLANDS_MB_WAY"
	ApmType_PAYCELL_DCB            ApmType = "PAYCELL_DCB"
	ApmType_IWALLET                ApmType = "IWALLET"
	ApmType_FUND_TRANSFER          ApmType = "FUND_TRANSFER"
	ApmType_CASH_ON_DELIVERY       ApmType = "CASH_ON_DELIVERY"
)

// card provider declaration
const (
	CardProvider_IYZICO         CardProvider = "IYZICO"
	CardProvider_IPARA          CardProvider = "IPARA"
	CardProvider_BIRLESIK_ODEME CardProvider = "BIRLESIK_ODEME"
	CardProvider_MEX            CardProvider = "MEX"
)

// payment provider declaration
const (
	PaymentProvider_BANK                        PaymentProvider = "BANK"
	PaymentProvider_CG_WALLET                   PaymentProvider = "CG_WALLET"
	PaymentProvider_MASTERPASS                  PaymentProvider = "MASTERPASS"
	PaymentProvider_GARANTI_PAY                 PaymentProvider = "GARANTI_PAY"
	PaymentProvider_YKB_WORLD_PAY               PaymentProvider = "YKB_WORLD_PAY"
	PaymentProvider_PAPARA                      PaymentProvider = "PAPARA"
	PaymentProvider_PAYONEER                    PaymentProvider = "PAYONEER"
	PaymentProvider_SODEXO                      PaymentProvider = "SODEXO"
	PaymentProvider_METROPOL                    PaymentProvider = "METROPOL"
	PaymentProvider_EDENRED                     PaymentProvider = "EDENRED"
	PaymentProvider_ALIPAY                      PaymentProvider = "ALIPAY"
	PaymentProvider_PAYPAL                      PaymentProvider = "PAYPAL"
	PaymentProvider_KLARNA                      PaymentProvider = "KLARNA"
	PaymentProvider_AFTERPAY                    PaymentProvider = "AFTERPAY"
	PaymentProvider_APPLEPAY                    PaymentProvider = "APPLEPAY"
	PaymentProvider_GOOGLEPAY                   PaymentProvider = "GOOGLEPAY"
	PaymentProvider_HEPSIPAY                    PaymentProvider = "HEPSIPAY"
	PaymentProvider_STRIPE                      PaymentProvider = "STRIPE"
	PaymentProvider_KASPI                       PaymentProvider = "KASPI"
	PaymentProvider_INSTANT_TRANSFER            PaymentProvider = "INSTANT_TRANSFER"
	PaymentProvider_MASLAK                      PaymentProvider = "MASLAK"
	PaymentProvider_TOMPAY                      PaymentProvider = "TOMPAY"
	PaymentProvider_TOM_FINANCE                 PaymentProvider = "TOM_FINANCE"
	PaymentProvider_ALFABANK                    PaymentProvider = "ALFABANK"
	PaymentProvider_PAYCELL                     PaymentProvider = "PAYCELL"
	PaymentProvider_HASO                        PaymentProvider = "HASO"
	PaymentProvider_MULTINET                    PaymentProvider = "MULTINET"
	PaymentProvider_YKB_WORLD_PAY_SHOPPING_LOAN PaymentProvider = "YKB_WORLD_PAY_SHOPPING_LOAN"
	PaymentProvider_ZIP                         PaymentProvider = "ZIP"
	PaymentProvider_CHIPPIN                     PaymentProvider = "CHIPPIN"
	PaymentProvider_ISPAY                       PaymentProvider = "ISPAY"
	PaymentProvider_VODAFONE                    PaymentProvider = "VODAFONE"
	PaymentProvider_PAYMOB                      PaymentProvider = "PAYMOB"
	PaymentProvider_BIZUM                       PaymentProvider = "BIZUM"
	PaymentProvider_PAYLANDS_MB_WAY             PaymentProvider = "PAYLANDS_MB_WAY"
	PaymentProvider_PAYCELL_DCB                 PaymentProvider = "PAYCELL_DCB"
	PaymentProvider_IWALLET                     PaymentProvider = "IWALLET"
	PaymentProvider_OFFLINE                     PaymentProvider = "OFFLINE"
)

// pos apm payment provider declaration
const (
	PosApmPaymentProvider_YKB_WORLD_PAY               PosApmPaymentProvider = "YKB_WORLD_PAY"
	PosApmPaymentProvider_YKB_WORLD_PAY_SHOPPING_LOAN PosApmPaymentProvider = "YKB_WORLD_PAY_SHOPPING_LOAN"
	PosApmPaymentProvider_GOOGLEPAY                   PosApmPaymentProvider = "GOOGLEPAY"
	PosApmPaymentProvider_GARANTI_PAY                 PosApmPaymentProvider = "GARANTI_PAY"
)

// payment status declaration
const (
	PaymentStatus_FAILURE          PaymentStatus = "FAILURE"
	PaymentStatus_SUCCESS          PaymentStatus = "SUCCESS"
	PaymentStatus_INIT_THREEDS     PaymentStatus = "INIT_THREEDS"
	PaymentStatus_CALLBACK_THREEDS PaymentStatus = "CALLBACK_THREEDS"
	PaymentStatus_WAITING          PaymentStatus = "WAITING"
)

// payment source declaration
const (
	PaymentSource_API           PaymentSource = "API"
	PaymentSource_CHECKOUT_FORM PaymentSource = "CHECKOUT_FORM"
	PaymentSource_PAY_BY_LINK   PaymentSource = "PAY_BY_LINK"
	PaymentSource_BKM_EXPRESS   PaymentSource = "BKM_EXPRESS"
	PaymentSource_HEPSIPAY      PaymentSource = "HEPSIPAY"
	PaymentSource_JUZDAN        PaymentSource = "JUZDAN"
	PaymentSource_MONO          PaymentSource = "MONO"
)

// currency declaration
const (
	Currency_TRY Currency = "TRY"
	Currency_USD Currency = "USD"
	Currency_EUR Currency = "EUR"
	Currency_GBP Currency = "GBP"
	Currency_CNY Currency = "CNY"
	Currency_ARS Currency = "ARS"
	Currency_BRL Currency = "BRL"
	Currency_AED Currency = "AED"
	Currency_IQD Currency = "IQD"
	Currency_AZN Currency = "AZN"
	Currency_KZT Currency = "KZT"
	Currency_KWD Currency = "KWD"
	Currency_SAR Currency = "SAR"
	Currency_BHD Currency = "BHD"
	Currency_RUB Currency = "RUB"
	Currency_JPY Currency = "JPY"
	Currency_EGP Currency = "EGP"
	Currency_MXN Currency = "MXN"
)

// payment group declaration
const (
	PaymentGroup_PRODUCT                 PaymentGroup = "PRODUCT"
	PaymentGroup_LISTING_OR_SUBSCRIPTION PaymentGroup = "LISTING_OR_SUBSCRIPTION"
)

// payment phase declaration
const (
	PaymentPhase_AUTH      PaymentPhase = "AUTH"
	PaymentPhase_PRE_AUTH  PaymentPhase = "PRE_AUTH"
	PaymentPhase_POST_AUTH PaymentPhase = "POST_AUTH"
)

// payment method declaration
const (
	PaymentMethod_CARD              PaymentMethod = "CARD"
	PaymentMethod_MASTERPASS        PaymentMethod = "MASTERPASS"
	PaymentMethod_PAPARA            PaymentMethod = "PAPARA"
	PaymentMethod_PAYONEER          PaymentMethod = "PAYONEER"
	PaymentMethod_SODEXO            PaymentMethod = "SODEXO"
	PaymentMethod_EDENRED           PaymentMethod = "EDENRED"
	PaymentMethod_EDENRED_GIFT      PaymentMethod = "EDENRED_GIFT"
	PaymentMethod_ALIPAY            PaymentMethod = "ALIPAY"
	PaymentMethod_PAYPAL            PaymentMethod = "PAYPAL"
	PaymentMethod_KLARNA            PaymentMethod = "KLARNA"
	PaymentMethod_AFTERPAY          PaymentMethod = "AFTERPAY"
	PaymentMethod_INSTANT_TRANSFER  PaymentMethod = "INSTANT_TRANSFER"
	PaymentMethod_STRIPE            PaymentMethod = "STRIPE"
	PaymentMethod_MULTINET          PaymentMethod = "MULTINET"
	PaymentMethod_MULTINET_GIFT     PaymentMethod = "MULTINET_GIFT"
	PaymentMethod_MULTINET_NEO_GIFT PaymentMethod = "MULTINET_NEO_GIFT"
	PaymentMethod_PAYLANDS_MB_WAY   PaymentMethod = "PAYLANDS_MB_WAY"
	PaymentMethod_PAYCELL_DCB       PaymentMethod = "PAYCELL_DCB"
	PaymentMethod_IWALLET           PaymentMethod = "IWALLET"
)

// card type declaration
const (
	CardType_CREDIT_CARD  CardType = "CREDIT_CARD"
	CardType_DEBIT_CARD   CardType = "DEBIT_CARD"
	CardType_PREPAID_CARD CardType = "PREPAID_CARD"
)

// card association declaration
const (
	CardAssociation_VISA        CardAssociation = "VISA"
	CardAssociation_MASTER_CARD CardAssociation = "MASTER_CARD"
	CardAssociation_AMEX        CardAssociation = "AMEX"
	CardAssociation_TROY        CardAssociation = "TROY"
	CardAssociation_JCB         CardAssociation = "JCB"
	CardAssociation_UNION_PAY   CardAssociation = "UNION_PAY"
	CardAssociation_MAESTRO     CardAssociation = "MAESTRO"
	CardAssociation_DISCOVER    CardAssociation = "DISCOVER"
	CardAssociation_DINERS_CLUB CardAssociation = "DINERS_CLUB"
)

// card expiry status declaration
const (
	CardExpiryStatus_EXPIRED                CardExpiryStatus = "EXPIRED"
	CardExpiryStatus_WILL_EXPIRE_NEXT_MONTH CardExpiryStatus = "WILL_EXPIRE_NEXT_MONTH"
	CardExpiryStatus_NOT_EXPIRED            CardExpiryStatus = "NOT_EXPIRED"
)

// loyalty type declaration
const (
	LoyaltyType_REWARD_MONEY           LoyaltyType = "REWARD_MONEY"
	LoyaltyType_ADDITIONAL_INSTALLMENT LoyaltyType = "ADDITIONAL_INSTALLMENT"
	LoyaltyType_POSTPONING_INSTALLMENT LoyaltyType = "POSTPONING_INSTALLMENT"
	LoyaltyType_EXTRA_POINTS           LoyaltyType = "EXTRA_POINTS"
	LoyaltyType_GAINING_MINUTES        LoyaltyType = "GAINING_MINUTES"
	LoyaltyType_POSTPONING_STATEMENT   LoyaltyType = "POSTPONING_STATEMENT"
)

// multi payment status declaration
const (
	MultiPaymentStatus_CREATED   MultiPaymentStatus = "CREATED"
	MultiPaymentStatus_COMPLETED MultiPaymentStatus = "COMPLETED"
)

// payment refund status declaration
const (
	PaymentRefundStatus_NO_REFUND        PaymentRefundStatus = "NO_REFUND"
	PaymentRefundStatus_NOT_REFUNDED     PaymentRefundStatus = "NOT_REFUNDED"
	PaymentRefundStatus_PARTIAL_REFUNDED PaymentRefundStatus = "PARTIAL_REFUNDED"
	PaymentRefundStatus_FULLY_REFUNDED   PaymentRefundStatus = "FULLY_REFUNDED"
)

// refund status declaration
const (
	RefundStatus_SUCCESS RefundStatus = "SUCCESS"
	RefundStatus_FAILURE RefundStatus = "FAILURE"
)

// onboarding status declaration
const (
	OnboardingStatus_APPLICATION_REJECTED OnboardingStatus = "APPLICATION_REJECTED"
	OnboardingStatus_REGISTERED           OnboardingStatus = "REGISTERED"
	OnboardingStatus_REGISTER_CONFIRMED   OnboardingStatus = "REGISTER_CONFIRMED"
	OnboardingStatus_APPLIED              OnboardingStatus = "APPLIED"
	OnboardingStatus_INTEGRATION          OnboardingStatus = "INTEGRATION"
	OnboardingStatus_LIVE                 OnboardingStatus = "LIVE"
)

// refund type declaration
const (
	RefundType_CANCEL RefundType = "CANCEL"
	RefundType_REFUND RefundType = "REFUND"
)

// approval status declaration
const (
	ApprovalStatus_SUCCESS ApprovalStatus = "SUCCESS"
	ApprovalStatus_FAILURE ApprovalStatus = "FAILURE"
)

// status declaration
const (
	Status_ACTIVE  Status = "ACTIVE"
	Status_PASSIVE Status = "PASSIVE"
)

// member type declaration
const (
	MemberType_PERSONAL                       MemberType = "PERSONAL"
	MemberType_PRIVATE_COMPANY                MemberType = "PRIVATE_COMPANY"
	MemberType_LIMITED_OR_JOINT_STOCK_COMPANY MemberType = "LIMITED_OR_JOINT_STOCK_COMPANY"
)

// settlementEarningsDestination type declaration
const (
	SettlementEarningsDestination_IBAN         SettlementEarningsDestination = "IBAN"
	SettlementEarningsDestination_WALLET       SettlementEarningsDestination = "WALLET"
	SettlementEarningsDestination_CROSS_BORDER SettlementEarningsDestination = "CROSS_BORDER"
	SettlementEarningsDestination_NONE         SettlementEarningsDestination = "NONE"
)

// refundDestinationType type declaration
const (
	RefundDestinationType_PROVIDER RefundDestinationType = "PROVIDER"
	RefundDestinationType_WALLET   RefundDestinationType = "WALLET"
)

// transaction status declaration
const (
	TransactionStatus_WAITING_FOR_APPROVAL TransactionStatus = "WAITING_FOR_APPROVAL"
	TransactionStatus_APPROVED             TransactionStatus = "APPROVED"
	TransactionStatus_PAYOUT_STARTED       TransactionStatus = "PAYOUT_STARTED"
)

// transaction payout status declaration
const (
	TransactionPayoutStatus_CANCELLED          TransactionPayoutStatus = "CANCELLED"
	TransactionPayoutStatus_NO_PAYOUT          TransactionPayoutStatus = "NO_PAYOUT"
	TransactionPayoutStatus_WAITING_FOR_PAYOUT TransactionPayoutStatus = "WAITING_FOR_PAYOUT"
	TransactionPayoutStatus_PAYOUT_STARTED     TransactionPayoutStatus = "PAYOUT_STARTED"
	TransactionPayoutStatus_PAYOUT_COMPLETED   TransactionPayoutStatus = "PAYOUT_COMPLETED"
)

// settlement type declaration
const (
	SettlementType_SETTLEMENT         SettlementType = "SETTLEMENT"
	SettlementType_BOUNCED_SETTLEMENT SettlementType = "BOUNCED_SETTLEMENT"
	SettlementType_WITHDRAW           SettlementType = "WITHDRAW"
)

// wallet transaction refund type declaration
const (
	WalletTransactionRefundCardTransactionType_PAYMENT    WalletTransactionRefundTransactionType = "PAYMENT"
	WalletTransactionRefundCardTransactionType_PAYMENT_TX WalletTransactionRefundTransactionType = "PAYMENT_TX"
	WalletTransactionRefundCardTransactionType_WALLET_TX  WalletTransactionRefundTransactionType = "WALLET_TX"
)

// fraud action type declaration
const (
	FraudAction_BLOCK  FraudAction = "BLOCK"
	FraudAction_REVIEW FraudAction = "REVIEW"
)

// fraud check status type declaration
const (
	FraudCheckStatus_WAITING   FraudCheckStatus = "WAITING"
	FraudCheckStatus_NOT_FRAUD FraudCheckStatus = "NOT_FRAUD"
	FraudCheckStatus_FRAUD     FraudCheckStatus = "FRAUD"
)

// fraud value type type declaration
const (
	FraudValueType_CARD         FraudValueType = "CARD"
	FraudValueType_IP           FraudValueType = "IP"
	FraudValueType_PHONE_NUMBER FraudValueType = "PHONE_NUMBER"
	FraudValueType_EMAIL        FraudValueType = "EMAIL"
	FraudValueType_OTHER        FraudValueType = "OTHER"
)

// apm additional action type declaration
const (
	ApmAdditionalAction_REDIRECT_TO_URL   ApmAdditionalAction = "REDIRECT_TO_URL"
	ApmAdditionalAction_OTP_REQUIRED      ApmAdditionalAction = "OTP_REQUIRED"
	ApmAdditionalAction_SHOW_HTML_CONTENT ApmAdditionalAction = "SHOW_HTML_CONTENT"
	ApmAdditionalAction_WAIT_FOR_WEBHOOK  ApmAdditionalAction = "WAIT_FOR_WEBHOOK"
	ApmAdditionalAction_APPROVAL_REQUIRED ApmAdditionalAction = "APPROVAL_REQUIRED"
	ApmAdditionalAction_NONE              ApmAdditionalAction = "NONE"
)

// report file type declaration
const (
	ReportFileType_CSV  ReportFileType = "CSV"
	ReportFileType_XLSX ReportFileType = "XLSX"
)

// wallet transaction type declaration
const (
	WalletTransactionType_PAYMENT_REDEEM                 WalletTransactionType = "PAYMENT_REDEEM"
	WalletTransactionType_REFUND_DEPOSIT                 WalletTransactionType = "REFUND_DEPOSIT"
	WalletTransactionType_REFUND_TX_DEPOSIT              WalletTransactionType = "REFUND_TX_DEPOSIT"
	WalletTransactionType_WITHDRAW                       WalletTransactionType = "WITHDRAW"
	WalletTransactionType_CANCEL_REFUND_WALLET_TO_CARD   WalletTransactionType = "CANCEL_REFUND_WALLET_TO_CARD"
	WalletTransactionType_REFUND_WALLET_TX_TO_CARD       WalletTransactionType = "REFUND_WALLET_TX_TO_CARD"
	WalletTransactionType_REFUND_WALLET_TX_FUND_TRANSFER WalletTransactionType = "REFUND_WALLET_TX_FUND_TRANSFER"
	WalletTransactionType_CANCEL_REFUND_TO_WALLET        WalletTransactionType = "CANCEL_REFUND_TO_WALLET"
	WalletTransactionType_REFUND_TX_TO_WALLET            WalletTransactionType = "REFUND_TX_TO_WALLET"
	WalletTransactionType_MANUAL_REFUND_TX_TO_WALLET     WalletTransactionType = "MANUAL_REFUND_TX_TO_WALLET"
	WalletTransactionType_SETTLEMENT_EARNINGS            WalletTransactionType = "SETTLEMENT_EARNINGS"
	WalletTransactionType_DEPOSIT_FROM_CARD              WalletTransactionType = "DEPOSIT_FROM_CARD"
	WalletTransactionType_DEPOSIT_FROM_APM               WalletTransactionType = "DEPOSIT_FROM_APM"
	WalletTransactionType_DEPOSIT_FROM_FUND_TRANSFER     WalletTransactionType = "DEPOSIT_FROM_FUND_TRANSFER"
	WalletTransactionType_REMITTANCE                     WalletTransactionType = "REMITTANCE"
	WalletTransactionType_LOYALTY                        WalletTransactionType = "LOYALTY"
	WalletTransactionType_WITHDRAW_CANCEL                WalletTransactionType = "WITHDRAW_CANCEL"
	WalletTransactionType_MERCHANT_BALANCE_RESET         WalletTransactionType = "MERCHANT_BALANCE_RESET"
)

const (
	WebhookEventType_API_AUTH            WebhookEventType = "API_AUTH"
	WebhookEventType_API_VERIFY_AND_AUTH WebhookEventType = "API_VERIFY_AND_AUTH"
	WebhookEventType_CHECKOUTFORM_AUTH   WebhookEventType = "CHECKOUTFORM_AUTH"
	WebhookEventType_THREEDS_VERIFY      WebhookEventType = "THREEDS_VERIFY"
	WebhookEventType_REFUND              WebhookEventType = "REFUND"
	WebhookEventType_REFUND_TX           WebhookEventType = "REFUND_TX"
	WebhookEventType_PAYOUT_COMPLETED    WebhookEventType = "PAYOUT_COMPLETED"
	WebhookEventType_AUTOPILOT           WebhookEventType = "AUTOPILOT"
	WebhookEventType_WALLET_CREATED      WebhookEventType = "WALLET_CREATED"
	WebhookEventType_WALLET_TX_CREATED   WebhookEventType = "WALLET_TX_CREATED"
	WebhookEventType_BNPL_NOTIFICATION   WebhookEventType = "BNPL_NOTIFICATION"
)

const (
	WebhookStatus_SUCCESS WebhookStatus = "SUCCESS"
	WebhookStatus_FAILURE WebhookStatus = "FAILURE"
)

const (
	PosStatus_DELETED     PosStatus = "DELETED"
	PosStatus_PASSIVE     PosStatus = "PASSIVE"
	PosStatus_ACTIVE      PosStatus = "ACTIVE"
	PosStatus_REFUND_ONLY PosStatus = "REFUND_ONLY"
	PosStatus_AUTOPILOT   PosStatus = "AUTOPILOT"
)

const (
	PosIntegrator_YKB               PosIntegrator = "YKB"
	PosIntegrator_GARANTI           PosIntegrator = "GARANTI"
	PosIntegrator_ISBANK            PosIntegrator = "ISBANK"
	PosIntegrator_AKBANK            PosIntegrator = "AKBANK"
	PosIntegrator_ZIRAATBANK        PosIntegrator = "ZIRAATBANK"
	PosIntegrator_ZIRAATBANK_INNOVA PosIntegrator = "ZIRAATBANK_INNOVA"
	PosIntegrator_ZIRAATKATILIM     PosIntegrator = "ZIRAATKATILIM"
	PosIntegrator_KUVEYTTURK        PosIntegrator = "KUVEYTTURK"
	PosIntegrator_HALKBANK          PosIntegrator = "HALKBANK"
	PosIntegrator_DENIZBANK         PosIntegrator = "DENIZBANK"
	PosIntegrator_VAKIFBANK         PosIntegrator = "VAKIFBANK"
	PosIntegrator_VAKIFKATILIM      PosIntegrator = "VAKIFKATILIM"
	PosIntegrator_FINANSBANK        PosIntegrator = "FINANSBANK"
	PosIntegrator_FIBABANK          PosIntegrator = "FIBABANK"
	PosIntegrator_FIBABANK_ASSECO   PosIntegrator = "FIBABANK_ASSECO"
	PosIntegrator_ANADOLUBANK       PosIntegrator = "ANADOLUBANK"
	PosIntegrator_PARAM_POS         PosIntegrator = "PARAM_POS"
	PosIntegrator_IYZICO            PosIntegrator = "IYZICO"
	PosIntegrator_SIPAY             PosIntegrator = "SIPAY"
	PosIntegrator_PAYNET            PosIntegrator = "PAYNET"
	PosIntegrator_PAYTR             PosIntegrator = "PAYTR"
	PosIntegrator_BIRLESIK_ODEME    PosIntegrator = "BIRLESIK_ODEME"
	PosIntegrator_MOKA              PosIntegrator = "MOKA"
	PosIntegrator_STRIPE            PosIntegrator = "STRIPE"
	PosIntegrator_TEB               PosIntegrator = "TEB"
	PosIntegrator_IPARA             PosIntegrator = "IPARA"
	PosIntegrator_OZAN              PosIntegrator = "OZAN"
	PosIntegrator_BRAINTREE         PosIntegrator = "BRAINTREE"
	PosIntegrator_NKOLAY            PosIntegrator = "NKOLAY"
	PosIntegrator_PAYTABS           PosIntegrator = "PAYTABS"
	PosIntegrator_PAYBULL           PosIntegrator = "PAYBULL"
	PosIntegrator_ELEKSE            PosIntegrator = "ELEKSE"
	PosIntegrator_ALGORITMA         PosIntegrator = "ALGORITMA"
	PosIntegrator_PAYCELL           PosIntegrator = "PAYCELL"
	PosIntegrator_TAMI              PosIntegrator = "TAMI"
	PosIntegrator_QNB_PAY           PosIntegrator = "QNB_PAY"
	PosIntegrator_AKBANK_VPOS       PosIntegrator = "AKBANK_VPOS"
	PosIntegrator_TAP               PosIntegrator = "TAP"
	PosIntegrator_RUBIK             PosIntegrator = "RUBIK"
)

const (
	PosUserType_API PosUserType = "API"
)

const (
	PosOperationType_STANDARD PosOperationType = "STANDARD"
	PosOperationType_PROVAUT  PosOperationType = "PROVAUT"
	PosOperationType_PROVRFN  PosOperationType = "PROVRFN"
	PosOperationType_PAYMENT  PosOperationType = "PAYMENT"
	PosOperationType_REFUND   PosOperationType = "REFUND"
	PosOperationType_INQUIRY  PosOperationType = "INQUIRY"
)

const (
	FileStatus_CREATED  FileStatus = "CREATED"
	FileStatus_UPLOADED FileStatus = "UPLOADED"
	FileStatus_APPROVED FileStatus = "APPROVED"
)

const (
	AdditionalAction_CONTINUE_IN_CLIENT AdditionalAction = "CONTINUE_IN_CLIENT"
	AdditionalAction_SHOW_HTML_CONTENT  AdditionalAction = "SHOW_HTML_CONTENT"
	AdditionalAction_REDIRECT_TO_URL    AdditionalAction = "REDIRECT_TO_URL"
	AdditionalAction_NONE               AdditionalAction = "NONE"
)

const (
	AccountOwner_MERCHANT            AccountOwner = "MERCHANT"
	AccountOwner_SUB_MERCHANT_MEMBER AccountOwner = "SUB_MERCHANT_MEMBER"
)

const (
	PayoutAccountType_WISE PayoutAccountType = "WISE"
)

// BnplCartItemType type declaration
const (
	BnplCartItemType_MOBILE_PHONE_OVER_5000_TRY                BnplCartItemType = "MOBILE_PHONE_OVER_5000_TRY"
	BnplCartItemType_MOBILE_PHONE_BELOW_5000_TRY               BnplCartItemType = "MOBILE_PHONE_BELOW_5000_TRY"
	BnplCartItemType_MOBILE_PHONE_PRICE_ABOVE_REGULATION_LIMIT BnplCartItemType = "MOBILE_PHONE_PRICE_ABOVE_REGULATION_LIMIT"
	BnplCartItemType_MOBILE_PHONE_PRICE_BELOW_REGULATION_LIMIT BnplCartItemType = "MOBILE_PHONE_PRICE_BELOW_REGULATION_LIMIT"
	BnplCartItemType_TABLET                                    BnplCartItemType = "TABLET"
	BnplCartItemType_COMPUTER                                  BnplCartItemType = "COMPUTER"
	BnplCartItemType_CONSTRUCTION_MARKET                       BnplCartItemType = "CONSTRUCTION_MARKET"
	BnplCartItemType_GOLD                                      BnplCartItemType = "GOLD"
	BnplCartItemType_DIGITAL_PRODUCTS                          BnplCartItemType = "DIGITAL_PRODUCTS"
	BnplCartItemType_SUPERMARKET                               BnplCartItemType = "SUPERMARKET"
	BnplCartItemType_WHITE_GOODS                               BnplCartItemType = "WHITE_GOODS"
	BnplCartItemType_WEARABLE_TECHNOLOGY                       BnplCartItemType = "WEARABLE_TECHNOLOGY"
	BnplCartItemType_SMALL_HOME_APPLIANCES                     BnplCartItemType = "SMALL_HOME_APPLIANCES"
	BnplCartItemType_TV                                        BnplCartItemType = "TV"
	BnplCartItemType_GAMES_CONSOLES                            BnplCartItemType = "GAMES_CONSOLES"
	BnplCartItemType_AIR_CONDITIONER_AND_HEATER                BnplCartItemType = "AIR_CONDITIONER_AND_HEATER"
	BnplCartItemType_ELECTRONICS                               BnplCartItemType = "ELECTRONICS"
	BnplCartItemType_ACCESSORIES                               BnplCartItemType = "ACCESSORIES"
	BnplCartItemType_MOM_AND_BABY_AND_KIDS                     BnplCartItemType = "MOM_AND_BABY_AND_KIDS"
	BnplCartItemType_SHOES                                     BnplCartItemType = "SHOES"
	BnplCartItemType_CLOTHING                                  BnplCartItemType = "CLOTHING"
	BnplCartItemType_COSMETICS_AND_PERSONAL_CARE               BnplCartItemType = "COSMETICS_AND_PERSONAL_CARE"
	BnplCartItemType_FURNITURE                                 BnplCartItemType = "FURNITURE"
	BnplCartItemType_HOME_LIVING                               BnplCartItemType = "HOME_LIVING"
	BnplCartItemType_AUTOMOBILE_MOTORCYCLE                     BnplCartItemType = "AUTOMOBILE_MOTORCYCLE"
	BnplCartItemType_MOBILE_PHONE_REFURBISHED                  BnplCartItemType = "MOBILE_PHONE_REFURBISHED"
	BnplCartItemType_OTHER                                     BnplCartItemType = "OTHER"
)

// RecordType declaration
const (
	RecordType_SEND    RecordType = "SEND"
	RecordType_RECEIVE RecordType = "RECEIVE"
)

// BankAccountTrackingSource declaration
const (
	BankAccountTrackingSource_YKB     BankAccountTrackingSource = "YKB"
	BankAccountTrackingSource_GARANTI BankAccountTrackingSource = "GARANTI"
)

const (
	PaymentAuthenticationType_THREE_DS     PaymentAuthenticationType = "THREE_DS"
	PaymentAuthenticationType_NON_THREE_DS PaymentAuthenticationType = "NON_THREE_DS"
	PaymentAuthenticationType_BKM_EXPRESS  PaymentAuthenticationType = "BKM_EXPRESS"
)

const (
	CardBrand_BONUS          CardBrand = "Bonus"
	CardBrand_AXESS          CardBrand = "Axess"
	CardBrand_MAXIMUM        CardBrand = "Maximum"
	CardBrand_WORLD          CardBrand = "World"
	CardBrand_PARAF          CardBrand = "Paraf"
	CardBrand_CARD_FINANS    CardBrand = "CardFinans"
	CardBrand_BANKKART_COMBO CardBrand = "Bankkart Combo"
	CardBrand_ADVANTAGE      CardBrand = "Advantage"
	CardBrand_SAGLAM_KART    CardBrand = "Sağlam Kart"
)

const (
	ClientType_W ClientType = "W"
	ClientType_M ClientType = "M"
)

// tokenized card type declaration
const (
	TokenizedCardType_APPLE_PAY TokenizedCardType = "APPLE_PAY"
)

const (
	MasterpassValidationType_NONE     MasterpassValidationType = "NONE"
	MasterpassValidationType_OTP      MasterpassValidationType = "OTP"
	MasterpassValidationType_THREE_DS MasterpassValidationType = "THREE_DS"
)

// requests
type CreatePaymentRequest struct {
	Price            float64                `json:"price,omitempty"`
	PaidPrice        float64                `json:"paidPrice,omitempty"`
	WalletPrice      float64                `json:"walletPrice,omitempty"`
	PosAlias         string                 `json:"posAlias,omitempty"`
	Installment      int                    `json:"installment,omitempty"`
	Currency         Currency               `json:"currency,omitempty"`
	PaymentGroup     PaymentGroup           `json:"paymentGroup,omitempty"`
	ConversationId   string                 `json:"conversationId,omitempty"`
	ExternalId       string                 `json:"externalId,omitempty"`
	ClientIp         string                 `json:"clientIp,omitempty"`
	PaymentPhase     PaymentPhase           `json:"paymentPhase,omitempty"`
	PaymentChannel   string                 `json:"paymentChannel,omitempty"`
	BuyerMemberId    int64                  `json:"buyerMemberId,omitempty"`
	BankOrderId      string                 `json:"bankOrderId,omitempty"`
	Card             *Card                  `json:"card,omitempty"`
	FraudParams      *FraudCheckParameters  `json:"fraudParams,omitempty"`
	Items            []PaymentItem          `json:"items"`
	AdditionalParams map[string]interface{} `json:"additionalParams,omitempty"`
	Retry            bool                   `json:"retry"`
}

type CreateApmPaymentRequest struct {
	ApmType        ApmType       `json:"apmType,omitempty"`
	Price          float64       `json:"price,omitempty"`
	PaidPrice      float64       `json:"paidPrice,omitempty"`
	Currency       Currency      `json:"currency,omitempty"`
	PaymentGroup   PaymentGroup  `json:"paymentGroup,omitempty"`
	PaymentChannel string        `json:"paymentChannel,omitempty"`
	ConversationId string        `json:"conversationId,omitempty"`
	ExternalId     string        `json:"externalId,omitempty"`
	BuyerMemberId  int64         `json:"buyerMemberId,omitempty"`
	ApmOrderId     string        `json:"apmOrderId,omitempty"`
	ClientIp       string        `json:"clientIp,omitempty"`
	Items          []PaymentItem `json:"items"`
}

type Init3DSPaymentRequest struct {
	Price            float64                `json:"price,omitempty"`
	PaidPrice        float64                `json:"paidPrice,omitempty"`
	WalletPrice      float64                `json:"walletPrice,omitempty"`
	PosAlias         string                 `json:"posAlias,omitempty"`
	Installment      int                    `json:"installment,omitempty"`
	Currency         Currency               `json:"currency,omitempty"`
	PaymentGroup     PaymentGroup           `json:"paymentGroup,omitempty"`
	ConversationId   string                 `json:"conversationId,omitempty"`
	ExternalId       string                 `json:"externalId,omitempty"`
	ClientIp         string                 `json:"clientIp,omitempty"`
	PaymentPhase     PaymentPhase           `json:"paymentPhase,omitempty"`
	PaymentChannel   string                 `json:"paymentChannel,omitempty"`
	BuyerMemberId    int64                  `json:"buyerMemberId,omitempty"`
	BankOrderId      string                 `json:"bankOrderId,omitempty"`
	Card             *Card                  `json:"card,omitempty"`
	CallbackUrl      string                 `json:"callbackUrl,omitempty"`
	Items            []PaymentItem          `json:"items"`
	AdditionalParams map[string]interface{} `json:"additionalParams"`
	Retry            bool                   `json:"retry"`
}

type Complete3DSPaymentRequest struct {
	PaymentId int64 `json:"paymentId"`
}

type InitCheckoutPaymentRequest struct {
	Price                       float64                        `json:"price,omitempty"`
	PaidPrice                   float64                        `json:"paidPrice,omitempty"`
	Currency                    Currency                       `json:"currency,omitempty"`
	PaymentGroup                PaymentGroup                   `json:"paymentGroup,omitempty"`
	ConversationId              string                         `json:"conversationId,omitempty"`
	ExternalId                  string                         `json:"externalId,omitempty"`
	OrderId                     string                         `json:"orderId,omitempty"`
	CallbackUrl                 string                         `json:"callbackUrl,omitempty"`
	ClientIp                    string                         `json:"clientIp,omitempty"`
	PaymentPhase                PaymentPhase                   `json:"paymentPhase,omitempty"`
	PaymentChannel              string                         `json:"paymentChannel,omitempty"`
	EnabledPaymentMethods       []PaymentMethod                `json:"enabledPaymentMethods,omitempty"`
	MasterpassGsmNumber         string                         `json:"masterpassGsmNumber,omitempty"`
	MasterpassUserId            string                         `json:"masterpassUserId,omitempty"`
	CardUserKey                 string                         `json:"cardUserKey,omitempty"`
	BuyerMemberId               int64                          `json:"buyerMemberId,omitempty"`
	EnabledInstallments         []int                          `json:"enabledInstallments,omitempty"`
	AlwaysStoreCardAfterPayment bool                           `json:"alwaysStoreCardAfterPayment,omitempty"`
	AllowOnlyStoredCards        bool                           `json:"allowOnlyStoredCards,omitempty"`
	AllowOnlyCreditCard         bool                           `json:"allowOnlyCreditCard,omitempty"`
	ForceThreeDS                bool                           `json:"forceThreeDS,omitempty"`
	ForceAuthForNonCreditCards  bool                           `json:"forceAuthForNonCreditCards,omitempty"`
	DepositPayment              bool                           `json:"depositPayment,omitempty"`
	Ttl                         int64                          `json:"ttl,omitempty"`
	CustomInstallments          []CustomInstallment            `json:"customInstallments,omitempty"`
	Items                       []PaymentItem                  `json:"items"`
	FraudParams                 *FraudCheckParameters          `json:"fraudParams,omitempty"`
	AdditionalParams            map[string]interface{}         `json:"additionalParams,omitempty"`
	CardBrandInstallments       map[string][]CustomInstallment `json:"cardBrandInstallments,omitempty"`
}

type InitApmPaymentRequest struct {
	ApmType          ApmType           `json:"apmType,omitempty"`
	MerchantApmId    int64             `json:"merchantApmId,omitempty"`
	Price            float64           `json:"price,omitempty"`
	PaidPrice        float64           `json:"paidPrice,omitempty"`
	Currency         Currency          `json:"currency,omitempty"`
	PaymentGroup     PaymentGroup      `json:"paymentGroup,omitempty"`
	PaymentChannel   string            `json:"paymentChannel,omitempty"`
	ConversationId   string            `json:"conversationId,omitempty"`
	ExternalId       string            `json:"externalId,omitempty"`
	CallbackUrl      string            `json:"callbackUrl,omitempty"`
	BuyerMemberId    int64             `json:"buyerMemberId,omitempty"`
	ApmOrderId       string            `json:"apmOrderId,omitempty"`
	ApmUserIdentity  string            `json:"apmUserIdentity,omitempty"`
	ClientIp         string            `json:"clientIp,omitempty"`
	AdditionalParams map[string]string `json:"additionalParams,omitempty"`
	Items            []PaymentItem     `json:"items"`
}

type CompleteApmPaymentRequest struct {
	PaymentId        int64             `json:"paymentId,omitempty"`
	AdditionalParams map[string]string `json:"additionalParams,omitempty"`
}

type InitPosApmPaymentRequest struct {
	Price             float64               `json:"price,omitempty"`
	PaidPrice         float64               `json:"paidPrice,omitempty"`
	PosAlias          string                `json:"posAlias,omitempty"`
	Currency          Currency              `json:"currency,omitempty"`
	ConversationId    string                `json:"conversationId,omitempty"`
	ExternalId        string                `json:"externalId,omitempty"`
	CallbackUrl       string                `json:"callbackUrl,omitempty"`
	PaymentPhase      PaymentPhase          `json:"paymentPhase,omitempty"`
	PaymentGroup      PaymentGroup          `json:"paymentGroup,omitempty"`
	PaymentChannel    string                `json:"paymentChannel,omitempty"`
	BuyerMemberId     int64                 `json:"buyerMemberId,omitempty"`
	BankOrderId       string                `json:"bankOrderId,omitempty"`
	ClientIp          string                `json:"clientIp,omitempty"`
	Items             []PaymentItem         `json:"items"`
	AdditionalParams  map[string]string     `json:"additionalParams"`
	Installments      []PosApmInstallment   `json:"installments"`
	PaymentProvider   PosApmPaymentProvider `json:"paymentProvider,omitempty"`
	FraudParams       *FraudCheckParameters `json:"fraudParams,omitempty"`
	CheckoutFormToken string                `json:"checkoutFormToken,omitempty"`
}

type CompletePosApmPaymentRequest struct {
	PaymentId        int64                  `json:"paymentId"`
	AdditionalParams map[string]interface{} `json:"additionalParams"`
}

type PostAuthPaymentRequest struct {
	PaidPrice float64 `json:"paidPrice"`
}

type DepositPaymentRequest struct {
	BuyerMemberId  int64    `json:"buyerMemberId,omitempty"`
	Price          float64  `json:"price,omitempty"`
	Currency       Currency `json:"currency,omitempty"`
	ConversationId string   `json:"conversationId,omitempty"`
	CallbackUrl    string   `json:"callbackUrl,omitempty"`
	PosAlias       string   `json:"posAlias,omitempty"`
	ClientIp       string   `json:"clientIp,omitempty"`
	Card           Card     `json:"card"`
}

type CreateFundTransferDepositPaymentRequest struct {
	Price          float64 `json:"price,omitempty"`
	BuyerMemberId  int64   `json:"buyerMemberId,omitempty"`
	ConversationId string  `json:"conversationId,omitempty"`
	ClientIp       string  `json:"clientIp,omitempty"`
}

type InitApmDepositPaymentRequest struct {
	ApmType          ApmType           `json:"apmType,omitempty"`
	MerchantApmId    int64             `json:"merchantApmId,omitempty"`
	Price            float64           `json:"price,omitempty"`
	Currency         Currency          `json:"currency,omitempty"`
	BuyerMemberId    int64             `json:"buyerMemberId,omitempty"`
	PaymentChannel   string            `json:"paymentChannel,omitempty"`
	ConversationId   string            `json:"conversationId,omitempty"`
	ExternalId       string            `json:"externalId,omitempty"`
	CallbackUrl      string            `json:"callbackUrl,omitempty"`
	ApmOrderId       string            `json:"apmOrderId,omitempty"`
	ApmUserIdentity  string            `json:"apmUserIdentity,omitempty"`
	ClientIp         string            `json:"clientIp,omitempty"`
	AdditionalParams map[string]string `json:"additionalParams,omitempty"`
}

type RetrieveLoyaltiesRequest struct {
	CardNumber  string `json:"cardNumber,omitempty"`
	ExpireYear  string `json:"expireYear,omitempty"`
	ExpireMonth string `json:"expireMonth,omitempty"`
	Cvc         string `json:"cvc,omitempty"`
	CardUserKey string `json:"cardUserKey,omitempty"`
	CardToken   string `json:"cardToken,omitempty"`
}

type RetrieveProviderCardRequest struct {
	ProviderCardToken  string `json:"providerCardToken,omitempty"`
	ExternalId         string `json:"externalId,omitempty"`
	ProviderCardUserId string `json:"providerCardUserId,omitempty"`
	CardProvider       string `json:"cardProvider,omitempty"`
}

type MasterpassRetrieveLoyaltiesRequest struct {
	Msisdn    string `json:"msisdn,omitempty"`
	BinNumber string `json:"binNumber,omitempty"`
	CardName  string `json:"cardName,omitempty"`
}

type InitGarantiPayPaymentRequest struct {
	Price               float64                 `json:"price,omitempty"`
	PaidPrice           float64                 `json:"paidPrice,omitempty"`
	Currency            Currency                `json:"currency,omitempty"`
	PosAlias            string                  `json:"posAlias,omitempty"`
	PaymentGroup        PaymentGroup            `json:"paymentGroup,omitempty"`
	ConversationId      string                  `json:"conversationId,omitempty"`
	ExternalId          string                  `json:"externalId,omitempty"`
	CallbackUrl         string                  `json:"callbackUrl,omitempty"`
	ClientIp            string                  `json:"clientIp,omitempty"`
	PaymentChannel      string                  `json:"paymentChannel,omitempty"`
	BuyerMemberId       int64                   `json:"buyerMemberId,omitempty"`
	BankOrderId         string                  `json:"bankOrderId,omitempty"`
	Items               []PaymentItem           `json:"items"`
	Installments        []GarantiPayInstallment `json:"installments,omitempty"`
	EnabledInstallments []int                   `json:"enabledInstallments"`
}

type RefundPaymentTransactionRequest struct {
	PaymentTransactionId  int64                 `json:"paymentTransactionId"`
	ConversationId        string                `json:"conversationId"`
	RefundPrice           float64               `json:"refundPrice"`
	RefundDestinationType RefundDestinationType `json:"refundDestinationType"`
	ChargeFromMe          bool                  `json:"chargeFromMe"`
}

type UpdatePaymentTransactionRequest struct {
	SubMerchantMemberId    int64   `json:"subMerchantMemberId,omitempty"`
	SubMerchantMemberPrice float64 `json:"subMerchantMemberPrice,omitempty"`
}

type UpdateStoredCardRequest struct {
	CardUserKey string `json:"cardUserKey,omitempty"`
	CardToken   string `json:"cardToken,omitempty"`
	ExpireYear  string `json:"expireYear,omitempty"`
	ExpireMonth string `json:"expireMonth,omitempty"`
}

type CloneStoredCardRequest struct {
	SourceCardUserKey string `json:"sourceCardUserKey"`
	SourceCardToken   string `json:"sourceCardToken"`
	TargetCardUserKey string `json:"targetCardUserKey,omitempty"`
	TargetMerchantId  int64  `json:"targetMerchantId"`
}

type DeleteStoredCardRequest struct {
	CardUserKey string `json:"cardUserKey,omitempty"`
	CardToken   string `json:"cardToken,omitempty"`
}

type SearchStoredCardsRequest struct {
	CardAlias       string          `schema:"cardAlias,omitempty"`
	CardBrand       string          `schema:"cardBrand,omitempty"`
	CardType        CardType        `schema:"cardType,omitempty"`
	CardUserKey     string          `schema:"cardUserKey,omitempty"`
	CardToken       string          `schema:"cardToken,omitempty"`
	CardBankName    string          `schema:"cardBankName,omitempty"`
	CardAssociation CardAssociation `schema:"cardAssociation,omitempty"`
	MinCreatedDate  time.Time       `schema:"minCreatedDate,omitempty"`
	MaxCreatedDate  time.Time       `schema:"maxCreatedDate,omitempty"`
	Page            int             `schema:"page,omitempty"`
	Size            int             `schema:"size,omitempty"`
}

type PaymentTransactionsApprovalRequest struct {
	PaymentTransactionIds []int64 `json:"paymentTransactionIds,omitempty"`
	IsTransactional       bool    `json:"isTransactional,omitempty"`
}

type RefundPaymentRequest struct {
	PaymentId             int64                 `json:"paymentId,omitempty"`
	ConversationId        string                `json:"conversationId,omitempty"`
	RefundDestinationType RefundDestinationType `json:"refundDestinationType,omitempty"`
	ChargeFromMe          bool                  `json:"chargeFromMe,omitempty"`
}

type StoreCardRequest struct {
	CardHolderName string `json:"cardHolderName,omitempty"`
	CardNumber     string `json:"cardNumber,omitempty"`
	ExpireYear     string `json:"expireYear,omitempty"`
	ExpireMonth    string `json:"expireMonth,omitempty"`
	CardAlias      string `json:"cardAlias,omitempty"`
	CardUserKey    string `json:"cardUserKey,omitempty"`
}

type ApplePayMerchantSessionCreateRequest struct {
	MerchantIdentifier string `json:"merchantIdentifier,omitempty"`
	DisplayName        string `json:"displayName,omitempty"`
	Initiative         string `json:"initiative,omitempty"`
	InitiativeContext  string `json:"initiativeContext,omitempty"`
	ValidationUrl      string `json:"validationUrl,omitempty"`
}

type CheckMasterpassUserRequest struct {
	MasterpassGsmNumber string `json:"masterpassGsmNumber"`
}

type CreatePayoutAccountRequest struct {
	AccountType         PayoutAccountType `json:"type,omitempty"`
	ExternalAccountId   string            `json:"externalAccountId,omitempty"`
	Currency            Currency          `json:"currency,omitempty"`
	AccountOwner        AccountOwner      `json:"accountOwner,omitempty"`
	SubMerchantMemberId int64             `json:"subMerchantMemberId,omitempty"`
}

type UpdatePayoutAccountRequest struct {
	AccountType       PayoutAccountType `json:"type,omitempty"`
	ExternalAccountId string            `json:"externalAccountId,omitempty"`
}

type SearchPayoutAccountRequest struct {
	Currency            Currency     `json:"currency,omitempty"`
	AccountOwner        AccountOwner `json:"accountOwner,omitempty"`
	SubMerchantMemberId int64        `json:"subMerchantMemberId,omitempty"`
	Page                int          `schema:"page,omitempty"`
	Size                int          `schema:"size,omitempty"`
}

type MasterpassPaymentTokenGenerateRequest struct {
	Msisdn                       string                   `json:"msisdn,omitempty"`
	UserId                       string                   `json:"userId,omitempty"`
	BinNumber                    string                   `json:"binNumber,omitempty"`
	ForceThreeDS                 bool                     `json:"forceThreeDS,omitempty"`
	CreatePayment                MasterpassCreatePayment  `json:"createPayment,omitempty"`
	Loyalty                      *Loyalty                 `json:"loyalty,omitempty"`
	MasterpassIntegrationVersion int                      `json:"masterpassIntegrationVersion,omitempty"`
	ValidationType               MasterpassValidationType `json:"validationType,omitempty"`
}

type MasterpassPaymentCompleteRequest struct {
	ReferenceId string `json:"referenceId,omitempty"`
	Token       string `json:"token,omitempty"`
}

type MasterpassPaymentThreeDSInitRequest struct {
	ReferenceId string `json:"referenceId,omitempty"`
	CallbackUrl string `json:"callbackUrl,omitempty"`
}

type MasterpassPaymentThreeDSCompleteRequest struct {
	PaymentId int64 `json:"paymentId,omitempty"`
}

type InitBnplPaymentRequest struct {
	ApmType          ApmType               `json:"apmType"`
	MerchantApmId    int64                 `json:"merchantApmId,omitempty"`
	Price            float64               `json:"price"`
	PaidPrice        float64               `json:"paidPrice"`
	CommissionRate   float64               `json:"commissionRate,omitempty"`
	Currency         Currency              `json:"currency"`
	PaymentType      PaymentType           `json:"paymentType"`
	PaymentGroup     PaymentGroup          `json:"paymentGroup"`
	PaymentSource    PaymentSource         `json:"paymentSource,omitempty"`
	PaymentChannel   string                `json:"paymentChannel,omitempty"`
	ConversationId   string                `json:"conversationId,omitempty"`
	ExternalId       string                `json:"externalId,omitempty"`
	CallbackUrl      string                `json:"callbackUrl"`
	BuyerMemberId    int64                 `json:"buyerMemberId,omitempty"`
	ApmOrderId       string                `json:"apmOrderId,omitempty"`
	ClientIp         string                `json:"clientIp,omitempty"`
	ApmUserIdentity  string                `json:"apmUserIdentity,omitempty"`
	AdditionalParams map[string]string     `json:"additionalParams"`
	Items            []PaymentItem         `json:"items"`
	BankCode         string                `json:"bankCode,omitempty"`
	CartItems        []BnplPaymentCartItem `json:"cartItems"`
}

type BnplPaymentCartItem struct {
	Id        string           `json:"id"`
	Name      string           `json:"name"`
	BrandName string           `json:"brandName"`
	Type      BnplCartItemType `json:"type"`
	UnitPrice float64          `json:"unitPrice"`
	Quantity  int64            `json:"quantity"`
}

type BnplPaymentOfferRequest struct {
	ApmType          ApmType               `json:"apmType"`
	MerchantApmId    int64                 `json:"merchantApmId,omitempty"`
	Price            float64               `json:"price"`
	Currency         Currency              `json:"currency"`
	ApmOrderId       string                `json:"apmOrderId,omitempty"`
	AdditionalParams map[string]string     `json:"additionalParams"`
	Items            []BnplPaymentCartItem `json:"items"`
}

// responses
type PaymentResponse struct {
	Id                           *int64                       `json:"id"`
	CreatedDate                  *TimeResponse                `json:"createdDate"`
	Price                        *float64                     `json:"price"`
	PaidPrice                    *float64                     `json:"paidPrice"`
	WalletPrice                  *float64                     `json:"walletPrice"`
	Currency                     *string                      `json:"currency"`
	BuyerMemberId                *int64                       `json:"buyerMemberId"`
	FraudId                      *int64                       `json:"fraudId"`
	Installment                  *int                         `json:"installment"`
	ConversationId               *string                      `json:"conversationId"`
	ExternalId                   *string                      `json:"externalId"`
	PaymentType                  *PaymentType                 `json:"paymentType"`
	PaymentProvider              *PaymentProvider             `json:"paymentProvider"`
	PaymentSource                *PaymentSource               `json:"paymentSource"`
	PaymentGroup                 *PaymentGroup                `json:"paymentGroup"`
	PaymentStatus                *PaymentStatus               `json:"paymentStatus"`
	PaymentPhase                 *PaymentPhase                `json:"paymentPhase"`
	FraudAction                  *FraudAction                 `json:"fraudAction"`
	PaymentChannel               *string                      `json:"paymentChannel"`
	IsThreeDS                    *bool                        `json:"isThreeDS"`
	MerchantCommissionRate       *float64                     `json:"merchantCommissionRate"`
	MerchantCommissionRateAmount *float64                     `json:"merchantCommissionRateAmount"`
	BankCommissionRate           *float64                     `json:"bankCommissionRate"`
	BankCommissionRateAmount     *float64                     `json:"bankCommissionRateAmount"`
	PaidWithStoredCard           *bool                        `json:"paidWithStoredCard"`
	BinNumber                    *string                      `json:"binNumber"`
	LastFourDigits               *string                      `json:"lastFourDigits"`
	AuthCode                     *string                      `json:"authCode"`
	HostReference                *string                      `json:"hostReference"`
	TransId                      *string                      `json:"transId"`
	MdStatus                     *int                         `json:"mdStatus"`
	OrderId                      *string                      `json:"orderId"`
	CardHolderName               *string                      `json:"cardHolderName"`
	BankCardHolderName           *string                      `json:"bankCardHolderName"`
	CardType                     *string                      `json:"cardType"`
	CardAssociation              *string                      `json:"cardAssociation"`
	CardBrand                    *string                      `json:"cardBrand"`
	RequestedPosAlias            *string                      `json:"requestedPosAlias"`
	Pos                          *MerchantPos                 `json:"pos"`
	Loyalty                      *Loyalty                     `json:"loyalty"`
	PaymentError                 *PaymentError                `json:"paymentError"`
	CardUserKey                  *string                      `json:"cardUserKey"`
	CardToken                    *string                      `json:"cardToken"`
	PaymentTransactions          []PaymentTransactionResponse `json:"paymentTransactions"`
	AdditionalData               map[string]any               `json:"additionalData"`
}

type PaymentTransactionResponse struct {
	ID                            *int64   `json:"id"`
	ExternalID                    *string  `json:"externalId"`
	Name                          *string  `json:"name"`
	TransactionStatus             *string  `json:"transactionStatus"`
	BlockageResolvedDate          *string  `json:"blockageResolvedDate"`
	Price                         *float64 `json:"price"`
	PaidPrice                     *float64 `json:"paidPrice"`
	WalletPrice                   *float64 `json:"walletPrice"`
	MerchantCommissionRate        *float64 `json:"merchantCommissionRate"`
	MerchantCommissionRateAmount  *float64 `json:"merchantCommissionRateAmount"`
	MerchantPayoutAmount          *float64 `json:"merchantPayoutAmount"`
	SubMerchantMemberID           *int64   `json:"subMerchantMemberId"`
	SubMerchantMemberPrice        *float64 `json:"subMerchantMemberPrice"`
	SubMerchantMemberPayoutRate   *float64 `json:"subMerchantMemberPayoutRate"`
	SubMerchantMemberPayoutAmount *float64 `json:"subMerchantMemberPayoutAmount"`
	Payout                        *Payout  `json:"payout"`
}

type Init3DSPaymentResponse struct {
	HtmlContent      *string           `json:"htmlContent"`
	PaymentId        *int64            `json:"paymentId"`
	PaymentStatus    *PaymentStatus    `json:"paymentStatus"`
	AdditionalAction *AdditionalAction `json:"additionalAction"`
}

type InitCheckoutPaymentResponse struct {
	Token           *string       `json:"token"`
	PageUrl         *string       `json:"pageUrl"`
	TokenExpireDate *TimeResponse `json:"tokenExpireDate"`
}

type InitApmPaymentResponse struct {
	PaymentId           *int64               `json:"paymentId"`
	RedirectUrl         *string              `json:"redirectUrl"`
	HtmlContent         *string              `json:"htmlContent"`
	PaymentStatus       *PaymentStatus       `json:"paymentStatus"`
	ApmAdditionalAction *ApmAdditionalAction `json:"additionalAction"`
	PaymentError        *PaymentError        `json:"paymentError"`
	AdditionalData      map[string]any       `json:"additionalData"`
}

type CompleteApmPaymentResponse struct {
	PaymentId     *int64         `json:"paymentId"`
	PaymentStatus *PaymentStatus `json:"paymentStatus"`
	PaymentError  *PaymentError  `json:"paymentError"`
}

type InitPosApmPaymentResponse struct {
	PaymentId        *int64            `json:"paymentId"`
	HtmlContent      *string           `json:"htmlContent"`
	PaymentStatus    *PaymentStatus    `json:"paymentStatus"`
	AdditionalAction *AdditionalAction `json:"additionalAction"`
	AdditionalData   map[string]any    `json:"additionalData"`
	PaymentError     *PaymentError     `json:"paymentError"`
}

type DepositPaymentResponse struct {
	Id                       *int64             `json:"id"`
	CreatedDate              *TimeResponse      `json:"createdDate"`
	Price                    *float64           `json:"price"`
	Currency                 *string            `json:"currency"`
	BuyerMemberId            *int64             `json:"buyerMemberId"`
	ConversationId           *string            `json:"conversationId"`
	BankCommissionRate       *float64           `json:"bankCommissionRate"`
	BankCommissionRateAmount *float64           `json:"bankCommissionRateAmount"`
	AuthCode                 *string            `json:"authCode"`
	HostReference            *string            `json:"hostReference"`
	TransId                  *string            `json:"transId"`
	OrderId                  *string            `json:"orderId"`
	PaymentType              *PaymentType       `json:"paymentType"`
	PaymentStatus            *PaymentStatus     `json:"paymentStatus"`
	CardUserKey              *string            `json:"cardUserKey"`
	CardToken                *string            `json:"cardToken"`
	WalletTransaction        *WalletTransaction `json:"walletTransaction"`
	FraudId                  *int64             `json:"fraudId"`
	FraudAction              *FraudAction       `json:"fraudAction"`
}

type ApmDepositPaymentResponse struct {
	PaymentId           *int64               `json:"paymentId"`
	RedirectUrl         *string              `json:"redirectUrl"`
	PaymentStatus       *PaymentStatus       `json:"paymentStatus"`
	ApmAdditionalAction *ApmAdditionalAction `json:"additionalAction"`
	PaymentError        *PaymentError        `json:"paymentError"`
	WalletTransaction   *WalletTransaction   `json:"walletTransaction"`
}

type PayoutAccountResponse struct {
	Id                  int64             `json:"id"`
	AccountType         PayoutAccountType `json:"type"`
	ExternalAccountId   string            `json:"externalAccountId"`
	Currency            Currency          `json:"currency"`
	AccountOwner        AccountOwner      `json:"accountOwner"`
	SubMerchantMemberId *int64            `json:"subMerchantMemberId"`
}

type MasterpassPaymentThreeDSInitResponse struct {
	PaymentId *int64  `json:"paymentId"`
	ReturnUrl *string `json:"returnUrl"`
}

type MasterpassPaymentTokenGenerateResponse struct {
	Token       *string `json:"token"`
	ReferenceId *string `json:"referenceId"`
	OrderNo     *string `json:"orderNo"`
}

type RefundWalletTransactionRequest struct {
	RefundPrice float64 `json:"refundPrice"`
}

type RemittanceRequest struct {
	MemberId             int64    `json:"memberId"`
	Price                float64  `json:"price"`
	Currency             Currency `json:"currency,omitempty"`
	Description          string   `json:"description"`
	RemittanceReasonType string   `json:"remittanceReasonType"`
}

type CreateMemberWalletRequest struct {
	NegativeAmountLimit float64  `json:"negativeAmountLimit"`
	Currency            Currency `json:"currency"`
}

type UpdateMemberWalletRequest struct {
	NegativeAmountLimit float64 `json:"negativeAmountLimit"`
}

type CreateWithdrawRequest struct {
	MemberId    int64    `json:"memberId"`
	Price       float64  `json:"price"`
	Description string   `json:"description"`
	Currency    Currency `json:"currency"`
}

type SearchWalletTransactionsRequest struct {
	WalletTransactionTypes []WalletTransactionType `schema:"walletTransactionTypes,omitempty"`
	MinAmount              float64                 `schema:"minWithdrawPrice,omitempty"`
	MaxAmount              float64                 `schema:"maxWithdrawPrice,omitempty"`
	MinCreatedDate         time.Time               `schema:"minCreatedDate,omitempty"`
	MaxCreatedDate         time.Time               `schema:"maxCreatedDate,omitempty"`
	Page                   int                     `schema:"page,omitempty"`
	Size                   int                     `schema:"size,omitempty"`
}

type SearchWithdrawsRequest struct {
	MemberId         int64     `schema:"walletId,omitempty"`
	PayoutStatus     string    `schema:"payoutStatus,omitempty"`
	Currency         Currency  `schema:"currency,omitempty"`
	MinWithdrawPrice float64   `schema:"minWithdrawPrice,omitempty"`
	MaxWithdrawPrice float64   `schema:"maxWithdrawPrice,omitempty"`
	MinCreatedDate   time.Time `schema:"minCreatedDate,omitempty"`
	MaxCreatedDate   time.Time `schema:"maxCreatedDate,omitempty"`
	Page             int       `schema:"page,omitempty"`
	Size             int       `schema:"size,omitempty"`
}

type MemberWalletResponse struct {
	Id                  *int64        `json:"id"`
	CreatedDate         *TimeResponse `json:"createdDate"`
	UpdatedDate         *TimeResponse `json:"updatedDate"`
	Amount              *float64      `json:"amount"`
	WithdrawalAmount    *float64      `json:"withdrawalAmount"`
	NegativeAmountLimit *float64      `json:"negativeAmountLimit"`
	Currency            *Currency     `json:"currency"`
	MemberId            *int64        `json:"memberId"`
}

type RemittanceResponse struct {
	Id                   *int64        `json:"id"`
	CreatedDate          *TimeResponse `json:"createdDate"`
	Status               *Status       `json:"status"`
	Price                *float64      `json:"price"`
	Currency             Currency      `json:"currency,omitempty"`
	MemberId             *int64        `json:"memberId"`
	RemittanceType       *string       `json:"remittanceType"`
	RemittanceReasonType *string       `json:"remittanceReasonType"`
	Description          *string       `json:"description"`
}

type WithdrawResponse struct {
	Id           *int64                   `json:"id"`
	CreatedDate  *TimeResponse            `json:"createdDate"`
	Status       *Status                  `json:"status"`
	MemberId     *int64                   `json:"memberId"`
	PayoutId     *int64                   `json:"payoutId"`
	Price        *float64                 `json:"price"`
	Description  *string                  `json:"description"`
	Currency     *Currency                `json:"currency"`
	PayoutStatus *TransactionPayoutStatus `json:"payoutStatus"`
}

type RefundWalletTransactionResponse struct {
	Id                  *int64                                  `json:"id"`
	CreatedDate         *TimeResponse                           `json:"createdDate"`
	RefundStatus        *string                                 `json:"refundStatus"`
	RefundPrice         *float64                                `json:"refundPrice"`
	AuthCode            *string                                 `json:"authCode"`
	HostReference       *string                                 `json:"hostReference"`
	TransId             *string                                 `json:"transId"`
	TransactionId       *int64                                  `json:"transactionId"`
	WalletTransactionId *int64                                  `json:"walletTransactionId"`
	PaymentError        *PaymentError                           `json:"paymentError"`
	TransactionType     *WalletTransactionRefundTransactionType `json:"transactionType"`
}

type SearchWalletTransactionsResponse struct {
	ID                    *int64        `json:"id"`
	CreatedDate           *TimeResponse `json:"createdDate"`
	WalletTransactionType *string       `json:"walletTransactionType"`
	Amount                *float64      `json:"amount"`
	TransactionID         *int64        `json:"transactionId"`
	WalletID              *int64        `json:"walletId"`
}

type ResetMerchantMemberWalletBalanceRequest struct {
	WalletAmount float64 `json:"walletAmount"`
}

type RetrieveWalletTransactionRefundableAmountResponse struct {
	RefundableAmount *float64 `json:"refundableAmount"`
}

type FundTransferDepositPaymentResponse struct {
	Price             *float64           `json:"price"`
	Currency          *string            `json:"currency"`
	ConversationId    *string            `json:"conversationId"`
	BuyerMemberId     *int64             `json:"buyerMemberId"`
	WalletTransaction *WalletTransaction `json:"walletTransaction"`
}

type WalletTransaction struct {
	Id                    *int64                 `json:"id"`
	WalletTransactionType *WalletTransactionType `json:"walletTransactionType"`
	Amount                *float64               `json:"amount"`
	WalletId              *int64                 `json:"walletId"`
}

type CustomInstallment struct {
	Number     int     `json:"number,omitempty"`
	TotalPrice float64 `json:"totalPrice,omitempty"`
}

type GarantiPayInstallment struct {
	Number     int     `json:"number,omitempty"`
	TotalPrice float64 `json:"totalPrice,omitempty"`
}

type InitGarantiPayPaymentResponse struct {
	HtmlContent *string `json:"htmlContent"`
	PaymentId   *int64  `json:"paymentId"`
}

type RetrieveLoyaltiesResponse struct {
	CardBrand *string      `json:"cardBrand"`
	Force3ds  *bool        `json:"force3ds"`
	Pos       *MerchantPos `json:"pos"`
	Loyalties []Loyalty    `json:"loyalties"`
}

type PaymentTransactionRefundResponse struct {
	Id                    *int64                 `json:"id"`
	CreatedDate           *TimeResponse          `json:"createdDate"`
	Status                *RefundStatus          `json:"status"`
	RefundDestinationType *RefundDestinationType `json:"refundDestinationType"`
	RefundPrice           *float64               `json:"refundPrice"`
	RefundBankPrice       *float64               `json:"refundBankPrice"`
	RefundWalletPrice     *float64               `json:"refundWalletPrice"`
	ConversationId        *string                `json:"conversationId"`
	AuthCode              *string                `json:"authCode"`
	HostReference         *string                `json:"hostReference"`
	TransId               *string                `json:"transId"`
	IsAfterSettlement     *bool                  `json:"isAfterSettlement"`
	PaymentTransactionId  *int64                 `json:"paymentTransactionId"`
	Currency              *Currency              `json:"currency"`
	PaymentId             *int64                 `json:"paymentId"`
}

type PaymentRefundResponse struct {
	Id                        *int64                             `json:"id"`
	CreatedDate               *TimeResponse                      `json:"createdDate"`
	Status                    *RefundStatus                      `json:"status"`
	RefundDestinationType     *RefundDestinationType             `json:"refundDestinationType"`
	RefundPrice               *float64                           `json:"refundPrice"`
	RefundBankPrice           *float64                           `json:"refundBankPrice"`
	RefundWalletPrice         *float64                           `json:"refundWalletPrice"`
	ConversationId            *string                            `json:"conversationId"`
	AuthCode                  *string                            `json:"authCode"`
	HostReference             *string                            `json:"hostReference"`
	TransId                   *string                            `json:"transId"`
	PaymentId                 *int64                             `json:"paymentId"`
	RefundType                *RefundType                        `json:"refundType"`
	Currency                  *Currency                          `json:"currency"`
	PaymentTransactionRefunds []PaymentTransactionRefundResponse `json:"paymentTransactionRefunds"`
}

type StoredCardResponse struct {
	BinNumber        *string           `json:"binNumber"`
	LastFourDigits   *string           `json:"lastFourDigits"`
	CardUserKey      *string           `json:"cardUserKey"`
	CardToken        *string           `json:"cardToken"`
	CardHolderName   *string           `json:"cardHolderName"`
	CardAlias        *string           `json:"cardAlias"`
	CardType         *CardType         `json:"cardType"`
	CardAssociation  *CardAssociation  `json:"cardAssociation"`
	CardExpiryStatus *CardExpiryStatus `json:"cardExpiryStatus"`
	CardBrand        *string           `json:"cardBrand"`
	CardBankName     *string           `json:"cardBankName"`
	CardBankId       *int64            `json:"cardBankId"`
	CreatedAt        *TimeResponse     `json:"createdAt"`
}

type PaymentTransactionsApprovalResponse struct {
	PaymentTransactionId *int64          `json:"paymentTransactionId"`
	ApprovalStatus       *ApprovalStatus `json:"approvalStatus"`
	FailedReason         *string         `json:"failedReason"`
}

type CheckMasterpassUserResponse struct {
	IsEligibleToUseMasterpass             *bool   `json:"isEligibleToUseMasterpass"`
	IsAnyCardSavedInCustomerProgram       *bool   `json:"isAnyCardSavedInCustomerProgram"`
	HasMasterpassAccount                  *bool   `json:"hasMasterpassAccount"`
	HashAnyCardSavedInMasterpassAccount   *bool   `json:"hashAnyCardSavedInMasterpassAccount"`
	IsMasterpassAccountLinkedWithMerchant *bool   `json:"isMasterpassAccountLinkedWithMerchant"`
	IsMasterpassAccountLocked             *bool   `json:"isMasterpassAccountLocked"`
	IsPhoneNumberUpdatedInAnotherMerchant *bool   `json:"isPhoneNumberUpdatedInAnotherMerchant"`
	AccountStatus                         *string `json:"accountStatus"`
}

type InitBkmExpressResponse struct {
	Id    *string `json:"id"`
	Path  *string `json:"path"`
	Token *string `json:"token"`
}

type InstallmentPrice struct {
	InstallmentPrice       *float64 `json:"installmentPrice"`
	BankCommissionRate     *float64 `json:"bankCommissionRate"`
	MerchantCommissionRate *float64 `json:"merchantCommissionRate"`
	TotalPrice             *float64 `json:"totalPrice"`
	InstallmentNumber      *int     `json:"installmentNumber"`
	InstallmentLabel       *string  `json:"installmentLabel"`
	Force3ds               *bool    `json:"force3ds"`
	CvcRequired            *bool    `json:"cvcRequired"`
}

type SearchInstallmentsRequest struct {
	BinNumber                               string   `schema:"binNumber"`
	Price                                   float64  `schema:"price"`
	Currency                                Currency `schema:"currency"`
	DistinctCardBrandsWithLowestCommissions bool     `schema:"distinctCardBrandsWithLowestCommissions"`
	LoyaltyExists                           bool     `schema:"loyaltyExists"`
}

type MerchantApmListResponse struct {
	Items []MerchantApmResponse `json:"items"`
}

type MerchantApmResponse struct {
	Id                  int64      `json:"id"`
	Status              Status     `json:"status"`
	Name                string     `json:"name"`
	ApmType             ApmType    `json:"apmType"`
	Hostname            string     `json:"hostname"`
	SupportedCurrencies []Currency `json:"supportedCurrencies"`
}
type InstallmentListResponse struct {
	Items []InstallmentResponse `json:"items"`
}

type InstallmentResponse struct {
	BinNumber         *string            `json:"binNumber"`
	Price             *float64           `json:"price"`
	CardType          *string            `json:"cardType"`
	CardAssociation   *string            `json:"cardAssociation"`
	CardBrand         *string            `json:"cardBrand"`
	BankName          *string            `json:"bankName"`
	BankCode          *int               `json:"bankCode"`
	Force3Ds          *bool              `json:"force3ds"`
	CvcRequired       *bool              `json:"cvcRequired"`
	Commercial        *bool              `json:"commercial"`
	PosAlias          *string            `json:"posAlias"`
	InstallmentPrices []InstallmentPrice `json:"installmentPrices"`
}

type InstantTransferBanksResponse struct {
	Items []InstantTransferBank `json:"items"`
}

type InstantTransferBank struct {
	BankCode    *string `json:"bankCode"`
	BankName    *string `json:"bankName"`
	BankLogoUrl *string `json:"bankLogoUrl"`
}

type RetrieveBinNumberResponse struct {
	BinNumber       *string `json:"binNumber"`
	CardType        *string `json:"cardType"`
	CardAssociation *string `json:"cardAssociation"`
	CardBrand       *string `json:"cardBrand"`
	BankName        *string `json:"bankName"`
	BankCode        *int    `json:"bankCode"`
	Commercial      *bool   `json:"commercial"`
}

type CreateMemberRequest struct {
	MemberExternalId                         string                        `json:"memberExternalId,omitempty"`
	Name                                     string                        `json:"name,omitempty"`
	Address                                  string                        `json:"address,omitempty"`
	Email                                    string                        `json:"email,omitempty"`
	PhoneNumber                              string                        `json:"phoneNumber,omitempty"`
	ContactName                              string                        `json:"contactName,omitempty"`
	ContactSurname                           string                        `json:"contactSurname,omitempty"`
	MemberType                               MemberType                    `json:"memberType,omitempty"`
	LegalCompanyTitle                        string                        `json:"legalCompanyTitle,omitempty"`
	TaxOffice                                string                        `json:"taxOffice,omitempty"`
	TaxNumber                                string                        `json:"taxNumber,omitempty"`
	Iban                                     string                        `json:"iban,omitempty"`
	SettlementEarningsDestination            SettlementEarningsDestination `json:"settlementEarningsDestination,omitempty"`
	IsBuyer                                  bool                          `json:"isBuyer,omitempty"`
	IsSubMerchant                            bool                          `json:"isSubMerchant,omitempty"`
	SubMerchantMaximumAllowedNegativeBalance float64                       `json:"subMerchantMaximumAllowedNegativeBalance,omitempty"`
	SettlementDelayCount                     int64                         `json:"settlementDelayCount,omitempty"`
}

type UpdateMemberRequest struct {
	Name                                     string                        `json:"name,omitempty"`
	Address                                  string                        `json:"address,omitempty"`
	Email                                    string                        `json:"email,omitempty"`
	PhoneNumber                              string                        `json:"phoneNumber,omitempty"`
	ContactName                              string                        `json:"contactName,omitempty"`
	ContactSurname                           string                        `json:"contactSurname,omitempty"`
	MemberType                               MemberType                    `json:"memberType,omitempty"`
	LegalCompanyTitle                        string                        `json:"legalCompanyTitle,omitempty"`
	TaxOffice                                string                        `json:"taxOffice,omitempty"`
	TaxNumber                                string                        `json:"taxNumber,omitempty"`
	Iban                                     string                        `json:"iban,omitempty"`
	SettlementEarningsDestination            SettlementEarningsDestination `json:"settlementEarningsDestination,omitempty"`
	IsBuyer                                  bool                          `json:"isBuyer,omitempty"`
	IsSubMerchant                            bool                          `json:"isSubMerchant,omitempty"`
	SubMerchantMaximumAllowedNegativeBalance float64                       `json:"subMerchantMaximumAllowedNegativeBalance,omitempty"`
	SettlementDelayCount                     int64                         `json:"settlementDelayCount,omitempty"`
}

type SearchMembersRequest struct {
	Page             int        `schema:"page,omitempty"`
	Size             int        `schema:"size,omitempty"`
	IsBuyer          bool       `schema:"isBuyer,omitempty"`
	IsSubMerchant    bool       `schema:"isSubMerchant,omitempty"`
	Name             string     `schema:"name,omitempty"`
	MemberIds        []int64    `schema:"memberIds,omitempty"`
	MemberType       MemberType `schema:"memberType,omitempty"`
	MemberExternalId string     `schema:"memberExternalId,omitempty"`
}

type MemberResponse struct {
	Id                            *int64                         `json:"id"`
	CreatedDate                   *TimeResponse                  `json:"createdDate"`
	UpdatedDate                   *TimeResponse                  `json:"updatedDate"`
	Status                        *Status                        `json:"status"`
	IsBuyer                       *bool                          `json:"isBuyer"`
	IsSubMerchant                 *bool                          `json:"isSubMerchant"`
	MemberType                    *MemberType                    `json:"memberType"`
	MemberExternalId              *string                        `json:"memberExternalId"`
	Name                          *string                        `json:"name"`
	Email                         *string                        `json:"email"`
	Address                       *string                        `json:"address"`
	PhoneNumber                   *string                        `json:"phoneNumber"`
	ContactName                   *string                        `json:"contactName"`
	ContactSurname                *string                        `json:"contactSurname"`
	LegalCompanyTitle             *string                        `json:"legalCompanyTitle"`
	TaxOffice                     *string                        `json:"taxOffice"`
	TaxNumber                     *string                        `json:"taxNumber"`
	SettlementEarningsDestination *SettlementEarningsDestination `json:"settlementEarningsDestination"`
	Iban                          *string                        `json:"iban"`
	SettlementDelayCount          *int64                         `json:"settlementDelayCount"`
}

type CreateProductRequest struct {
	Name                string   `json:"name"`
	Channel             string   `json:"channel,omitempty"`
	OrderId             string   `json:"orderId,omitempty"`
	ConversationId      string   `json:"conversationId,omitempty"`
	ExternalId          string   `json:"externalId,omitempty"`
	Stock               int      `json:"stock,omitempty"`
	Price               float64  `json:"price"`
	Currency            Currency `json:"currency"`
	Description         string   `json:"description,omitempty"`
	MultiPayment        bool     `json:"multiPayment,omitempty"`
	EnabledInstallments []int    `json:"enabledInstallments"`
}

type UpdateProductRequest struct {
	Name                string   `json:"name"`
	Channel             string   `json:"channel,omitempty"`
	OrderId             string   `json:"orderId,omitempty"`
	ConversationId      string   `json:"conversationId,omitempty"`
	ExternalId          string   `json:"externalId,omitempty"`
	Stock               int      `json:"stock,omitempty"`
	Status              Status   `json:"status"`
	Price               float64  `json:"price"`
	Currency            Currency `json:"currency"`
	Description         string   `json:"description,omitempty"`
	EnabledInstallments []int    `json:"enabledInstallments"`
}

type SearchProductsRequest struct {
	Name     string   `schema:"name,omitempty"`
	MinPrice float64  `schema:"minPrice,omitempty"`
	MaxPrice float64  `schema:"maxPrice,omitempty"`
	Currency Currency `schema:"currency,omitempty"`
	Channel  string   `schema:"channel,omitempty"`
	Page     int      `schema:"page,omitempty"`
	Size     int      `schema:"size,omitempty"`
}

type ProductResponse struct {
	Id                  *int64    `json:"id"`
	Name                *string   `json:"name"`
	Description         *string   `json:"description"`
	OrderId             *string   `json:"orderId,omitempty"`
	ConversationId      *string   `json:"conversationId,omitempty"`
	ExternalId          *string   `json:"externalId,omitempty"`
	Status              *Status   `json:"status"`
	Price               *float64  `json:"price"`
	Currency            *Currency `json:"currency"`
	Stock               *int      `json:"stock"`
	SoldCount           *int      `json:"soldCount"`
	Token               *string   `json:"token"`
	EnabledInstallments []int     `json:"enabledInstallments"`
	Url                 *string   `json:"url"`
	QrCodeUrl           *string   `json:"qrCodeUrl"`
	Channel             *string   `json:"channel"`
}

type SearchPaymentsRequest struct {
	Page                 int             `schema:"page,omitempty"`
	Size                 int             `schema:"size,omitempty"`
	PaymentId            int64           `schema:"paymentId,omitempty"`
	PaymentTransactionId int64           `schema:"paymentTransactionId,omitempty"`
	BuyerMemberId        int64           `schema:"buyerMemberId,omitempty"`
	SubMerchantMemberId  int64           `schema:"subMerchantMemberId,omitempty"`
	ConversationId       string          `schema:"conversationId,omitempty"`
	ExternalId           string          `schema:"externalId,omitempty"`
	OrderId              string          `schema:"orderId,omitempty"`
	PaymentType          PaymentType     `schema:"paymentType,omitempty"`
	PaymentProvider      PaymentProvider `schema:"paymentProvider,omitempty"`
	PaymentStatus        PaymentStatus   `schema:"paymentStatus,omitempty"`
	PaymentSource        PaymentSource   `schema:"paymentSource,omitempty"`
	PaymentChannel       string          `schema:"paymentChannel,omitempty"`
	BinNumber            string          `schema:"binNumber,omitempty"`
	LastFourDigits       string          `schema:"lastFourDigits,omitempty"`
	Currency             Currency        `schema:"currency,omitempty"`
	MinPaidPrice         float64         `schema:"minPaidPrice,omitempty"`
	MaxPaidPrice         float64         `schema:"maxPaidPrice,omitempty"`
	Installment          int             `schema:"installment,omitempty"`
	IsThreeDS            bool            `schema:"isThreeDS,omitempty"`
	MinCreatedDate       time.Time       `schema:"minCreatedDate,omitempty"`
	MaxCreatedDate       time.Time       `schema:"maxCreatedDate,omitempty"`
}

type SearchPaymentRefundsRequest struct {
	Page           int          `schema:"page,omitempty"`
	Size           int          `schema:"size,omitempty"`
	Id             int64        `schema:"id,omitempty"`
	PaymentId      int64        `schema:"paymentId,omitempty"`
	BuyerMemberId  int64        `schema:"buyerMemberId,omitempty"`
	ConversationId string       `schema:"conversationId,omitempty"`
	Status         RefundStatus `schema:"status,omitempty"`
	Currency       Currency     `schema:"currency,omitempty"`
	MinRefundPrice float64      `schema:"minRefundPrice,omitempty"`
	MaxRefundPrice float64      `schema:"maxRefundPrice,omitempty"`
	MinCreatedDate time.Time    `schema:"minCreatedDate,omitempty"`
	MaxCreatedDate time.Time    `schema:"maxCreatedDate,omitempty"`
}

type SearchPaymentTransactionRefundsRequest struct {
	Page                 int          `schema:"page,omitempty"`
	Size                 int          `schema:"size,omitempty"`
	Id                   int64        `schema:"id,omitempty"`
	PaymentId            int64        `schema:"paymentId,omitempty"`
	PaymentTransactionId int64        `schema:"paymentTransactionId,omitempty"`
	BuyerMemberId        int64        `schema:"buyerMemberId,omitempty"`
	ConversationId       string       `schema:"conversationId,omitempty"`
	Status               RefundStatus `schema:"status,omitempty"`
	Currency             Currency     `schema:"currency,omitempty"`
	IsAfterSettlement    bool         `schema:"isAfterSettlement,omitempty"`
	MinRefundPrice       float64      `schema:"minRefundPrice,omitempty"`
	MaxRefundPrice       float64      `schema:"maxRefundPrice,omitempty"`
	MinCreatedDate       time.Time    `schema:"minCreatedDate,omitempty"`
	MaxCreatedDate       time.Time    `schema:"maxCreatedDate,omitempty"`
}

type ReportingPaymentResponse struct {
	Id                           *int64                            `json:"id"`
	CreatedDate                  *TimeResponse                     `json:"createdDate"`
	Price                        *float64                          `json:"price"`
	PaidPrice                    *float64                          `json:"paidPrice"`
	WalletPrice                  *float64                          `json:"walletPrice"`
	Currency                     *Currency                         `json:"currency"`
	BuyerMemberId                *int64                            `json:"buyerMemberId"`
	Installment                  *int                              `json:"installment"`
	ConversationId               *string                           `json:"conversationId"`
	ExternalId                   *string                           `json:"externalId"`
	PaymentType                  *PaymentType                      `json:"paymentType"`
	PaymentProvider              *PaymentProvider                  `json:"paymentProvider"`
	PaymentSource                *PaymentSource                    `json:"paymentSource"`
	PaymentGroup                 *PaymentGroup                     `json:"paymentGroup"`
	PaymentStatus                *PaymentStatus                    `json:"paymentStatus"`
	PaymentPhase                 *PaymentPhase                     `json:"paymentPhase"`
	PaymentChannel               *string                           `json:"paymentChannel"`
	IsThreeDS                    *bool                             `json:"isThreeDS"`
	MerchantCommissionRate       *float64                          `json:"merchantCommissionRate"`
	MerchantCommissionRateAmount *float64                          `json:"merchantCommissionRateAmount"`
	BankCommissionRate           *float64                          `json:"bankCommissionRate"`
	BankCommissionRateAmount     *float64                          `json:"bankCommissionRateAmount"`
	PaidWithStoredCard           *bool                             `json:"paidWithStoredCard"`
	BinNumber                    *string                           `json:"binNumber"`
	LastFourDigits               *string                           `json:"lastFourDigits"`
	AuthCode                     *string                           `json:"authCode"`
	HostReference                *string                           `json:"hostReference"`
	OrderId                      *string                           `json:"orderId"`
	TransId                      *string                           `json:"transId"`
	CardHolderName               *string                           `json:"cardHolderName"`
	BankCardHolderName           *string                           `json:"bankCardHolderName"`
	CardType                     *CardType                         `json:"cardType"`
	CardAssociation              *CardAssociation                  `json:"cardAssociation"`
	CardBrand                    *string                           `json:"cardBrand"`
	RequestedPosAlias            *string                           `json:"requestedPosAlias"`
	RetryCount                   *int                              `json:"retryCount"`
	RefundablePrice              *float64                          `json:"refundablePrice"`
	RefundStatus                 *PaymentRefundStatus              `json:"refundStatus"`
	CardIssuerBankName           *string                           `json:"cardIssuerBankName"`
	MdStatus                     *int                              `json:"mdStatus"`
	FraudId                      *int64                            `json:"fraudId"`
	FraudAction                  *FraudAction                      `json:"fraudAction"`
	BuyerMember                  *MemberResponse                   `json:"buyerMember"`
	Refunds                      *[]ReportingPaymentRefundResponse `json:"refunds"`
	Pos                          *MerchantPos                      `json:"pos"`
	Loyalty                      *Loyalty                          `json:"loyalty"`
	PaymentError                 *PaymentError                     `json:"paymentError"`
}

type ReportingPaymentRefundResponse struct {
	Id                    *int64                 `json:"id"`
	CreatedDate           *TimeResponse          `json:"createdDate"`
	Status                *RefundStatus          `json:"status"`
	RefundDestinationType *RefundDestinationType `json:"refundDestinationType"`
	RefundPrice           *float64               `json:"refundPrice"`
	RefundBankPrice       *float64               `json:"refundBankPrice"`
	RefundWalletPrice     *float64               `json:"refundWalletPrice"`
	ConversationId        *string                `json:"conversationId"`
	AuthCode              *string                `json:"authCode"`
	HostReference         *string                `json:"hostReference"`
	PaymentType           *PaymentType           `json:"paymentType"`
	TransId               *string                `json:"transId"`
	PaymentId             *int64                 `json:"paymentId"`
	PaymentError          *PaymentError          `json:"paymentError"`
}

type ReportingPaymentTransactionRefundResponse struct {
	Id                    *int64                 `json:"id"`
	CreatedDate           *TimeResponse          `json:"createdDate"`
	Status                *RefundStatus          `json:"status"`
	RefundDestinationType *RefundDestinationType `json:"refundDestinationType"`
	RefundPrice           *float64               `json:"refundPrice"`
	RefundBankPrice       *float64               `json:"refundBankPrice"`
	RefundWalletPrice     *float64               `json:"refundWalletPrice"`
	ConversationId        *string                `json:"conversationId"`
	AuthCode              *string                `json:"authCode"`
	HostReference         *string                `json:"hostReference"`
	TransId               *string                `json:"transId"`
	IsAfterSettlement     *bool                  `json:"isAfterSettlement"`
	PaymentTransactionId  *int64                 `json:"paymentTransactionId"`
	PaymentError          *PaymentError          `json:"paymentError"`
}

type ReportingPaymentTransactionResponse struct {
	Id                            *int64               `json:"id"`
	Name                          *string              `json:"name"`
	ExternalId                    *string              `json:"externalId"`
	Price                         *float64             `json:"price"`
	PaidPrice                     *float64             `json:"paidPrice"`
	WalletPrice                   *float64             `json:"walletPrice"`
	MerchantCommissionRate        *float64             `json:"merchantCommissionRate"`
	MerchantCommissionRateAmount  *float64             `json:"merchantCommissionRateAmount"`
	MerchantPayoutAmount          *float64             `json:"merchantPayoutAmount"`
	SubMerchantMemberId           *int64               `json:"subMerchantMemberId"`
	SubMerchantMemberPrice        *float64             `json:"subMerchantMemberPrice"`
	SubMerchantMemberPayoutRate   *float64             `json:"subMerchantMemberPayoutRate"`
	SubMerchantMemberPayoutAmount *float64             `json:"subMerchantMemberPayoutAmount"`
	TransactionStatus             *TransactionStatus   `json:"transactionStatus"`
	BlockageResolvedDate          *TimeResponse        `json:"blockageResolvedDate"`
	CreatedDate                   *TimeResponse        `json:"createdDate"`
	TransactionStatusDate         *TimeResponse        `json:"transactionStatusDate"`
	RefundablePrice               *float64             `json:"refundablePrice"`
	BankCommissionRate            *float64             `json:"bankCommissionRate"`
	BankCommissionRateAmount      *float64             `json:"bankCommissionRateAmount"`
	Payout                        *Payout              `json:"payout"`
	SubMerchantMember             *MemberResponse      `json:"subMerchantMember"`
	RefundStatus                  *PaymentRefundStatus `json:"refundStatus"`
	PayoutStatus                  *PayoutStatus        `json:"payoutStatus"`
}

type InitBnplPaymentResponse struct {
	PaymentId        int64               `json:"paymentId"`
	RedirectUrl      string              `json:"redirectUrl"`
	PaymentStatus    PaymentStatus       `json:"paymentStatus"`
	AdditionalAction ApmAdditionalAction `json:"additionalAction"`
	PaymentError     PaymentError        `json:"paymentError"`
	AdditionalData   map[string]any      `json:"additionalData"`
}

type BnplPaymentVerifyResponse struct {
	PaymentStatus    PaymentStatus       `json:"paymentStatus"`
	AdditionalAction ApmAdditionalAction `json:"additionalAction"`
	PaymentError     PaymentError        `json:"paymentError"`
}

type BnplPaymentOfferResponse struct {
	OfferId        string           `json:"offerId"`
	Price          *float64         `json:"price"`
	BnplBankOffers *[]BnplBankOffer `json:"nnplBankOffers"`
}

type BnplBankOffer struct {
	BankCode                      string               `json:"bankCode"`
	BankName                      string               `json:"bankName"`
	BankIconUrl                   string               `json:"bankIconUrl"`
	BankTableBannerMessage        string               `json:"bankTableBannerMessage"`
	BankSmallBannerMessage        string               `json:"bankSmallBannerMessage"`
	PreApprovedApplicationId      string               `json:"preApprovedApplicationId"`
	IsSupportNonCustomer          bool                 `json:"isSupportNonCustomer"`
	IsPaymentPlanCalculatedByBank bool                 `json:"isPaymentPlanCalculatedByBank"`
	BnplBankOfferTerm             *[]BnplBankOfferTerm `json:"bankOfferTerms"`
}

type BnplBankOfferTerm struct {
	Term               int64    `json:"term"`
	Amount             *float64 `json:"amount"`
	TotalAmount        *float64 `json:"totalAmount"`
	InterestRate       *float64 `json:"interestRate"`
	AnnualInterestRate *float64 `json:"annualInterestRate"`
}

type Payout struct {
	PaidPrice                     *float64  `json:"paidPrice"`
	Parity                        *float64  `json:"parity"`
	Currency                      *Currency `json:"currency"`
	MerchantPayoutAmount          *float64  `json:"merchantPayoutAmount"`
	SubMerchantMemberPayoutAmount *float64  `json:"subMerchantMemberPayoutAmount"`
}

type PayoutStatus struct {
	MerchantStatus              *TransactionPayoutStatus `json:"merchantStatus"`
	MerchantStatusDate          *TimeResponse            `json:"merchantStatusDate"`
	SubMerchantMemberStatus     *TransactionPayoutStatus `json:"subMerchantMemberStatus"`
	SubMerchantMemberStatusDate *TimeResponse            `json:"subMerchantMemberStatusDate"`
}

type CreateInstantWalletSettlementRequest struct {
	ExcludedSubMerchantMemberIds []int64
}

type CreateInstantWalletSettlementResponse struct {
	SettlementResultStatus *string `json:"settlementResultStatus"`
}

type SearchPayoutCompletedTransactionsRequest struct {
	SettlementFileId int64          `schema:"settlementFileId,omitempty"`
	SettlementType   SettlementType `schema:"settlementType,omitempty"`
	StartDate        time.Time      `schema:"startDate,omitempty"`
	EndDate          time.Time      `schema:"endDate,omitempty"`
	Page             int            `schema:"page"`
	Size             int            `schema:"size"`
}

type SearchPayoutBouncedTransactionsRequest struct {
	StartDate time.Time `schema:"startDate,omitempty"`
	EndDate   time.Time `schema:"endDate,omitempty"`
}

type SearchPayoutRowRequest struct {
	Page       int        `schema:"page,omitempty"`
	Size       int        `schema:"size,omitempty"`
	FileStatus FileStatus `schema:"fileStatus,omitempty"`
	StartDate  time.Time  `schema:"startDate,omitempty"`
	EndDate    time.Time  `schema:"endDate,omitempty"`
}

type RetrievePayoutDetailsRequest struct {
	PayoutDetailId int64
}

type SearchPayoutCompletedTransactionsResponse struct {
	PayoutId                      *int64        `json:"payoutId"`
	TransactionId                 *int64        `json:"transactionId"`
	TransactionType               *string       `json:"transactionType"`
	PayoutAmount                  *float64      `json:"payoutAmount"`
	PayoutDate                    *TimeResponse `json:"payoutDate"`
	Currency                      *Currency     `json:"currency"`
	MerchantId                    *int64        `json:"merchantId"`
	MerchantType                  *string       `json:"merchantType"`
	SettlementEarningsDestination *string       `json:"settlementEarningsDestination"`
	SettlementSource              *string       `json:"settlementSource"`
}

type SearchPayoutBouncedTransactionsResponse struct {
	Id                *int64        `json:"id"`
	Iban              *string       `json:"iban"`
	CreatedDate       *TimeResponse `json:"createdDate"`
	UpdatedDate       *TimeResponse `json:"updatedDate"`
	PayoutId          *int64        `json:"payoutId"`
	PayoutAmount      *float64      `json:"payoutAmount"`
	ContactName       *string       `json:"contactName"`
	ContactSurname    *string       `json:"contactSurname"`
	LegalCompanyTitle *string       `json:"legalCompanyTitle"`
	RowDescription    *string       `json:"rowDescription"`
}

type PayoutDetailResponse struct {
	RowDescription                *string                           `json:"rowDescription"`
	PayoutDate                    *TimeResponse                     `json:"payoutDate"`
	Name                          *string                           `json:"name"`
	Iban                          *string                           `json:"iban"`
	PayoutAmount                  *float64                          `json:"payoutAmount"`
	Currency                      *string                           `json:"currency"`
	MerchantId                    *int64                            `json:"merchantId"`
	MerchantType                  *string                           `json:"merchantType"`
	SettlementEarningsDestination *string                           `json:"settlementEarningsDestination"`
	SettlementSource              *string                           `json:"settlementSource"`
	BounceStatus                  *string                           `json:"bounceStatus"`
	PayoutTransactions            []PayoutDetailTransactionResponse `json:"payoutTransactions"`
}

type PayoutDetailTransactionResponse struct {
	TransactionId   *int64   `json:"transactionId"`
	TransactionType *string  `json:"transactionType"`
	PayoutAmount    *float64 `json:"payoutAmount"`
}

type RetrieveDailyTransactionReportRequest struct {
	ReportDate Date           `schema:"reportDate,omitempty"`
	FileType   ReportFileType `schema:"fileType,omitempty"`
}

type RetrieveDailyPaymentReportRequest struct {
	ReportDate Date           `schema:"reportDate,omitempty"`
	FileType   ReportFileType `schema:"fileType,omitempty"`
}

type SearchFraudChecksRequest struct {
	Page           int              `schema:"page,omitempty"`
	Size           int              `schema:"size,omitempty"`
	Action         FraudAction      `schema:"action,omitempty"`
	CheckStatus    FraudCheckStatus `schema:"checkStatus,omitempty"`
	MinCreatedDate time.Time        `schema:"minCreatedDate,omitempty"`
	MaxCreatedDate time.Time        `schema:"maxCreatedDate,omitempty"`
	RuleId         int64            `schema:"ruleId,omitempty"`
	PaymentId      int64            `schema:"paymentId,omitempty"`
	PaymentStatus  PaymentStatus    `schema:"paymentStatus,omitempty"`
}

type FraudCheckResponse struct {
	Id             *int64            `json:"id"`
	Status         *Status           `json:"status"`
	Action         *FraudAction      `json:"action"`
	CheckStatus    *FraudCheckStatus `json:"checkStatus"`
	PaymentData    *FraudPaymentData `json:"paymentData"`
	RuleId         *int64            `json:"ruleId"`
	RuleName       *string           `json:"ruleName"`
	RuleConditions *string           `json:"ruleConditions"`
	PaymentId      *int64            `json:"paymentId"`
	PaymentStatus  *PaymentStatus    `json:"paymentStatus"`
}

type FraudPaymentData struct {
	PaymentDate    *time.Time `json:"paymentDate"`
	ConversationId *string    `json:"conversationId"`
	PaidPrice      *float64   `json:"paidPrice"`
	Currency       *Currency  `json:"currency"`
	BuyerId        *int64     `json:"buyerId"`
	ClientIp       *string    `json:"clientIp"`
}

type FraudValueListRequest struct {
	ListName          string         `json:"listName,omitempty"`
	Type              FraudValueType `json:"type,omitempty"`
	Label             string         `json:"label,omitempty"`
	Value             string         `json:"value,omitempty"`
	PaymentId         int64          `json:"paymentId"`
	DurationInSeconds int            `json:"durationInSeconds,omitempty"`
}

type FraudValuesResponse struct {
	Name   string       `json:"name"`
	Values []FraudValue `json:"values"`
}

type FraudValue struct {
	Id              *string `json:"id"`
	Label           *string `json:"label"`
	Value           *string `json:"value"`
	ExpireInSeconds *int    `json:"expireInSeconds"`
}

type PayoutRowResponse struct {
	Name         *string       `json:"name"`
	Iban         *string       `json:"iban"`
	PayoutId     *int64        `json:"payoutId"`
	MerchantId   *int64        `json:"merchantId"`
	MerchantType *string       `json:"merchantType"`
	PayoutAmount *float64      `json:"payoutAmount"`
	PayoutDate   *TimeResponse `json:"payoutDate"`
}

type WebhookData struct {
	EventType      WebhookEventType
	EventTime      time.Time
	EventTimestamp int64
	Status         WebhookStatus
	PayloadId      string
}

type SearchBankAccountTrackingRecordRequest struct {
	SenderName    string    `schema:"senderName,omitempty"`
	SenderIban    string    `schema:"senderIban,omitempty"`
	Description   string    `schema:"description,omitempty"`
	Currency      Currency  `schema:"currency,omitempty"`
	MinRecordDate time.Time `schema:"minRecordDate,omitempty"`
	MaxRecordDate time.Time `schema:"maxRecordDate,omitempty"`
	Page          int       `schema:"page,omitempty"`
	Size          int       `schema:"size,omitempty"`
}

type BankAccountTrackingRecordResponse struct {
	Id                        int64                     `json:"id"`
	Key                       string                    `json:"key"`
	SenderName                string                    `json:"senderName"`
	SenderIban                string                    `json:"senderIban"`
	Description               string                    `json:"description"`
	Currency                  Currency                  `json:"currency"`
	Amount                    float64                   `json:"amount"`
	RecordDate                TimeResponse              `json:"recordDate"`
	RecordType                RecordType                `json:"recordType"`
	BankAccountTrackingSource BankAccountTrackingSource `json:"bankAccountTrackingSource"`
}

type RequestOptions struct {
	BaseURL   string
	ApiKey    string
	SecretKey string
}

type Response[T any] struct {
	Data   *T             `json:"data"`
	Errors *ErrorResponse `json:"errors"`
}

func (r Response[ErrorResponse]) Error() string {
	if r.Errors.ErrorGroup != nil {
		return *r.Errors.ErrorGroup + "-" + *r.Errors.ErrorCode + "-" + *r.Errors.ErrorDescription
	}

	return *r.Errors.ErrorCode + "-" + *r.Errors.ErrorDescription
}

type ErrorResponse struct {
	ErrorGroup       *string `json:"errorGroup"`
	ErrorDescription *string `json:"errorDescription"`
	ErrorCode        *string `json:"errorCode"`
}

type DataResponse[T any] struct {
	Items     []T   `json:"items"`
	Page      int   `json:"page"`
	Size      int   `json:"size"`
	TotalSize int64 `json:"totalSize"`
}

type MerchantPos struct {
	Id     *int64  `json:"id"`
	Name   *string `json:"name"`
	Alias  *string `json:"alias"`
	BankId *int64  `json:"bankId"`
}

type Reward struct {
	CardRewardMoney *float64 `json:"cardRewardMoney,omitempty"`
	FirmRewardMoney *float64 `json:"firmRewardMoney,omitempty"`
}

type Loyalty struct {
	LoyaltyType *LoyaltyType `json:"type,omitempty"`
	Reward      *Reward      `json:"reward,omitempty"`
	Message     *string      `json:"message,omitempty"`
}

type TokenizedCard struct {
	TokenizedCardType TokenizedCardType      `json:"type,omitempty"`
	Data              map[string]interface{} `json:"data,omitempty"`
}

type Card struct {
	CardHolderName               string         `json:"cardHolderName,omitempty"`
	CardNumber                   string         `json:"cardNumber,omitempty"`
	ExpireYear                   string         `json:"expireYear,omitempty"`
	ExpireMonth                  string         `json:"expireMonth,omitempty"`
	Cvc                          string         `json:"cvc,omitempty"`
	CardAlias                    string         `json:"cardAlias,omitempty"`
	CardUserKey                  string         `json:"cardUserKey,omitempty"`
	CardToken                    string         `json:"cardToken,omitempty"`
	BinNumber                    string         `json:"binNumber,omitempty"`
	LastFourDigits               string         `json:"lastFourDigits,omitempty"`
	CardHolderIdentityNumber     string         `json:"cardHolderIdentityNumber,omitempty"`
	Loyalty                      *Loyalty       `json:"loyalty,omitempty"`
	StoreCardAfterSuccessPayment bool           `json:"storeCardAfterSuccessPayment,omitempty"`
	TokenizedCard                *TokenizedCard `json:"tokenizedCard,omitempty"`
}

type FraudCheckParameters struct {
	BuyerExternalId     string `json:"buyerExternalId,omitempty"`
	BuyerPhoneNumber    string `json:"buyerPhoneNumber,omitempty"`
	BuyerEmail          string `json:"buyerEmail,omitempty"`
	CustomFraudVariable string `json:"customFraudVariable,omitempty"`
}

type PosApmInstallment struct {
	Number     *int     `json:"number,omitempty"`
	TotalPrice *float64 `json:"totalPrice,omitempty"`
}

type PaymentItem struct {
	Name                   string  `json:"name,omitempty"`
	Price                  float64 `json:"price,omitempty"`
	ExternalId             string  `json:"externalId,omitempty"`
	SubMerchantMemberId    int64   `json:"subMerchantMemberId,omitempty"`
	SubMerchantMemberPrice float64 `json:"subMerchantMemberPrice,omitempty"`
}

type MasterpassCreatePayment struct {
	Price            float64                `json:"price,omitempty"`
	PaidPrice        float64                `json:"paidPrice,omitempty"`
	PosAlias         string                 `json:"posAlias,omitempty"`
	Installment      int                    `json:"installment,omitempty"`
	Currency         Currency               `json:"currency,omitempty"`
	PaymentGroup     PaymentGroup           `json:"paymentGroup,omitempty"`
	ConversationId   string                 `json:"conversationId,omitempty"`
	ExternalId       string                 `json:"externalId,omitempty"`
	ClientIp         string                 `json:"clientIp,omitempty"`
	PaymentPhase     PaymentPhase           `json:"paymentPhase,omitempty"`
	PaymentChannel   string                 `json:"paymentChannel,omitempty"`
	BuyerMemberId    int64                  `json:"buyerMemberId,omitempty"`
	BankOrderId      string                 `json:"bankOrderId,omitempty"`
	Items            []PaymentItem          `json:"items"`
	AdditionalParams map[string]interface{} `json:"additionalParams,omitempty"`
}

type InitBkmExpressRequest struct {
	Price              float64       `json:"price,omitempty"`
	PaidPrice          float64       `json:"paidPrice,omitempty"`
	Currency           Currency      `json:"currency,omitempty"`
	PaymentGroup       PaymentGroup  `json:"paymentGroup,omitempty"`
	ConversationId     string        `json:"conversationId,omitempty"`
	PaymentPhase       PaymentPhase  `json:"paymentPhase,omitempty"`
	BuyerMemberId      int64         `json:"buyerMemberId,omitempty"`
	BankOrderId        string        `json:"bankOrderId,omitempty"`
	Items              []PaymentItem `json:"items"`
	EnabledInstallment bool          `json:"enabledInstallments"`
}

type CreateMerchantRequest struct {
	Name               string `json:"name"`
	LegalCompanyTitle  string `json:"legalCompanyTitle"`
	Email              string `json:"email"`
	SecretWord         string `json:"secretWord,omitempty"`
	Website            string `json:"website"`
	PhoneNumber        string `json:"phoneNumber,omitempty"`
	ContactName        string `json:"contactName"`
	ContactSurname     string `json:"contactSurname"`
	ContactPhoneNumber string `json:"contactPhoneNumber"`
}

type CompleteBkmExpressRequest struct {
	Status   bool   `json:"status"`
	Message  string `json:"message"`
	TicketId string `json:"ticketId"`
}

type MerchantApiCredential struct {
	Name      string `json:"name"`
	ApiKey    string `json:"apiKey"`
	SecretKey string `json:"secretKey"`
}

type CreateMerchantResponse struct {
	Id                     *int64                  `json:"id"`
	Name                   *string                 `json:"name"`
	MerchantApiCredentials []MerchantApiCredential `json:"merchantApiCredentials"`
}

type CreateMerchantPosUser struct {
	PosUsername      string           `json:"posUsername"`
	PosPassword      string           `json:"posPassword"`
	PosUserType      PosUserType      `json:"posUserType"`
	PosOperationType PosOperationType `json:"posOperationType"`
}

type CreateMerchantPosRequest struct {
	Status                            PosStatus                   `json:"status"`
	Name                              string                      `json:"name"`
	ClientId                          string                      `json:"clientId"`
	Currency                          Currency                    `json:"currency"`
	PosnetId                          string                      `json:"posnetId,omitempty"`
	TerminalId                        string                      `json:"terminalId,omitempty"`
	ThreedsPosnetId                   string                      `json:"threedsPosnetId,omitempty"`
	ThreedsTerminalId                 string                      `json:"threedsTerminalId,omitempty"`
	ThreedsKey                        string                      `json:"threedsKey,omitempty"`
	EnableForeignCard                 bool                        `json:"enableForeignCard"`
	EnableInstallment                 bool                        `json:"enableInstallment"`
	EnablePaymentWithoutCvc           bool                        `json:"enablePaymentWithoutCvc"`
	EnableLoyalty                     bool                        `json:"enableLoyalty"`
	NewIntegration                    bool                        `json:"newIntegration"`
	OrderNumber                       int64                       `json:"orderNumber"`
	PosIntegrator                     PosIntegrator               `json:"posIntegrator"`
	EnabledPaymentAuthenticationTypes []PaymentAuthenticationType `json:"enabledPaymentAuthenticationTypes"`
	MerchantPosUsers                  []CreateMerchantPosUser     `json:"merchantPosUsers"`
}

type AutopilotState struct {
	IsThreeDsUp    bool `json:"isThreeDsUp"`
	IsNonThreeDsUp bool `json:"isNonThreeDsUp"`
}

type MerchantPosUser struct {
	Id               int64            `json:"id"`
	PosUsername      string           `json:"posUsername"`
	PosPassword      string           `json:"posPassword"`
	PosUserType      PosUserType      `json:"posUserType"`
	PosOperationType PosOperationType `json:"posOperationType"`
}
type MerchantPosResponse struct {
	Id                                int64                       `json:"id"`
	Status                            PosStatus                   `json:"status"`
	Name                              string                      `json:"name"`
	Alias                             string                      `json:"alias"`
	PosIntegrator                     PosIntegrator               `json:"posIntegrator"`
	Hostname                          string                      `json:"hostname"`
	ClientId                          string                      `json:"clientId"`
	PosCurrencyCode                   string                      `json:"posCurrencyCode"`
	Mode                              string                      `json:"mode"`
	Path                              string                      `json:"path"`
	Port                              int64                       `json:"port"`
	PosnetId                          string                      `json:"posnetId"`
	TerminalId                        string                      `json:"terminalId"`
	ThreedsPosnetId                   string                      `json:"threedsPosnetId"`
	ThreedsTerminalId                 string                      `json:"threedsTerminalId"`
	ThreedsKey                        string                      `json:"threedsKey"`
	ThreedsPath                       string                      `json:"threedsPath"`
	EnableForeignCard                 bool                        `json:"enableForeignCard"`
	EnableInstallment                 bool                        `json:"enableInstallment"`
	EnablePaymentWithoutCvc           bool                        `json:"enablePaymentWithoutCvc"`
	EnableLoyalty                     bool                        `json:"enableLoyalty"`
	NewIntegration                    bool                        `json:"newIntegration"`
	OrderNumber                       int64                       `json:"orderNumber"`
	AutopilotState                    AutopilotState              `json:"autopilotState"`
	Currency                          Currency                    `json:"currency"`
	BankId                            int64                       `json:"bankId"`
	BankName                          string                      `json:"bankName"`
	IsPf                              bool                        `json:"isPf"`
	MerchantPosUsers                  []MerchantPosUser           `json:"merchantPosUsers"`
	SupportedCardAssociations         []CardAssociation           `json:"supportedCardAssociations"`
	EnabledPaymentAuthenticationTypes []PaymentAuthenticationType `json:"enabledPaymentAuthenticationTypes"`
}

type MerchantPosCommissionResponse struct {
	Id                                  int64     `json:"id"`
	Status                              Status    `json:"status"`
	Installment                         int64     `json:"installment"`
	BlockageDay                         int64     `json:"blockageDay"`
	InstallmentLabel                    string    `json:"installmentLabel"`
	CardBrandName                       CardBrand `json:"cardBrandName"`
	BankOnUsCreditCardCommissionRate    float64   `json:"bankOnUsCreditCardCommissionRate"`
	BankNotOnUsCreditCardCommissionRate float64   `json:"bankNotOnUsCreditCardCommissionRate"`
	BankOnUsDebitCardCommissionRate     float64   `json:"bankOnUsDebitCardCommissionRate"`
	BankNotOnUsDebitCardCommissionRate  float64   `json:"bankNotOnUsDebitCardCommissionRate"`
	BankForeignCardCommissionRate       float64   `json:"bankForeignCardCommissionRate"`
	MerchantCommissionRate              float64   `json:"merchantCommissionRate"`
}

type MultiPaymentResponse struct {
	Id                 *int64              `json:"id"`
	MultiPaymentStatus *MultiPaymentStatus `json:"multiPaymentStatus"`
	Token              *string             `json:"token"`
	ConversationId     *string             `json:"conversationId"`
	ExternalId         *string             `json:"externalId"`
	PaidPrice          *float64            `json:"paidPrice"`
	RemainingAmount    *float64            `json:"remainingAmount"`
	TokenExpireDate    *TimeResponse       `json:"tokenExpireDate"`
	PaymentIds         []int64             `json:"paymentIds"`
}

type CreateMerchantPosCommission struct {
	CardBrandName                       CardBrand `json:"cardBrandName,omitempty"`
	Installment                         int64     `json:"installment"`
	Status                              Status    `json:"status"`
	BlockageDay                         int64     `json:"blockageDay"`
	InstallmentLabel                    string    `json:"installmentLabel,omitempty"`
	BankOnUsCreditCardCommissionRate    float64   `json:"bankOnUsCreditCardCommissionRate"`
	BankOnUsDebitCardCommissionRate     float64   `json:"bankOnUsDebitCardCommissionRate,omitempty"`
	BankNotOnUsCreditCardCommissionRate float64   `json:"bankNotOnUsCreditCardCommissionRate,omitempty"`
	BankNotOnUsDebitCardCommissionRate  float64   `json:"bankNotOnUsDebitCardCommissionRate,omitempty"`
	BankForeignCardCommissionRate       float64   `json:"bankForeignCardCommissionRate,omitempty"`
	MerchantCommissionRate              float64   `json:"merchantCommissionRate,omitempty"`
}

type SearchMerchantPosRequest struct {
	Name              string   `schema:"name,omitempty"`
	Alias             string   `schema:"alias,omitempty"`
	Currency          Currency `schema:"currency,omitempty"`
	EnableInstallment bool     `schema:"enableInstallment,omitempty"`
	EnableForeignCard bool     `schema:"enableForeignCard,omitempty"`
	BankName          string   `schema:"bankName,omitempty"`
	Page              int64    `schema:"page,omitempty"`
	Size              int64    `schema:"size,omitempty"`
}

type CreateMerchantPosCommissionRequest struct {
	Commissions []CreateMerchantPosCommission `json:"commissions"`
}

type InitJuzdanPaymentRequest struct {
	Price          float64       `json:"price,omitempty"`
	PaidPrice      float64       `json:"paidPrice,omitempty"`
	Currency       Currency      `json:"currency,omitempty"`
	PaymentGroup   PaymentGroup  `json:"paymentGroup,omitempty"`
	ConversationId string        `json:"conversationId,omitempty"`
	ExternalId     string        `json:"externalId,omitempty"`
	CallbackUrl    string        `json:"callbackUrl,omitempty"`
	PaymentPhase   PaymentPhase  `json:"paymentPhase,omitempty"`
	PaymentChannel string        `json:"paymentChannel,omitempty"`
	BuyerMemberId  int64         `json:"buyerMemberId,omitempty"`
	BankOrderId    string        `json:"bankOrderId,omitempty"`
	Items          []PaymentItem `json:"items"`
	ClientType     ClientType    `json:"clientType,omitempty"`
	LoanCampaignId int64         `json:"loanCampaignId,omitempty"`
}

type InitJuzdanPaymentResponse struct {
	ReferenceId string `json:"referenceId"`
	JuzdanQrUrl string `json:"juzdanQrUrl"`
}

type PaymentError ErrorResponse

type Void struct {
}
