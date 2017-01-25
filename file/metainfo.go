package file

type MetaInfo struct {
	Info         InfoDict   `bencode:"info"`
	Announce     string     `bencode:"announce"`
	AnnounceList [][]string `bencode:"announce-list"`
	CreationDate int        `bencode:"creation date"`
	Comment      string     `bencode:"comment"`
	CreatedBy    string     `bencode:"created by"`
	Encoding     string     `bencode:"encoding"`
}

// InfoDict is the structure that contains all the InfoDict
// about the pieces of the torrent.
type InfoDict struct {
	PieceLength int    `bencode:"piece length"`
	Pieces      string `bencode:"pieces"`
	Private     int    `bencode:"private"`

	Name string `bencode:"name"`

	// Single File Mode
	Length int    `bencode:"length"`
	MD5Sum string `bencode:"md5sum"`

	// Multiple File Mode
	Files []ArtifactDict `bencode:"files"`
}

// ArtifactDict is a structure to hold file information in a multi-file torrent.
type ArtifactDict struct {
	Length int      `bencode:"length"`
	MD5Sum string   `bencode:"md5"`
	Path   []string `bencode:"path"`
}
