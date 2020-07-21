package clouddriver

import (
	"fmt"
	"path/filepath"

	"github.com/cucumber/godog"
	"github.com/cucumber/godog/colors"

	"citihub.com/probr/internal/coreengine"
	"citihub.com/probr/test/features"
)

//this is the "TEST HANDLER" impl  and will get called when probr is invoked from the CLI or API
//all we do here is set the godog args based on what has been supplied (e.g. output path)
//and call to the "feature" implementation (i.e the same impl when godog / go test is invoked)

//Init ...
func init() {
	n, c := "account_manager", coreengine.General
	td := coreengine.TestDescriptor{Category: c, Name: n}

	coreengine.TestHandleFunc(td, TH)
}

//TH ...
func TH() (int, error) {
	r, err := features.GetRootDir()

	if err != nil {
		return -1, fmt.Errorf("unable to determine root directory - not able to perform tests")
	}

	var t = "clouddriver"
	featPath := filepath.Join(r, "test", "features", "clouddriver", "features")

	f, err := features.GetOutputPath(&t)
	if err != nil {
		return -2, err
	}

	opts := godog.Options{
		Format: "cucumber",
		Output: colors.Colored(f),
		Paths:  []string{featPath},
	}

	status := godog.TestSuite{
		Name:                 "account_manager",
		TestSuiteInitializer: TestSuiteInitialize,
		ScenarioInitializer:  ScenarioInitialize,
		Options:              &opts,
	}.Run()

	return status, nil
}
