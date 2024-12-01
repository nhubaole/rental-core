package responses

type RoomOnChainRes struct {
	ID     int64
	TotalPrice int
	Deposit    int64
	Status     int64
	IsRent     bool
	CreatedAt int64
	UpdatedAt int64
}

type MContractOnChainRes struct {
	ID                   int64  // ID của hợp đồng
	Code                 string // Mã hợp đồng
	Landlord             int64  // ID của chủ nhà
	Tenant               int64  // ID của người thuê
	RoomID               int64  // ID của phòng
	ActualPrice          int64  // Giá thực tế của hợp đồng
	Deposit              int64  // Tiền đặt cọc
	BeginDate            int64  // Thời gian bắt đầu hợp đồng (timestamp)
	EndDate              int64  // Thời gian kết thúc hợp đồng (timestamp)
	PaymentMethod        string // Phương thức thanh toán
	ElectricityMethod    string // Phương thức tính điện
	ElectricityCost      int64  // Giá điện
	WaterMethod          string // Phương thức tính nước
	WaterCost            int64  // Giá nước
	InternetCost         int64  // Giá internet
	ParkingFee           int64  // Phí gửi xe
	ResponsibilityA      string // Trách nhiệm bên A
	ResponsibilityB      string // Trách nhiệm bên B
	GeneralResponsibility string // Trách nhiệm chung
	SignatureA           string // Chữ ký của bên A
	SignedTimeA          int64  // Thời gian ký của bên A (timestamp)
	SignatureB           string // Chữ ký của bên B
	SignedTimeB          int64  // Thời gian ký của bên B (timestamp)
	ContractTemplateID   int64  // ID mẫu hợp đồng
	PreRentalStatus      uint8  // Trạng thái trước khi thuê
	RentalProcessStatus  uint8  // Trạng thái trong quá trình thuê
	PostRentalStatus     uint8  // Trạng thái sau khi thuê
	CreatedAt            int64  // Thời gian tạo hợp đồng (timestamp)
	UpdatedAt            int64  // Thời gian cập nhật hợp đồng (timestamp)
}
