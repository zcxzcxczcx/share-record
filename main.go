package main

import "sync"

type (
	Record struct {
		Val interface{}
	}
	ShareRecord struct {
		M       sync.Mutex
		Records map[string]*Record
	}
)

func main() {

}

func NewShareRecord() *ShareRecord {
	return &ShareRecord{
		Records: map[string]*Record{},
	}
}
func (sr *ShareRecord) Do(key string, fn func() (interface{}, error)) (interface{}, error) {
	sr.M.Lock()
	if record, ok := sr.Records[key]; ok {
		sr.M.Unlock()
		return record.Val, nil
	}
	val, err := fn()
	if err != nil {
		return nil, err
	}
	sr.Records[key].Val = val
	sr.M.Unlock()
	return sr.Records[key].Val, nil

}
