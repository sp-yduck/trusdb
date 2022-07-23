package main

import (
	"bytes"
	"encoding/binary"
)

// PageSize is the size of a page in disk (4KB)
const PageSize = 4096

// InvalidPageID represents an invalid page ID
const InvalidPageID = PageID(-1)

// PageID is the type of the page identifier
type PageID int32

// Page represents an abstract page on disk
type Page struct {
	id       PageID          // idenfies the page. It is used to find the offset of the page on disk
	pinCount uint32          // counts how many goroutines are acessing it
	isDirty  bool            // the page was modified but not flushed
	data     *[PageSize]byte // bytes stored in disk
}

// IsValid checks if id is valid
func (id PageID) IsValid() bool {
	return id != InvalidPageID || id >= 0
}

// Serialize casts it to []byte
func (id PageID) Serialize() []byte {
	buf := new(bytes.Buffer)
	binary.Write(buf, binary.LittleEndian, id)
	return buf.Bytes()
}

// NewPageIDFromBytes creates a page id from []byte
func NewPageIDFromBytes(data []byte) (ret PageID) {
	binary.Read(bytes.NewBuffer(data), binary.LittleEndian, &ret)
	return ret
}

// IncPinCount decrements pin count
func (p *Page) IncPinCount() {
	p.pinCount++
}

// DecPinCount decrements pin count
func (p *Page) DecPinCount() {
	if p.pinCount > 0 {
		p.pinCount--
	}
}

// PinCount retunds the pin count
func (p *Page) PinCount() uint32 {
	return p.pinCount
}

// ID retunds the page id
func (p *Page) ID() PageID {
	return p.id
}

// Data returns the data of the page
func (p *Page) Data() *[PageSize]byte {
	return p.data
}

// SetIsDirty sets the isDirty bit
func (p *Page) SetIsDirty(isDirty bool) {
	p.isDirty = isDirty
}

// IsDirty check if the page is dirty
func (p *Page) IsDirty() bool {
	return p.isDirty
}

// Copy copies data to the page's data. It is mainly used for testing
func (p *Page) Copy(offset uint32, data []byte) {
	copy(p.data[offset:], data)
}

// New creates a new page
func New(id PageID, isDirty bool, data *[PageSize]byte) *Page {
	return &Page{id, uint32(1), isDirty, data}
}

// NewEmpty creates a new empty page
func NewEmpty(id PageID) *Page {
	return &Page{id, uint32(1), false, &[PageSize]byte{}}
}
