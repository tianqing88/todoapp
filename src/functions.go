package removefunc

import "fmt"

func removeText(fileName){
	input, err := ioutil.ReadFile(fileName)
	if err != nil {
			log.Fatalln(err)
	}

	lines := strings.Split(string(input), "\n")

	for i, line := range lines {

					lines[i] = ""
			}
	}
	output := strings.Join(lines, "\n")
	err = ioutil.WriteFile(fileName, []byte(output), 0644)
	if err != nil {
			log.Fatalln(err)
	}
}