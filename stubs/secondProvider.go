package stubs

func SecondProvider(prompt string) (int, string, error) {
	var data map[string]string = map[string]string{
		"What is the capital of India":              "New Delhi",
		"What is the capital of France":             "Paris",
		"Where is olympic games being held in 2024": "Paris",
		"Where was olympic games held in 2021":      "Tokyo",
	}

	if _, ok := data[prompt]; ok {
		return 200, data[prompt], nil
	}

	return 200, "", nil
}
