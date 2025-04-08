package utils

import (
	"net/url"
)

func IsURL(path string) bool {
    u, err := url.Parse(path)
    if err != nil {
        return false
    }
    return u.Scheme == "http" || u.Scheme == "https"
}