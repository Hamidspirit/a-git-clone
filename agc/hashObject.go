package agc

import (
	"bufio"
	"crypto/sha1"
	"encoding/hex"
	"io"
	"log"
	"os"
	"path/filepath"

	"github.com/Hamidspirit/a-git-clone/util"
)

func HashObject(path, objectType string, filesname string) (fpath string, objectid string) {
	// File path
	var fp string

	if path == "." {
		// Construct file path
		fp = util.FilePathParser(path, filesname)

		// check if file exists
		if _, err := os.Stat(fp); os.IsNotExist(err) {
			log.Fatalf("File does not exist: %s", fp)
		}
	} else {
		fp = path
	}

	// Read and hash the file
	file, err := os.Open(fp)
	if err != nil {
		log.Fatal("Failed to open file:", err)
	}
	defer file.Close()

	// Add object file to repo
	reader := bufio.NewReader(file)
	oid, err := SaveHashedObject(reader, objectType)
	if err != nil {
		log.Fatal("Failed to add object file to repo")
	}

	return fp, oid
}

func SaveHashedObject(reader *bufio.Reader, objectType string) (string, error) {
	hasher := sha1.New()

	// Temp file to store data as i hash it
	tmpFile, err := os.CreateTemp("", "object_temp_*")
	if err != nil {
		log.Fatal("Failed to create temp file", err)
	}
	defer os.Remove(tmpFile.Name()) // in case there was an error

	// Add the type of object and after add a null byte
	tmpFile.WriteString(objectType + "\x00")

	buffer := make([]byte, 4096) // 4KB chunks of data

	for {
		// Why use a loop with reader.Read (buffer)?
		// This pattern allows reading the input stream in fixed-size chunks (e.g., 4KB).
		// It's memory-efficient and suitable for large files, as it avoids loading the entire file into memory.
		n, err := reader.Read(buffer)
		if n > 0 {

			// Why check n > 0 before processing the chunk?
			// n > 0 ensures we only process actual data read into the buffer.
			// reader.Read() can return 0 bytes without being EOF, especially in networked or slow streams.
			// We continue reading until we reach EOF (err == io.EOF), so we don't miss anything even if the file > 4KB.

			chunk := buffer[:n]
			hasher.Write(chunk)
			if _, writeErr := tmpFile.Write(chunk); writeErr != nil {
				tmpFile.Close()
				return "", writeErr
			}
			// How does chunked hashing work?
			// The SHA-1 hasher is incremental – it maintains internal state across multiple Write() calls.
			// Writing data in 4KB chunks produces the same final hash as writing the entire content at once.
			// This lets us compute the hash as we read, without storing the full file in memory.
		}

		if err == io.EOF {
			break
		}

		if err != nil {
			tmpFile.Close()
			return "", err
		}
	}

	// Get object id sha1 hash
	oid := hex.EncodeToString(hasher.Sum(nil))

	// Final object path
	objectPath := filepath.Join(GitRepo, "objects", oid)

	// Rewind temp file and copy to final destination
	_, err = tmpFile.Seek(0, 0)
	if err != nil {
		tmpFile.Close()
		return "", err
	}

	// Create output file
	outputFile, err := os.Create(objectPath)
	if err != nil {
		return "", err
	}
	defer outputFile.Close()

	// Copy temp file to destination
	if _, err := io.Copy(outputFile, tmpFile); err != nil {
		tmpFile.Close()
		return "", err
	}

	tmpFile.Close()
	return oid, nil
}
