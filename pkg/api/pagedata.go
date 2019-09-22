package api

import "gio-frontend-ms/pkg/model"

// Data structure for ListRooms page
type ListRoomsPageData struct {
	Title string
	Rooms []*model.Room
}

// Data structure for Room detail page
type RoomPageData struct {
	Title string
	Room  *model.Room
}

// Data structure for ListDevices page
type ListDevicesPageData struct {
	Title   string
	Devices []model.Device
}

// Data structure for Device detail page
type DevicePageData struct {
	Title        string
	Device       *model.Device
	Readings     []model.Reading
	ErrorMessage string
}
