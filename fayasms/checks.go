package fayasms

import (
	"fmt"
)

// mandatoryFields are required by FayaSMS to be present in every request
var mandatoryFields = []string{
	"AppKey", "AppSecret",
}

// checkMandatoryFields checks to ensure that the mandatory fields are set
func (f *FayaSMS) checkMandatoryFields(mandatoryFields []string) error {
	for _, field := range mandatoryFields {
		if f.payload.Get(field) == "" {
			return fmt.Errorf("fayasms: a mandatory field has not been set. please supply all mandatory fields which are: %v", mandatoryFields)
		}
	}

	return nil
}


// contingentFields are only required based on the endpoint being hit.
// This map shows the endpoints and the fields they require
var contingentFields = map[string][]map[string]string{
	"send": {
		{"name": "From", "errMsg": "no sender id has been set"},
		{"name": "Message", "errMsg": "no message body has been set"},
		{"name": "To", "errMsg": "no recipient has been set"},
	},
	"estimate": {
		{"name": "Recipients", "errMsg": "no recipient has been set"},
		{"name": "Message", "errMsg": "no message body set"},
	},
}

// checkContingentFields checks that all contingent fields required by endpoint are set
func (f *FayaSMS) checkContingentFields(endpoint string, contingentFields map[string][]map[string]string) error {
	fields, ok := contingentFields[endpoint]

	// Some endpoints do not have any contingent fields
	if !ok {
		return nil
	}

	for _, field := range fields {
		if f.payload.Get(field["name"]) == "" {
			return fmt.Errorf("fayasms: %v", field["errMsg"])
		}
	}

	return nil
}

