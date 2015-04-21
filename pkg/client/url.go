package client

import "net/url"

// Type - enum of different url types
type Type int

// enum types
const (
	Unknown    Type = iota // Unknown type
	Object                 // Minio and S3 compatible object storage
	Filesystem             // POSIX compatible file systems
)

// GetType returns the type of .
func GetType(urlStr string) Type {
	u, err := url.Parse(urlStr)
	if err != nil {
		return Unknown
	}

	if u.Scheme == "http" || u.Scheme == "https" {
		return Object
	}

	if u.Scheme == "" {
		return Filesystem
	}

	return Unknown
}
