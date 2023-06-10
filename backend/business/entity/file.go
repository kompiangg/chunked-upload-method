package entity

import "time"

type UploadFileMetadata struct {
	OriginName string `validate:"required,notblank" json:"origin_name"`
	Size       int    `validate:"required,notblank" json:"size"`

	ChunkSize  int    `json:"chunk_size" swaggerignore:"true"`
	ChunkCount int    `json:"chunk_count" swaggerignore:"true"`
	UniqueName string `json:"unique_name" swaggerignore:"true"`
}

type UploadBinaryFile struct {
	IdentityName string
	Order        int
	ByteCodeData []byte
}

type GetFileMetadata struct {
	IdentityName string `param:"identity_name"`
}

type GetFileMetadataResponse struct {
	OriginName string `json:"origin_name"`
	Size       int    `json:"size"`
	ChunkSize  int    `json:"chunk_size"`
	ChunkCount int    `json:"chunk_count"`
	UniqueName string `json:"unique_name"`
}

type UploadFileMetaDataResponse struct {
	IdentityName  string `json:"unique_name"`
	ChunkByteSize int    `json:"chunk_byte_size"`
	ChunkCount    int    `json:"chunk_count"`
}

type AssembleByteCode struct {
	UniqueName   string         `json:"unique_name"`
	FileByteCode []FileByteCode `json:"-"`
}

type FileByteCode struct {
	ByteCode []byte
	Order    int
}

type File struct {
	ID        int       `json:"id" db:"id"`
	Name      string    `json:"name" db:"name"`
	URL       string    `json:"url" db:"url"`
	CreatedAt time.Time `json:"created_at" db:"created_at"`
}
