/*
	*** Washizator ***
	NOTE: Credits go to the authors of the audio source.
 */

/* Settings:
 * Opening a browser with a URL from the command-line
 * Preferred way to open a browser with a URL from the command-line
 * macOS: 'open' <URL>
 * Windows: 'rundll32' <URL>
 * Linux: 'xdg-open' <URL>
 */

package main


import (
	"fmt"
	"os/exec"
	"runtime"
)

const (
	URL 	= "https://www.youtube.com/watch?v=TwB4uINHe9Y&t=60s&autoplay=1"
	GREEN   = "\033[32m"
	ITALICS = "\033[3m"
	RESET 	= "\033[0m"
)

// Create the main function
func main() {
	openBrowserWithUrl(URL)
	fmt.Printf("%s%s'Washizator' - Vašo P. was here...%s\n", GREEN, ITALICS, RESET)
}

// Create a function that opens a browser with a URL
func openBrowserWithUrl(URL string) {
	// Determine the OS
	userOS := runtime.GOOS

	// Buffer to handle errors
	var ERR error

	// Open the browser based on the OS
	switch userOS {
	case "linux":
		ERR = exec.Command("xdg-open", URL).Start()
	case "windows":
		ERR = exec.Command("rundll32", "url.dll,FileProtocolHandler", URL).Start()
	case "darwin":
		ERR = exec.Command("open", URL).Start()
	default:
		ERR = fmt.Errorf("unsupported platform: %s", userOS)
	}
	// Handle possible Errors
	if ERR != nil {
		panic(ERR)
	}
}

