package api

// File represents a file ready to be downloaded.
// The file can be downloaded via the link https://api.telegram.org/file/bot<token>/<file_path>.
// It is guaranteed that the link will be valid for at least 1 hour.
// When the link expires, a new one can be requested by calling getFile.
// Maximum file size to download is 20 MB
type File struct {
    // ID is a unique file identifier
    ID string `json:"file_id"`

    // Size itself, if known
    // Optional
    Size int32 `json:"file_size"`

    // Path of the File
    // Use https://api.telegram.org/file/bot<token>/<path> to get the file.
    // Optional
    Path string `json:"file_path"`
}