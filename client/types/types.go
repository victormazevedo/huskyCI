// Copyright 2019 Globo.com authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package types

import (
	"time"

	"gopkg.in/mgo.v2/bson"
)

// FoundVuln is the boolean that will be checked to return an os.exit(0) or os.exit(1)
var FoundVuln bool

// FoundInfo is the boolean that will be checked to verify if only low/info severity vulnerabilites were found.
var FoundInfo bool

// JSONPayload is a struct that represents the JSON payload needed to make a HuskyCI API request.
type JSONPayload struct {
	RepositoryURL    string `json:"repositoryURL"`
	RepositoryBranch string `json:"repositoryBranch"`
}

// Analysis is the struct that stores all data from analysis performed.
type Analysis struct {
	ID            bson.ObjectId  `bson:"_id,omitempty"`
	RID           string         `bson:"RID" json:"RID"`
	URL           string         `bson:"repositoryURL" json:"repositoryURL"`
	Branch        string         `bson:"repositoryBranch" json:"repositoryBranch"`
	SecurityTests []SecurityTest `bson:"securityTests" json:"securityTests"`
	Status        string         `bson:"status" json:"status"`
	Result        string         `bson:"result" json:"result"`
	Containers    []Container    `bson:"containers" json:"containers"`
}

// Container is the struct that stores all data from a container run.
type Container struct {
	CID          string       `bson:"CID" json:"CID"`
	SecurityTest SecurityTest `bson:"securityTest" json:"securityTest"`
	CStatus      string       `bson:"cStatus" json:"cStatus"`
	COutput      string       `bson:"cOutput" json:"cOutput"`
	CResult      string       `bson:"cResult" json:"cResult"`
	CInfo        string       `bson:"cInfo" json:"cInfo"`
	StartedAt    time.Time    `bson:"startedAt" json:"startedAt"`
	FinishedAt   time.Time    `bson:"finishedAt" json:"finishedAt"`
}

// SecurityTest is the struct that stores all data from the security tests to be executed.
type SecurityTest struct {
	ID               bson.ObjectId `bson:"_id,omitempty"`
	Name             string        `bson:"name" json:"name"`
	Image            string        `bson:"image" json:"image"`
	Cmd              string        `bson:"cmd" json:"cmd"`
	Language         string        `bson:"language" json:"language"`
	Default          bool          `bson:"default" json:"default"`
	TimeOutInSeconds int           `bson:"timeOutSeconds" json:"timeOutSeconds"`
}

// HuskyCIVulnerability is the struct that stores vulnerability information.
type HuskyCIVulnerability struct {
	SecurityTool   string `json:"securitytool,omitempty"`
	Severity       string `json:"severity,omitempty"`
	Confidence     string `json:"confidence,omitempty"`
	File           string `json:"file,omitempty"`
	Line           string `json:"line,omitempty"`
	Code           string `json:"code,omitempty"`
	Details        string `json:"details,omitempty"`
	Type           string `json:"type,omitempty"`
	VunerableBelow string `json:"vulnerablebelow,omitempty"`
	Version        string `json:"version,omitempty"`
}

// JSONOutput is a truct that represents huskyCI output in a JSON format.
type JSONOutput struct {
	GoResults         GoResults         `json:"goresults,omitempty"`
	PythonResults     PythonResults     `json:"pythonresults,omitempty"`
	JavaScriptResults JavaScriptResults `json:"javascriptresults,omitempty"`
	RubyResults       RubyResults       `json:"rubyresults,omitempty"`
}

// GoResults represents all Golang security tests results.
type GoResults struct {
	GosecOutput []HuskyCIVulnerability `json:"gosecoutput,omitempty"`
}

// PythonResults represents all Python security tests results.
type PythonResults struct {
	BanditOutput []HuskyCIVulnerability `json:"banditoutput,omitempty"`
	SafetyOutput []HuskyCIVulnerability `json:"safetyoutput,omitempty"`
}

// JavaScriptResults represents all JavaScript security tests results.
type JavaScriptResults struct {
	RetirejsResult []HuskyCIVulnerability `json:"retirejsoutput,omitempty"`
}

// RubyResults represents all Ruby security tests results.
type RubyResults struct {
	BrakemanOutput []HuskyCIVulnerability `json:"brakemanoutput,omitempty"`
}
