package message

import (
	"fmt"
	"log"
	"strings"
	"sync"
)

type (
	// DictionaryInfo is an item for Dictionary
	DictionaryInfo struct {
		// Fn is handler name in Servcie
		Fn interface{}
		// Code is route comporessed code
		Code uint16
	}
	// Dictionary is alias for []*DictionaryInfo
	Dictionary = []*DictionaryInfo
)

var (
	// Routes is a map from route to code
	Routes = make(map[string]uint16)
	// Codes is a map from code to route
	Codes = make(map[uint16]string)
	rw    sync.RWMutex
)

// ReadDictionary returns dictionary for compressed route.
func ReadDictionary() (map[string]uint16, map[uint16]string) {
	rw.RLock()
	defer rw.RUnlock()

	return Routes, Codes
}

// WriteDictionary is to set dictionary when new route dictionray is found.
func WriteDictionary(dict map[string]uint16) {
	routes, codes := TransformDictionary(dict)

	rw.Lock()
	defer rw.Unlock()

	for k, v := range routes {
		Routes[k] = v
	}
	for k, v := range codes {
		Codes[k] = v
	}
}

// TransformDictionary transfroms user defined dict into routes and codes
func TransformDictionary(dict map[string]uint16) (map[string]uint16, map[uint16]string) {
	routes := make(map[string]uint16)
	codes := make(map[uint16]string)

	for route, code := range dict {
		r := strings.TrimSpace(route)

		// duplication check
		if _, ok := routes[r]; ok {
			log.Println(fmt.Sprintf("duplicated route(route: %s, code: %d)", r, code))
		}

		if _, ok := codes[code]; ok {
			log.Println(fmt.Sprintf("duplicated route(route: %s, code: %d)", r, code))
		}

		// update map, using last value when key duplicated
		routes[r] = code
		codes[code] = r
	}

	return routes, codes
}
