package config

type UploadFileConfig struct {
	ChunkSizeChar string `validate:"required,notblank,numeric,gt=0"`
	ChunkSize     int
}
