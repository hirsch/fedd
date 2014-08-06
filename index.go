package main

import (
	"log"
	"os"
	"path/filepath"
	"time"
)

var (
	folderCount uint64 = 0
	fileCount   uint64 = 0
)

type FileEntry struct {
	Name string // File name
}

type DirEntry struct {
	Files []*FileEntry // Files in the directory
	Path  string       // Absolute path of the directory
}

// index checks the initial parameters and induces crawling
func index(start string) []*DirEntry {
	file, err := os.Open(start)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	info, err := file.Stat()
	if err != nil {
		log.Fatal(err)
	}
	if !info.IsDir() {
		log.Fatal("index ", start+": no directory")
	}

	elapsedTime := time.Now()
	log.Println("Indexing started")
	data := crawl(start)
	log.Println("Indexing finished\nTime:", time.Since(elapsedTime), "folders:", folderCount, "files:", fileCount)
	return data
}

// crawl recursively crawls the filesystem
func crawl(start string) []*DirEntry {
	file, err := os.Open(start)
	if err != nil {
		log.Println(err)
		return nil
	}
	defer file.Close()

	cont, err := file.Readdir(-1)
	if err != nil {
		log.Println(err)
		return nil
	}

	thisdir := &DirEntry{make([]*FileEntry, 0), start}
	data := make([]*DirEntry, 0)
	for _, info := range cont {
		name := info.Name()
		if info.IsDir() {
			data = append(data, crawl(filepath.Join(start, name))...)
			folderCount++
		} else {
			entry := &FileEntry{name}
			thisdir.Files = append(thisdir.Files, entry)
			fileCount++
		}
	}
	return append(data, thisdir)
}
