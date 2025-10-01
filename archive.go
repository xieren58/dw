package main

// ArchiveJSON represents the structure of archive_files.json
type ArchiveJSON struct {
	Content   []ArchiveContent `json:"content"`
	TotalSize int              `json:"total_size"`
}

// ArchiveContent represents each item in the content array
type ArchiveContent struct {
	Name   string         `json:"name"`
	Size   int            `json:"size"`
	Pieces []ArchivePiece `json:"pieces"`
}

// ArchivePiece represents each item in the pieces array
type ArchivePiece struct {
	Name   string `json:"name"`
	Offset int    `json:"offset"`
}
