package api

import "gio-frontend-ms/pkg/model"

type ListRoomsPageData struct {
	Title string
	Rooms []*model.Room
}

type RoomPageData struct {
	Title string
	Room  *model.Room
}

type ListDevicesPageData struct {
	Title   string
	Devices []*model.Device
}

type DevicePageData struct {
	Title         string
	ApiGatewayUrl string
	Device        *model.Device
	Readings      []*model.Reading
	ErrorMessage  string
}
