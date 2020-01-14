package util

import "github.com/CeruleanSong/gobox-server/src/model"

// ErrorFILENOTFOUND as
var ErrorFILENOTFOUND = &model.ErrorResponce{
	MESSAGE: "File not Found",
}

// ErrorDATABASE as
var ErrorDATABASE = &model.ErrorResponce{
	MESSAGE: "Error Document Not in Database",
}

// ErrorUPLOADERROR as
var ErrorUPLOADERROR = &model.ErrorResponce{
	MESSAGE: "Error During Upload",
}

// ErrorINTERNALCRASH as
var ErrorINTERNALCRASH = &model.ErrorResponce{
	MESSAGE: "Internal Server Crash",
}
