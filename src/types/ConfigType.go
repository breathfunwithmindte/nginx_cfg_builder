package types

import (
	"encoding/json"
	"fmt"
)

type Config struct {
	Name    string    `json:"name"`
	Version string    `json:"version"`
	Output  string    `json:"output"`
	Items   []Element `json:"items"`
}

func (c Config) PrettyPrint() {
	blue := "\033[34m"
	reset := "\033[0m"
	underline := "\033[4m"
	green := "\033[32m"
	red := "\033[31m"
	yellow := "\033[33m"

	fmt.Printf("%s%sConfig%s\n", blue, underline, reset)
	fmt.Printf("%sName%s: %s\n", red, reset, c.Name)
	fmt.Printf("%sOutput%s: %s\n", green, reset, c.Output)
	fmt.Printf("%sVersion%s: %s\n", yellow, reset, c.Version)
	fmt.Printf("%sElements%s:\n", green, reset)
	for _, item := range c.Items {
		item.PrettyPrint()
	}
	fmt.Println("--------------------------------------------------------------------- ENDL ---------------------------------------------------------------------")
	fmt.Print("\n\n")
}

func (c *Config) UnmarshalJSON(data []byte) error {
	type Alias Config
	temp := &struct {
		Items []json.RawMessage `json:"items"`
		*Alias
	}{
		Alias: (*Alias)(c),
	}
	if err := json.Unmarshal(data, &temp); err != nil {
		return err
	}

	for _, raw := range temp.Items {
		var kind struct {
			Kind string `json:"kind"`
		}
		if err := json.Unmarshal(raw, &kind); err != nil {
			return err
		}

		switch kind.Kind {
		case "server":
			var s Server
			if err := json.Unmarshal(raw, &s); err != nil {
				return err
			}
			s.Init()
			c.Items = append(c.Items, s)
		case "comment":
			var com Comment
			if err := json.Unmarshal(raw, &com); err != nil {
				return err
			}
			c.Items = append(c.Items, com)
		// ... other cases for different kinds as needed
		default:
			return fmt.Errorf("unknown kind: %s in file: %s", kind.Kind, c.Name)
		}
	}

	return nil
}
