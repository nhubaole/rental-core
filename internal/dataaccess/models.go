// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.25.0

package dataaccess

import (
	"github.com/jackc/pgx/v5/pgtype"
)

type Bank struct {
	ID        int32            `json:"id"`
	BankName  string           `json:"bank_name"`
	BankCode  string           `json:"bank_code"`
	CreatedAt pgtype.Timestamp `json:"created_at"`
	UpdatedAt pgtype.Timestamp `json:"updated_at"`
	ShortName *string          `json:"short_name"`
	Logo      *string          `json:"logo"`
}

type Billing struct {
	ID                   int32              `json:"id"`
	Code                 string             `json:"code"`
	ContractID           int32              `json:"contract_id"`
	AdditionFee          *int32             `json:"addition_fee"`
	AdditionNote         *string            `json:"addition_note"`
	TotalAmount          float64            `json:"total_amount"`
	Month                int32              `json:"month"`
	Year                 int32              `json:"year"`
	CreatedAt            pgtype.Timestamptz `json:"created_at"`
	UpdatedAt            pgtype.Timestamptz `json:"updated_at"`
	DeletedAt            pgtype.Timestamptz `json:"deleted_at"`
	OldWaterIndex        *int32             `json:"old_water_index"`
	OldElectricityIndex  *int32             `json:"old_electricity_index"`
	NewWaterIndex        *int32             `json:"new_water_index"`
	NewElectricityIndex  *int32             `json:"new_electricity_index"`
	TotalWaterCost       *float64           `json:"total_water_cost"`
	TotalElectricityCost *float64           `json:"total_electricity_cost"`
	Status               *int32             `json:"status"`
}

type Contract struct {
	ID         int32   `json:"id"`
	RoomID     *int32  `json:"room_id"`
	SignatureA *string `json:"signature_a"`
	SignatureB *string `json:"signature_b"`
}

type ContractTemplate struct {
	ID                    int32              `json:"id"`
	PartyA                int32              `json:"party_a"`
	Address               []string           `json:"address"`
	ElectricityMethod     string             `json:"electricity_method"`
	ElectricityCost       float64            `json:"electricity_cost"`
	WaterMethod           string             `json:"water_method"`
	WaterCost             float64            `json:"water_cost"`
	InternetCost          float64            `json:"internet_cost"`
	ParkingFee            float64            `json:"parking_fee"`
	ResponsibilityA       string             `json:"responsibility_a"`
	ResponsibilityB       string             `json:"responsibility_b"`
	GeneralResponsibility string             `json:"general_responsibility"`
	CreatedAt             pgtype.Timestamptz `json:"created_at"`
	UpdatedAt             pgtype.Timestamptz `json:"updated_at"`
	DeletedAt             pgtype.Timestamptz `json:"deleted_at"`
}

type Conversation struct {
	ID            int32            `json:"id"`
	UserA         int32            `json:"user_a"`
	UserB         int32            `json:"user_b"`
	LastMessageID *int32           `json:"last_message_id"`
	CreatedAt     pgtype.Timestamp `json:"created_at"`
}

type Index struct {
	ID               int32   `json:"id"`
	WaterIndex       float64 `json:"water_index"`
	ElectricityIndex float64 `json:"electricity_index"`
	RoomID           int32   `json:"room_id"`
	Month            int32   `json:"month"`
	Year             int32   `json:"year"`
}

type LandlordRating struct {
	ID                    int32            `json:"id"`
	LandlordID            *int32           `json:"landlord_id"`
	RatedBy               *int32           `json:"rated_by"`
	FriendlinessRating    *int32           `json:"friendliness_rating"`
	ProfessionalismRating *int32           `json:"professionalism_rating"`
	SupportRating         *int32           `json:"support_rating"`
	TransparencyRating    *int32           `json:"transparency_rating"`
	OverallRating         *int32           `json:"overall_rating"`
	Comments              *string          `json:"comments"`
	CreatedAt             pgtype.Timestamp `json:"created_at"`
}

