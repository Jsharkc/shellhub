// Package geoip helps in geolocation operations.
package geoip

import (
	"fmt"
	"io"
	"net"
	"net/http"
	"os"

	"github.com/sirupsen/logrus"

	"github.com/mholt/archiver/v3"

	"github.com/oschwald/geoip2-golang"
)

// DBPath is the default path for Database.
var DBPath = "/usr/share/GeoIP/"

// GeoLite2DbName is the default database name of GeoLite2 when extracted.
var GeoLite2DbName = "GeoLite2-City.mmdb"

// Check if geoLite2 implements Locator interface.
var _ Locator = (*geoLite2)(nil)

// Check if geoLite2 implements io.Closer interface.
var _ io.Closer = (*geoLite2)(nil)

// geoLite2 is a structure what stores a geoIp2Reader to a GeoIp2 database.
type geoLite2 struct {
	db *geoip2.Reader
}

// downloadGeoLite2Db downloads the GeoLite2 database and extract the files into the DBPath.
func downloadGeoLite2Db(maxmindDBLicense string) error {
	// Download the GeoLite2Db .tar.gz file with the database inside it.
	r, err := http.Get(fmt.Sprintf("https://download.maxmind.com/app/geoip_download?edition_id=GeoLite2-City&license_key=%s&suffix=tar.gz", maxmindDBLicense))
	if err != nil {
		return err
	}

	temp, err := os.CreateTemp("/tmp", "geoip*.tar.gz")
	defer func(temp *os.File) {
		err = temp.Close()
		if err != nil {
			logrus.Error("Could not close the:", temp.Name())
		}
	}(temp)
	if err != nil {
		return err
	}

	_, err = io.Copy(temp, r.Body)
	if err != nil {
		return err
	}

	if _, err := os.Stat(DBPath); os.IsNotExist(err) {
		// TODO Permission
		err := os.Mkdir(DBPath, 0o755)
		if err != nil {
			return err
		}
	}

	g, err := os.Create(DBPath + GeoLite2DbName)
	defer func(g *os.File) {
		err := g.Close()
		if err != nil {
			logrus.Error("Cloud not close the:", g.Name())
		}
	}(g)
	if err != nil {
		return err
	}

	err = archiver.Walk(temp.Name(), func(f archiver.File) error {
		// This could be problematic in large files, but it's okay for GeoLite2 case
		if f.Name() == GeoLite2DbName {
			_, err := io.Copy(g, f)
			if err != nil {
				return err
			}
		}

		return nil
	})

	return nil
}

// NewGeoLite2 opens a connection to GeoIp2 database and return a geoLite2 structure with the database connection.
//
// The connection uses the local database or try to download it from MaxMind's server (the download required `MAXMIND_LICENSE`).
func NewGeoLite2() (Locator, error) {
	if _, err := os.Stat(DBPath + GeoLite2DbName); os.IsNotExist(err) {
		if os.Getenv("MAXMIND_LICENSE") != "" {
			err := downloadGeoLite2Db(os.Getenv("MAXMIND_LICENSE"))
			if err != nil {
				return nil, err
			}
		} else {
			return nil, err
		}
	}

	db, err := geoip2.Open(DBPath + GeoLite2DbName)
	if err != nil {
		return nil, err
	}

	return &geoLite2{
		db: db,
	}, nil
}

// Close the connection with the GeoLite2 database, returning either error or nil.
func (g *geoLite2) Close() error {
	return g.db.Close()
}

// GetCountry gets an ip and return either an ISO 3166-1 code to a country or an empty string.
func (g *geoLite2) GetCountry(ip net.IP) (string, error) {
	record, err := g.db.Country(ip)
	if err != nil {
		return "", err
	}

	return record.Country.IsoCode, nil
}

// GetPosition gets an ip and return a Position structure with Longitude and Latitude with error nil or an empty Position structure with the error.
func (g *geoLite2) GetPosition(ip net.IP) (Position, error) {
	record, err := g.db.City(ip)
	if err != nil {
		return Position{}, err
	}

	return Position{
		Longitude: record.Location.Longitude,
		Latitude:  record.Location.Latitude,
	}, nil
}
