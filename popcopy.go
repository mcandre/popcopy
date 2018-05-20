package popcopy

import (
	"io"
	"io/ioutil"
	"log"
	"os"
	"path/filepath"
	"regexp"
)

// Copy recursively copies the named directory or file path
// to the target directory path,
// excluding any specified path patterns.
func Copy(source string, target string, exclusions []*regexp.Regexp) error {
	sourceAbs, err := filepath.Abs(source)

	if err != nil {
		return err
	}

	for _, exclusion := range exclusions {
		if exclusion.MatchString(sourceAbs) {
			return nil
		}
	}

	sourceFI, err := os.Stat(sourceAbs)

	if err != nil {
		return err
	}

	targetAbs, err := filepath.Abs(target)

	if err != nil {
		return err
	}

	sourceBase := filepath.Base(source)
	targetAbsWithSourceBase := filepath.Join(targetAbs, sourceBase)

	sourceMode := sourceFI.Mode()

	switch {
	case sourceMode.IsDir():
		if err := os.MkdirAll(targetAbsWithSourceBase, os.ModeDir|0775); err != nil {
			return err
		}

		childrenFI, err := ioutil.ReadDir(sourceAbs)

		if err != nil {
			return err
		}

		for _, childFI := range childrenFI {
			child := childFI.Name()

			if err := Copy(filepath.Join(source, child), targetAbsWithSourceBase, exclusions); err != nil {
				return err
			}
		}
	default:
		fIn, err := os.Open(sourceAbs)
		defer func() {
			if err = fIn.Close(); err != nil {
				log.Panic(err)
			}
		}()

		if err != nil {
			return err
		}

		fOut, err := os.Create(targetAbsWithSourceBase)
		defer func() {
			if err = fOut.Close(); err != nil {
				log.Panic(err)
			}
		}()

		if err != nil {
			return err
		}

		if _, err := io.Copy(fOut, fIn); err != nil {
			return err
		}
	}

	return nil
}
