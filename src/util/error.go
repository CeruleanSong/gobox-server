package util

import "github.com/CeruleanSong/gobox-server/src/model"

// ErrorFILENOTFOUND as
var ErrorFILENOTFOUND = &model.ErrorResponce{
	STATUS:  404,
	MESSAGE: "File not Found",
}

// ErrorDATABASE as
var ErrorDATABASE = &model.ErrorResponce{
	STATUS:  404,
	MESSAGE: "Error Document Not in Database",
}

// ErrorUPLOADERROR as
var ErrorUPLOADERROR = &model.ErrorResponce{
	STATUS:  500,
	MESSAGE: "Error During Upload",
}

// ErrorINTERNALCRASH as
var ErrorINTERNALCRASH = &model.ErrorResponce{
	STATUS:  500,
	MESSAGE: "Internal Server Crash",
}
