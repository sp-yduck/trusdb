package main

import (
	"fmt"
	"io"
	"log"
	"os"
)

type FileManager interface {
	ReadPage(PageID, []byte) error
	WritePage(PageID, []byte) error
	AllocatePage() PageID
	DeallocatePage(PageID)
	GetNumWrites() uint64
	ShutDown()
	Size() int64
}

//File is the disk implementation of FileManager
type File struct {
	db         *os.File
	fileName   string
	nextPageID PageID
	numWrites  uint64
	size       int64
}

// type BlockID struct {
// }

func NewFileManager(filename string) FileManager {
	file, err := os.OpenFile(filename, os.O_RDWR|os.O_CREATE, 0666)
	if err != nil {
		log.Fatalln("can't open db file")
		return nil
	}

	fileInfo, err := file.Stat()
	if err != nil {
		log.Fatalln("file info error")
		return nil
	}

	fileSize := fileInfo.Size()
	nPages := fileSize / PageSize

	nextPageID := PageID(0)
	if nPages > 0 {
		nextPageID = PageID(int32(nPages + 1))
	}

	return &File{file, filename, nextPageID, 0, fileSize}
}

// ShutDown closes of the database file
func (f *File) ShutDown() {
	f.db.Close()
}

// Write a page to the database file
func (f *File) WritePage(pageId PageID, pageData []byte) error {
	offset := int64(pageId * PageSize)
	f.db.Seek(offset, io.SeekStart)
	bytesWritten, err := f.db.Write(pageData)
	if err != nil {
		return err
	}

	if bytesWritten != PageSize {
		return fmt.Errorf("bytes written not equals page size")
	}

	if offset >= f.size {
		f.size = offset + int64(bytesWritten)
	}

	f.db.Sync()
	return nil
}

// Read a page from the database file
func (f *File) ReadPage(pageID PageID, pageData []byte) error {
	offset := int64(pageID * PageSize)

	fileInfo, err := f.db.Stat()
	if err != nil {
		return fmt.Errorf("file info error")
	}

	if offset > fileInfo.Size() {
		return fmt.Errorf("I/O error past end of file")
	}

	f.db.Seek(offset, io.SeekStart)

	bytesRead, err := f.db.Read(pageData)
	if err != nil {
		return fmt.Errorf("I/O error while reading")
	}

	if bytesRead < PageSize {
		for i := 0; i < PageSize; i++ {
			pageData[i] = 0
		}
	}
	return nil
}

//  AllocatePage allocates a new page
//  For now just keep an increasing counter
func (f *File) AllocatePage() PageID {
	ret := f.nextPageID
	f.nextPageID++
	return ret
}

// DeallocatePage deallocates page
// Need bitmap in header page for tracking pages
// This does not actually need to do anything for now.
func (f *File) DeallocatePage(pageID PageID) {
}

// GetNumWrites returns the number of disk writes
func (f *File) GetNumWrites() uint64 {
	return f.numWrites
}

// Size returns the size of the file in disk
func (f *File) Size() int64 {
	return f.size
}

// func (fm *FileManager) Write(blockId *BlockID, page *Page) {

// }
