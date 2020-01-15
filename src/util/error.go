package util

import "github.com/CeruleanSong/gobox-server/src/model"

// ErrorINVALIDAUTH as
var ErrorINVALIDAUTH = &model.ErrorResponce{
	STATUS:  500,
	MESSAGE: "Invalid Username and/or Password",
}

// ErrorDUPLICATE as
var ErrorDUPLICATE = &model.ErrorResponce{
	STATUS:  409,
	MESSAGE: "Duplicate user account",
}

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
