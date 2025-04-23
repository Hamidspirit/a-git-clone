package util

import (
	"bufio"
	"crypto/sha1"
	"encoding/hex"
	"io"
	"log"
	"os"
	"path/filepath"

	"github.com/Hamidspirit/a-git-clone/agc"
)

func HashObject(path *string, args []string) (string, []byte) {
	// File path
	var fp string

	if *path == "" {
		// Check if third argument args[3] exists
		if len(args) >= 3 {
			// Get current working directory
			wd, err := os.Getwd()
			if err != nil {
				log.Fatal("Failed to get current directory:", err)
			}

			// Construct file path
			fp = filepath.Join(wd, args[2])

			// check if file exists
			if _, err := os.Stat(fp); os.IsNotExist(err) {
				log.Fatalf("File does not exist: %s", fp)
			}
		} else {
			log.Fatal("No path file specified with -p and no file name provided")
		}
	} else {
		fp = *path
	}

	// Read and hash the file
	file, err := os.Open(fp)
	if err != nil {
		log.Fatal("Failed to open file:", err)
	}
	defer file.Close()

	// Add object file to repo
	reader := bufio.NewReader(file)
	_, err = SaveHashedObject(reader)
	if err != nil {
		log.Fatal("Failed to add object file to repo")
	}

	// Create a new sha1 hash
	hasher := sha1.New()

	// Copy the contents of reader to hasher
	if _, err := io.Copy(hasher, reader); err != nil {
		log.Fatal("Failed to read and hash file\n", err)
	}

	return fp, hasher.Sum(nil)
}

func SaveHashedObject(reader *bufio.Reader) (string, error) {
	hasher := sha1.New()

	// Temp file to store data as i hash it
	tmpFile, err := os.CreateTemp("", "object_temp_*")
	if err != nil {
		log.Fatal("Failed to create temp file", err)
	}
	defer os.Remove(tmpFile.Name()) // in case there was an error

	buffer := make([]byte, 4096) // 4KB chunks of data

	for {
		n, err := reader.Read(buffer)
		if n > 0 {
			chunk := buffer[:n]
			hasher.Write(chunk)
			if _, writeErr := tmpFile.Write(chunk); writeErr != nil {
				tmpFile.Close()
				return "", writeErr
			}
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
	objectPath := filepath.Join(agc.GitRepo, "objects", oid)

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
