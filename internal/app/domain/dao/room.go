package dao

type Room struct {
	BaseModel
	Title           string           `gorm:"title" json:"title"`
	Address         Address          `gorm:"polymorphic:Addressable;polymorphicValue:Room" json:"address"`
	Acreage         int              `gorm:"acreage" json:"acreage"`
	Price           float64          `gorm:"price" json:"price"`
	Description     string           `gorm:"description" json:"description"`
	DateSubmitted   string           `gorm:"date_submitted" json:"date_submitted"`
	OwnerID         uint             `gorm:"owner_id" json:"owner_id"`
	MaxPeople       int              `gorm:"max_people" json:"max_people"`
	RoomType        int              `gorm:"room_type" json:"room_type"`
	Deposit         float64          `gorm:"deposit" json:"deposit"`
	Services        []Service        `gorm:"many2many:room_services" json:"services"`
	Images          []string         `gorm:"images;serializer:json" json:"images"`
	BookingRequests []BookingRequest `gorm:"foreignKey:RoomID" json:"booking_requests"`
	BorrowedItems   []BorrowedItem   `gorm:"foreignKey:RoomID" json:"borrowed_items"`
}

type Image struct {
	FileName   string `json:"file_name"`
	FileBase64 string `json:"file_base64"`
}

type RoomRequest struct {
	Title         string         `json:"title"`
	Acreage       int            `json:"acreage"`
	Price         float64        `json:"price"`
	Description   string         `json:"description"`
	DateSubmitted string         `json:"date_submitted"`
	OwnerID       uint           `json:"owner_id"`
	MaxPeople     int            `json:"max_people"`
	RoomType      int            `json:"room_type"`
	Deposit       float64        `json:"deposit"`
	Services      []Service      `json:"services"`
	Images        []Image        `json:"images"`
	WardID        uint           `json:"ward_id"`
	AddressDetail string         `json:"address_detail"`
	BorrowedItems []BorrowedItem `json:"borrowed_items"`
}

type RoomResponse struct {
	ID            uint            `json:"id"`
	Title         string          `json:"title"`
	AddressID     uint            `json:"-"`
	Address       AddressResponse `json:"address"`
	Acreage       int             `json:"acreage"`
	Price         float64         `json:"price"`
	Description   string          `json:"description"`
	DateSubmitted string          `json:"date_submitted"`
	OwnerID       uint            `json:"owner_id"`
	MaxPeople     int             `json:"max_people"`
	RoomType      int             `json:"room_type"`
	Deposit       float64         `json:"deposit"`
	Services      []Service       `json:"services"`
	Images        []string        `json:"images"`
	BorrowedItems []BorrowedItem  `json:"borrowed_items"`
}

type RoomService struct {
	RoomID    uint `json:"room_id"`
	ServiceID uint `json:"service_id"`
}
