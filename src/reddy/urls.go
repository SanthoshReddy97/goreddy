package reddy

import (
	"github.com/SanthoshReddy97/goreddy/internal/httpReddy"
)

func RegisterUrls(path, method string, view httpReddy.HandlerFunc) {
	httpReddy.URLs = append(httpReddy.URLs,
		httpReddy.RouteInfo{Path: path, Method: method, HandlerFunc: view})
}
