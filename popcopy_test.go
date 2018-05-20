package popcopy_test

import (
	"io/ioutil"
	"os"
	"path/filepath"
	"reflect"
	"regexp"
	"testing"

	"github.com/mcandre/popcopy"
)

func TestCopy(t *testing.T) {
	sourcePrefix := "business-presentations"
	forecastName := "forecast.txt"
	forecast := []byte("30% growth month over month\n")
	junkName := "Thumbs.db"
	junk := []byte{}
	targetPrefix := "usbkey"
	exclusions := []*regexp.Regexp{
		regexp.MustCompile("Thumbs.db"),
	}

	sourceAbsPath, err := ioutil.TempDir("", sourcePrefix)
	defer func() {
		if err = os.RemoveAll(sourceAbsPath); err != nil {
			t.Error(err)
		}
	}()

	if err != nil {
		t.Error(err)
	}

	sourceBasePath := filepath.Base(sourceAbsPath)

	if err = ioutil.WriteFile(filepath.Join(sourceAbsPath, forecastName), forecast, 0666); err != nil {
		t.Error(err)
	}

	if err = ioutil.WriteFile(filepath.Join(sourceAbsPath, junkName), junk, 0666); err != nil {
		t.Error(err)
	}

	targetAbsPath, err := ioutil.TempDir("", targetPrefix)
	defer func() {
		if err = os.RemoveAll(targetAbsPath); err != nil {
			t.Error(err)
		}
	}()

	if err != nil {
		t.Error(err)
	}

	err = popcopy.Copy(sourceAbsPath, targetAbsPath, exclusions)

	if err != nil {
		t.Error(err)
	}

	observedForecast, err := ioutil.ReadFile(filepath.Join(targetAbsPath, sourceBasePath, forecastName))

	if err != nil {
		t.Error(err)
	}

	if !reflect.DeepEqual(observedForecast, forecast) {
		t.Errorf("Expected forecast %v to equal %v", observedForecast, forecast)
	}

	targetWithJunk := filepath.Join(targetAbsPath, sourceBasePath, junkName)

	if _, err := os.Stat(targetWithJunk); !os.IsNotExist(err) {
		t.Errorf("Expected %v to not exist", targetWithJunk)
	}
}
