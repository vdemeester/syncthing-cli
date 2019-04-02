package constants

const (
	AppName = "syncthing-cli"

	DefaultAddress = "127.0.0.1:8384"
	DynamicAddress = "dynamic"

	CompressionMetadataOnly = "metadata"
	CompressionAllData      = "always"
	CompressionOff          = "never"

	FolderTypeSendReceive = "sendreceive"
	FolderTypeSendOnly    = "sendonly"
	FolderTypeReceiveOnly = "receiveonly"

	FilePullOrderRandom        = "random"
	FilePullOrderAlphabetic    = "alphabetic"
	FilePullOrderSmallestFirst = "smallestFirst"
	FilePullOrderLargestFirst  = "largestFirst"
	FilePullOrderOldestFirst   = "oldestFirst"
	FilePullOrderNewestFirst   = "newestFirst"
)
