package log4go

import (
	"sync"
)

var (
	defaultSize = 10
)

type fieldsPool struct {
	// contains filtered or unexported fields
	cacheMu          sync.Mutex
	stringSliceCache map[int][][2]string // size: [][2]string
}

func newFieldsPool() *fieldsPool {
	return &fieldsPool{
		stringSliceCache: make(map[int][][2]string),
	}
}

func (p *fieldsPool) GetFields(size int) [][2]string {
	if size == 0 {
		size = defaultSize
	}

	if size%10 != 0 {
		size = (size/10 + 1) * 10
	}

	p.cacheMu.Lock()
	cache, ok := p.stringSliceCache[size]
	p.cacheMu.Unlock()

	if !ok {
		cache = make([][2]string, 0, size)
	}
	return cache
}

func (p *fieldsPool) ExtendFields(fields [][2]string, newSize int) [][2]string {
	currSize := cap(fields)
	if newSize < currSize {
		return fields
	}

	newCache := p.GetFields(newSize)
	copy(newCache, fields)
	p.ReturnFields(fields)

	return newCache
}

func (p *fieldsPool) ReturnFields(fields [][2]string) {
	p.cacheMu.Lock()
	p.stringSliceCache[cap(fields)] = fields
	p.cacheMu.Unlock()
}
