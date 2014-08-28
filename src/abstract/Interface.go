package abstract

import "db/entity"

// Database - Interface to define basic DB functions.
type Database interface {
	Save(interface{}) bool
	Read() []entity.Map
	Update(interface{}) bool
	Delete(interface{}) bool
        First() entity.Map
        Last() entity.Map
	Count() int
	Limit(int) []entity.Map
	FindById(string) entity.Map
	Where(string) []entity.Map
}
