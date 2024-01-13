package main

import (
	"encoding/json"
	"os"
)

type Pref struct {
	DarkMode bool `json:"darkMode"`
}

var defaultPref = Pref{DarkMode: false}

func writePreference(fileName string, pref Pref) error {

	f, err := os.Create(fileName)
	if err != nil {
		return err
	}
	defer f.Close()

	encoder := json.NewEncoder(f)
	if err := encoder.Encode(pref); err != nil {
		return err
	}

	return nil
}

func updatePreference(fileName string, pref Pref) error {
	// update the pref file

	if _, err := os.Stat(fileName); os.IsNotExist(err) {

		writePreference(fileName, pref)
	} else {
		f, err := os.Open(fileName)
		if err != nil {
			return err
		}
		defer f.Close()
		encoder := json.NewEncoder(f)
		if err := encoder.Encode(pref); err != nil {
			return err
		}
	}
	return nil

}

func readPreference(fileName string) (*Pref, error) {
	var pref Pref

	if _, err := os.Stat(fileName); os.IsNotExist(err) {
		// File does not exist, call writePreference with a default value
		pref = defaultPref

		if err := writePreference(fileName, pref); err != nil {
			return nil, err
		}

	} else {

		f, err := os.Open(fileName)
		if err != nil {
			return nil, err
		}
		defer f.Close()

		decoder := json.NewDecoder(f)
		if err := decoder.Decode(&pref); err != nil {
			return nil, err
		}
	}

	return &pref, nil
}
