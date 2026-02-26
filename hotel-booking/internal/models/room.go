package models

type RoomType string

const (
	RoomStandard RoomType = "standard"
	RoomDeluxe   RoomType = "deluxe"
	RoomVIP      RoomType = "vip"
)

type Room struct {
	ID          int
	HotelID     int
	Type        RoomType
	Name        string
	Price       int
	Beds        int
	Capacity    int
	SizeM2      int
	Facilities  []string
	FloorPlan   string
	Images      []string
	MapEmbedURL string
}
