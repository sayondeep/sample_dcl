package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

var (
	vendorID                     = 4891                      //imp
	productID                    = 32769                     //imp
	softwareVersions             = []int{10080000, 105, 106} //imp
	softwareVersion              = 106                       //imp
	softwareVersionString        = "106.0"                   //imp
	cdVersionNumber              = 1                         //imp
	firmwareInformation          = ""
	softwareVersionValid         = true                                                                //imp
	otaUrl                       = "https://sayon-rm.s3.amazonaws.com/apple-light-matter-ota_v106.bin" //imp
	otaFileSize                  = "1430977"
	otaChecksum                  = "l98JrB7hsstBBagr0VPudPzyBqZptDOpVZPMXzEgPj8="
	otaChecksumType              = 1
	minApplicableSoftwareVersion = 0    //imp
	maxApplicableSoftwareVersion = 1000 //imp
	releaseNotesUrl              = ""
	creator                      = "cosmos1d8vdghcynsl8r7ram2ref3zz6ez6n54f4aat76"
	schemaVersion                = 0
)

type queryModelVersions struct {
	Vid              int   `json:"vid"`
	Pid              int   `json:"pid"`
	SoftwareVersions []int `json:"softwareVersions"`
	SchemaVersion    int   `json:"schemaVersion"`
}

type otaCandidateModelVersion struct {
	Vid                          int    `json:"vid"`
	Pid                          int    `json:"pid"`
	SoftwareVersion              int    `json:"softwareVersion"`
	SoftwareVersionString        string `json:"softwareVersionString"`
	CdVersionNumber              int    `json:"cdVersionNumber"`
	FirmwareInformation          string `json:"firmwareInformation"`
	SoftwareVersionValid         bool   `json:"softwareVersionValid"`
	OtaUrl                       string `json:"otaUrl"`
	OtaFileSize                  string `json:"otaFileSize"`
	OtaChecksum                  string `json:"otaChecksum"`
	OtaChecksumType              int    `json:"otaChecksumType"`
	MinApplicableSoftwareVersion int    `json:"minApplicableSoftwareVersion"`
	MaxApplicableSoftwareVersion int    `json:"maxApplicableSoftwareVersion"`
	ReleaseNotesUrl              string `json:"releaseNotesUrl"`
	Creator                      string `json:"creator"`
	SchemaVersion                int    `json:"schemaVersion"`
}

type querySwVersionResponse struct {
	ModelVersion queryModelVersions `json:"modelVersions"`
}

type queryOtaCandidateResponse struct {
	ModelVersion otaCandidateModelVersion `json:"modelVersion"`
}

func querySoftwareVersion(w http.ResponseWriter, r *http.Request) {
	item := querySwVersionResponse{
		ModelVersion: queryModelVersions{
			Vid:              vendorID,
			Pid:              productID,
			SoftwareVersions: softwareVersions,
			SchemaVersion:    0,
		},
	}

	// Marshal struct to JSON
	jsonData, err := json.Marshal(item)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	// Set content type to application/json
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)

	// Write the raw JSON data directly
	w.Write(jsonData)
}

func queryOtaCandidate(w http.ResponseWriter, r *http.Request) {
	// Construct the response struct dynamically
	item := queryOtaCandidateResponse{
		ModelVersion: otaCandidateModelVersion{
			Vid:                          vendorID,
			Pid:                          productID,
			SoftwareVersion:              softwareVersion,
			SoftwareVersionString:        softwareVersionString,
			CdVersionNumber:              cdVersionNumber,
			FirmwareInformation:          firmwareInformation,
			SoftwareVersionValid:         softwareVersionValid,
			OtaUrl:                       otaUrl,
			OtaFileSize:                  otaFileSize,
			OtaChecksum:                  otaChecksum,
			OtaChecksumType:              otaChecksumType,
			MinApplicableSoftwareVersion: minApplicableSoftwareVersion,
			MaxApplicableSoftwareVersion: maxApplicableSoftwareVersion,
			ReleaseNotesUrl:              releaseNotesUrl,
			Creator:                      creator,
			SchemaVersion:                schemaVersion,
		},
	}

	// Marshal struct to JSON
	jsonData, err := json.Marshal(item)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	// Set content type to application/json
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)

	// Write the raw JSON data directly
	w.Write(jsonData)
}

func main() {
	router := mux.NewRouter()

	// Construct the route dynamically
	queryimageroutePath := fmt.Sprintf("/dcl/model/versions/%d/%d", vendorID, productID)
	router.HandleFunc(queryimageroutePath, querySoftwareVersion).Methods("GET")

	queryotacandidateroutePath := fmt.Sprintf("/dcl/model/versions/%d/%d/%d", vendorID, productID, softwareVersion)
	router.HandleFunc(queryotacandidateroutePath, queryOtaCandidate).Methods("GET")

	fs := http.FileServer(http.Dir("/ota_firmware/"))
	router.PathPrefix("/ota_firmware/").Handler(http.StripPrefix("/ota_firmware/", fs))

	log.Fatal(http.ListenAndServe(":8000", router))
}
