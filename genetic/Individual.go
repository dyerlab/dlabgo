package genetic

import (
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/sqlite"
)

// Individual struct is the main one that has the geospatial data associated with it.
type Individual struct {
	gorm.Model
}
