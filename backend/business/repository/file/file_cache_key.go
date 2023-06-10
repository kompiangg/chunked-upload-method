package file

import "time"

const (
	FileMetadata                                 = `fp:file:metadata:%s`
	FileBinary                                   = `fp:file:binary:%s:%d`
	PrefixFileBinary                             = `fp:file:binary:%s:*`
	DurationFileMetadataExpiration time.Duration = 10 * time.Minute
)
