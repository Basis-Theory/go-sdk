// Code generated by Fern. DO NOT EDIT.

package accountupdater

type JobsListRequest struct {
	// The maximum number of jobs to return
	Size *int `json:"-" url:"size,omitempty"`
	// Cursor for pagination
	Start *string `json:"-" url:"start,omitempty"`
}
