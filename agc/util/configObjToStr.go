package util

import "fmt"

func ConfigObjToStr(config map[string]map[string]map[string]bool) string {
	// Start with a formatted string
	configStr := ""

	// Iterate over the nested map and build the string
	for section, subsection := range config {
		for key, value := range subsection {
			for _, val := range value {
				// Format it to look like `core.bare=true` in the string
				configStr += fmt.Sprintf("%s.%s=%t\n", section, key, val)
			}
		}
	}
	return configStr
}