type Like struct {
	ID        int32              `json:"id"`
	RoomID    int32              `json:"room_id"`
	UserID    int32              `json:"user_id"`
	CreatedAt pgtype.Timestamptz `json:"created_at"`
	UpdatedAt pgtype.Timestamptz `json:"updated_at"`
	DeletedAt pgtype.Timestamptz `json:"deleted_at"`
}

type Message struct {
	ID              int32            `json:"id"`
	ConversationID  int32            `json:"conversation_id"`
	SenderID        int32            `json:"sender_id"`
	Type            int32            `json:"type"`
	Content         *string          `json:"content"`
	CreatedAt       pgtype.Timestamp `json:"created_at"`
	RentAutoContent []byte           `json:"rent_auto_content"`
}

type Notification struct {
	ID            int32            `json:"id"`
	UserID        int32            `json:"user_id"`
	ReferenceID   int32            `json:"reference_id"`
	ReferenceType string           `json:"reference_type"`
	Title         string           `json:"title"`
	IsRead        *bool            `json:"is_read"`
	CreatedAt     pgtype.Timestamp `json:"created_at"`
	UpdatedAt     pgtype.Timestamp `json:"updated_at"`
}

type Payment struct {
	ID              int32              `json:"id"`
	Code            string             `json:"code"`
	SenderID        int32              `json:"sender_id"`
	BillID          *int32             `json:"bill_id"`
	ContractID      *int32             `json:"contract_id"`
	Amount          float64            `json:"amount"`
	Status          int32              `json:"status"`
	ReturnRequestID *int32             `json:"return_request_id"`
	TransferContent *string            `json:"transfer_content"`
	EvidenceImage   *string            `json:"evidence_image"`
	PaidTime        pgtype.Timestamptz `json:"paid_time"`
}

type ProcessTracking struct {
	ID        int32              `json:"id"`
	Actor     int32              `json:"actor"`
	Action    string             `json:"action"`
	IssuedAt  pgtype.Timestamptz `json:"issued_at"`
	RequestID int32              `json:"request_id"`
}

type RentalRequest struct {
	ID              int32              `json:"id"`
	Code            string             `json:"code"`
	SenderID        int32              `json:"sender_id"`
	RoomID          int32              `json:"room_id"`
	SuggestedPrice  *float64           `json:"suggested_price"`
	NumOfPerson     *int32             `json:"num_of_person"`
	BeginDate       pgtype.Timestamptz `json:"begin_date"`
	EndDate         pgtype.Timestamptz `json:"end_date"`
	AdditionRequest *string            `json:"addition_request"`
	Status          int32              `json:"status"`
	CreatedAt       pgtype.Timestamptz `json:"created_at"`
	UpdatedAt       pgtype.Timestamptz `json:"updated_at"`
	DeletedAt       pgtype.Timestamptz `json:"deleted_at"`
}

type ReturnRequest struct {
	ID                 int32            `json:"id"`
	ContractID         *int32           `json:"contract_id"`
	Reason             *string          `json:"reason"`
	ReturnDate         pgtype.Timestamp `json:"return_date"`
	Status             *int32           `json:"status"`
	DeductAmount       *float64         `json:"deduct_amount"`
	TotalReturnDeposit *float64         `json:"total_return_deposit"`
	CreatedUser        *int32           `json:"created_user"`
	CreatedAt          pgtype.Timestamp `json:"created_at"`
	UpdatedAt          pgtype.Timestamp `json:"updated_at"`
	DeletedAt          pgtype.Timestamp `json:"deleted_at"`
}

