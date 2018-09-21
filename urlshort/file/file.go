package file

import "io/ioutil"

// Read reads a file and modifies the value if the file exists or returns an error
func Read(filepath string, original *[]byte) error {
	var err error

	if filepath != "" {
		*original, err = ioutil.ReadFile(filepath)
		if err != nil {
			return err
		}
	}

	return nil
}
