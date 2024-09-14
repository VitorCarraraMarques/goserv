package video

import (
	"encoding/binary"
	"fmt"
	"log"
	"os"
)

type AtomHeader struct {
    Pos int64
	Size uint32
	Type string
}

func main() {
	if len(os.Args) != 2 {
		log.Fatal("filename missing")
	}
	f := OpenMP4(os.Args[1])
	stat, err := f.Stat()
    if err != nil {
        log.Fatal(err)
    }
    filesize := stat.Size()

    var rootatoms []*AtomHeader
    var moovatoms []*AtomHeader
    var trakatoms []*AtomHeader


	var offset int64 = 0
    rootatoms = ReadAtoms(f, offset, filesize)
    for _, pheader := range rootatoms {
        header := (*pheader)
        fmt.Printf("header[ Type: %s | Size: %d | Pos: %d ] \n", header.Type, header.Size, header.Pos)
        if header.Type == "moov" && header.Size > 8 {
            offset = header.Pos + 8
            endatom := header.Pos + int64(header.Size)
            moovatoms = ReadAtoms(f, offset, endatom)
        }
    }
    for _, subheader := range moovatoms {
        sub := (*subheader)
        fmt.Printf("___subheader[ Type: %s | Size: %d | Pos: %d ] \n", sub.Type, sub.Size, sub.Pos)
        if sub.Type == "trak"{
            offset = sub.Pos + 8
            end := sub.Pos + int64(sub.Size)
            trakatoms = ReadAtoms(f, offset, end)
        }
    }
    for _, ptrak := range trakatoms {
        trak := (*ptrak)
        fmt.Printf("_______trak[ Type: %s | Size: %d | Pos: %d ] \n", trak.Type, trak.Size, trak.Pos)
    }
}

func OpenMP4(name string) *os.File {
	f, err := os.Open(name)
	if err != nil {
		log.Fatal(err)
	}
	return f
}

func ReadAtoms(f *os.File, offset int64, maxsize int64) []*AtomHeader {
    var buffer []*AtomHeader
    for {
        header := ReadAtomHeader(f, offset)
        buffer = append(buffer, header)
        offset += int64(header.Size)
        if offset >= maxsize {
            break
        }
    }
    return buffer 
}

func ReadAtomHeader(f *os.File, offset int64) *AtomHeader {
	buffer := make([]byte, 8)
	_, err := f.ReadAt(buffer, offset)
	if err != nil {
		log.Fatal(err)
	}
	header := AtomHeader{
        Pos: offset,
		Size: binary.BigEndian.Uint32(buffer[:4]),
		Type: string(buffer[4:]),
	}
	return &header
}

