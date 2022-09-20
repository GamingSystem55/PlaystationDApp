import "os"

func openSingleFile() {
	file, err := os.Open("path/to/my/file.jpg")
	if err != nil {
		// do something with error
	}
	// pass file into web3.storage client...
}
