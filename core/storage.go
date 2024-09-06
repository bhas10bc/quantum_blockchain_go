package core

import (
	"encoding/json"
	"fmt"

	"github.com/fjl/go-couchdb"
)

// Storage interface represents the storage layer for the blockchain.
type Storage interface {
	Put(*Block) error
}

// CouchDBStore represents a CouchDB-based storage implementation.
type CouchDBStore struct {
	connection *couchdb.DB
}

// NewCouchDBStore initializes a new CouchDB storage.
func NewCouchDBStore(couchDBURL string, dbName string, user string, pass string) (*CouchDBStore, error) {
	// Include the username and password in the CouchDB URL
	authURL := fmt.Sprintf("http://%s:%s@%s", user, pass, couchDBURL)

	// Create a new CouchDB client
	client, err := couchdb.NewClient(authURL, nil)
	if err != nil {
		return nil, fmt.Errorf("failed to create CouchDB client: %w", err)
	}

	// Create or open the database
	db := client.DB(dbName)

	// Check if the database exists by attempting to get its info
	
	if err != nil {
		if couchErr, ok := err.(*couchdb.Error); ok && couchErr.StatusCode == 404 {
			// Database does not exist, create it
			_, err := client.CreateDB(dbName)
			if err != nil {
				return nil, fmt.Errorf("failed to create database: %w", err)
			}
		} else {
			return nil, fmt.Errorf("failed to connect to CouchDB database: %w", err)
		}
	}

	return &CouchDBStore{connection: db}, nil
}

// Put inserts a block into the CouchDB database.
func (s *CouchDBStore) Put(b *Block) error {
	// Ensure the block hash is available for use as a document ID.
	blockHash :=  (b.DataHash).String()
	

	// Serialize the block data to JSON
	blockData, err := json.Marshal(b)
	if err != nil {
		return fmt.Errorf("failed to serialize block data: %w", err)
	}

	// Define the block data dynamically
	blockMap := map[string]interface{}{
		"_id": blockHash,
		"data": string(blockData),
	}

	// Print the JSON string for debugging purposes.
	fmt.Println("Block data JSON:", blockMap)

	// Store the block data as JSON in CouchDB.
	_, err = s.connection.Put(blockHash, blockMap, "")
	if err != nil {
		return fmt.Errorf("failed to store block in CouchDB: %w", err)
	}


	fmt.Println("Block saved successfully")
	return nil
}
