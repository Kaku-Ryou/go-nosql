package supported_db

import (
	"couch-go-master"
	"db/entity"
	"fmt"
)

// CouchDb - Struct for couch database.
type CouchDb struct {
	Conn couch.Database
}

// Save - Save generic record in couchDB.
func (this CouchDb) Save(record interface{}) bool {
	id, rev, err := this.Conn.Insert(record)
	if err == nil && id != "" && rev != "" {
		return true
	} else {
		return false
	}
}

// Read - Read all records from couchDB.
func (this CouchDb) Read() []entity.Map {
	ids := getIds(this)
	records := make([]entity.Map, len(ids))
	for i := 0; i < len(ids); i++ {
		err := this.Conn.RetrieveFast(ids[i], &records[i])
		if err != nil {
			panic(err)
		}
	}
	return records
}

// Delete - Delete generic record in couchDB.
func (this CouchDb) Delete(record map[string]interface{}) bool {
	rev, err := this.Conn.Retrieve(fmt.Sprintf("%s", record["_id"]), &record)
	if err == nil && rev != "" {
		err = this.Conn.Delete(fmt.Sprintf("%s", record["_id"]), rev)
		if err == nil {
			return true
		} else {
			return false
		}
	} else {
		return false
	}
}

// Update - Update record in couchDB.
func (this CouchDb) Update(record map[string]interface{}) bool {
	var r map[string]interface{}
	rev, err := this.Conn.Retrieve(fmt.Sprintf("%s", record["_id"]), &r)
	if err == nil {
		rev, err = this.Conn.EditWith(record, fmt.Sprintf("%s", record["_id"]), rev)
		if err == nil {
			return true
		} else {
			return false
		}
	} else {
		panic(err)
	}
}

// First - Read first record from couchDB
func (this CouchDb) First() entity.Map {
        ids := getIds(this)
        record := make(entity.Map)
        err := this.Conn.RetrieveFast(ids[0], &record)
        if err != nil {
                       panic(err)
                      }
        return record
}

// Last - Read last record from couchDB
func (this CouchDb) Last() entity.Map {
        ids := getIds(this)
        record := make(entity.Map)
        err := this.Conn.RetrieveFast(ids[len(ids)-1], &record)
        if err != nil {
                       panic(err)
                      }
        return record
}

// Count - Read number of records from couchDB
func (this CouchDb) Count() int {
        ids, err := this.Conn.QueryIds("_all_docs", nil)
	if err != nil {
		return 0
	}
        return len(ids)
}

// Limit - Read limited number of records from couchDB.
func (this CouchDb) Limit(limit int) []entity.Map {
	if limit <= 0 {
		return make([]entity.Map,0)
	}
        ids := getIds(this)
	if limit > len(ids) {
		limit = len(ids)
	}
        records := make([]entity.Map, limit)
        for i := 0; i < limit; i++ {
                err := this.Conn.RetrieveFast(ids[i], &records[i])
                if err != nil {
                             panic(err)
                            }
                }
        return records
}


// FindById - Read record by id from couchDB.
func (this CouchDb) FindById(id string) entity.Map {
	var record entity.Map
        err := this.Conn.RetrieveFast(id, &record)
        if err != nil {
                       panic(err)
                      }
        return record
}

// getIds - Read all document ids from couchDB.
func getIds(this CouchDb) []string {
	ids, err := this.Conn.QueryIds("_all_docs", nil)
	if err != nil {
		return make([]string,0)
	}
	return ids
}
