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
	junkName := "Thumbs.db"
	junk := []byte{}
	forecastName := "forecast.txt"
	forecast := []byte("30% growth month over month\n")
	imagesName := "images"
	imageName := "space.bmp"
	image := []byte{}
	gitDirectory := ".git"
	gitConfigName := "config"
	gitConfig := []byte("[core]\neditor = vi")
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

	if err = ioutil.WriteFile(filepath.Join(sourceAbsPath, forecastName), forecast, 0666); err != nil {
		t.Error(err)
	}

	if err = ioutil.WriteFile(filepath.Join(sourceAbsPath, junkName), junk, 0666); err != nil {
		t.Error(err)
	}

	if err = os.MkdirAll(filepath.Join(sourceAbsPath, imagesName), os.ModeDir|0775); err != nil {
		t.Error(err)
	}

	if err = ioutil.WriteFile(filepath.Join(sourceAbsPath, imagesName, imageName), image, 0666); err != nil {
		t.Error(err)
	}

	if err = os.MkdirAll(filepath.Join(sourceAbsPath, gitDirectory), os.ModeDir|0775); err != nil {
		t.Error(err)
	}

	if err = ioutil.WriteFile(filepath.Join(sourceAbsPath, gitDirectory, gitConfigName), gitConfig, 0666); err != nil {
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

	observedForecast, err := ioutil.ReadFile(filepath.Join(targetAbsPath, forecastName))

	if err != nil {
		t.Error(err)
	}

	if !reflect.DeepEqual(observedForecast, forecast) {
		t.Errorf("Expected forecast %v to equal %v", observedForecast, forecast)
	}

	observedImage, err := ioutil.ReadFile(filepath.Join(targetAbsPath, imagesName, imageName))

	if err != nil {
		t.Error(err)
	}

	if !reflect.DeepEqual(observedImage, image) {
		t.Errorf("Expected image %v to equal %v", observedImage, image)
	}

	targetWithJunk := filepath.Join(targetAbsPath, junkName)

	if _, err := os.Stat(targetWithJunk); !os.IsNotExist(err) {
		t.Errorf("Expected %v to not exist", targetWithJunk)
	}

	targetWithGitConfig := filepath.Join(targetAbsPath, gitDirectory, gitConfigName)

	if _, err := os.Stat(targetWithGitConfig); !os.IsNotExist(err) {
		t.Errorf("Expected %v to not exist", targetWithGitConfig)
	}
}