type Room struct {
	ID              int32              `json:"id"`
	Title           string             `json:"title"`
	Address         []string           `json:"address"`
	RoomNumber      int32              `json:"room_number"`
	RoomImages      []string           `json:"room_images"`
	Utilities       []string           `json:"utilities"`
	Description     string             `json:"description"`
	RoomType        *string            `json:"room_type"`
	Owner           int32              `json:"owner"`
	Capacity        int32              `json:"capacity"`
	Gender          *int32             `json:"gender"`
	Area            float64            `json:"area"`
	TotalPrice      *float64           `json:"total_price"`
	Deposit         float64            `json:"deposit"`
	ElectricityCost float64            `json:"electricity_cost"`
	WaterCost       float64            `json:"water_cost"`
	InternetCost    float64            `json:"internet_cost"`
	IsParking       bool               `json:"is_parking"`
	ParkingFee      *float64           `json:"parking_fee"`
	Status          int32              `json:"status"`
	IsRent          bool               `json:"is_rent"`
	CreatedAt       pgtype.Timestamptz `json:"created_at"`
	UpdatedAt       pgtype.Timestamptz `json:"updated_at"`
	DeletedAt       pgtype.Timestamptz `json:"deleted_at"`
	AvailableFrom   pgtype.Timestamptz `json:"available_from"`
}

type RoomRating struct {
	ID                int32            `json:"id"`
	RoomID            *int32           `json:"room_id"`
	RatedBy           *int32           `json:"rated_by"`
	AmenitiesRating   *int32           `json:"amenities_rating"`
	LocationRating    *int32           `json:"location_rating"`
	CleanlinessRating *int32           `json:"cleanliness_rating"`
	PriceRating       *int32           `json:"price_rating"`
	OverallRating     *int32           `json:"overall_rating"`
	Comments          *string          `json:"comments"`
	Images            []string         `json:"images"`
	CreatedAt         pgtype.Timestamp `json:"created_at"`
}

type Tenant struct {
	ID        int32              `json:"id"`
	RoomID    int32              `json:"room_id"`
	TenantID  int32              `json:"tenant_id"`
	BeginDate pgtype.Timestamptz `json:"begin_date"`
	EndDate   pgtype.Timestamptz `json:"end_date"`
	CreatedAt pgtype.Timestamptz `json:"created_at"`
	UpdatedAt pgtype.Timestamptz `json:"updated_at"`
	DeletedAt pgtype.Timestamptz `json:"deleted_at"`
}

type TenantRating struct {
	ID                            int32            `json:"id"`
	TenantID                      *int32           `json:"tenant_id"`
	RatedBy                       *int32           `json:"rated_by"`
	PaymentRating                 *int32           `json:"payment_rating"`
	PropertyCareRating            *int32           `json:"property_care_rating"`
	NeighborhoodDisturbanceRating *int32           `json:"neighborhood_disturbance_rating"`
	ContractComplianceRating      *int32           `json:"contract_compliance_rating"`
	OverallRating                 *int32           `json:"overall_rating"`
	Comments                      *string          `json:"comments"`
	Images                        []string         `json:"images"`
	CreatedAt                     pgtype.Timestamp `json:"created_at"`
}

type Transaction struct {
	ID              int32   `json:"id"`
	PaymentID       int32   `json:"payment_id"`
	SenderID        int32   `json:"sender_id"`
	Amount          float64 `json:"amount"`
	Status          int32   `json:"status"`
	TransactionType *int32  `json:"transaction_type"`
	GatewayResponse *string `json:"gateway_response"`
}

type User struct {
	ID            int32              `json:"id"`
	PhoneNumber   string             `json:"phone_number"`
	FullName      string             `json:"full_name"`
	Password      string             `json:"password"`
	Address       *string            `json:"address"`
	Role          int32              `json:"role"`
	CreatedAt     pgtype.Timestamptz `json:"created_at"`
	DeletedAt     pgtype.Timestamptz `json:"deleted_at"`
	Otp           *int32             `json:"otp"`
	WalletAddress *string            `json:"wallet_address"`
	PrivateKeyHex *string            `json:"private_key_hex"`
	AvatarUrl     *string            `json:"avatar_url"`
	Gender        *int32             `json:"gender"`
	Dob           pgtype.Date        `json:"dob"`
}

type UserBank struct {
	UserID        int32            `json:"user_id"`
	BankID        int32            `json:"bank_id"`
	AccountNumber string           `json:"account_number"`
	AccountName   string           `json:"account_name"`
	CardNumber    *string          `json:"card_number"`
	Currency      *string          `json:"currency"`
	CreatedAt     pgtype.Timestamp `json:"created_at"`
	UpdatedAt     pgtype.Timestamp `json:"updated_at"`
}
