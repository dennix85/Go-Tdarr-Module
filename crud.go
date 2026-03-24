package tdarr

import "encoding/json"

// GetFileByID is a convenience wrapper around Cruddb for the common case of
// fetching a single file record from the FileJSONDB collection by its ID.
func (c *Client) GetFileByID(id string) (json.RawMessage, error) {
	return c.Cruddb(CruddbRequest{
		Collection: CollectionFile,
		Mode:       CrudModeGetByID,
		DocID:      id,
	})
}

// Crud is a convenience alias for Cruddb with a shorter name.
// Use the Collection* and CrudMode* constants to build the request.
//
// Example — fetch a file by ID:
//
//	raw, err := client.Crud(tdarr.CruddbRequest{
//	    Collection: tdarr.CollectionFile,
//	    Mode:       tdarr.CrudModeGetByID,
//	    DocID:      fileID,
//	})
func (c *Client) Crud(req CruddbRequest) (json.RawMessage, error) {
	return c.Cruddb(req)
}
