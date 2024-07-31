package main

import (
	"fmt"
	"math/rand"
	"syscall/js"
	"time"
)

func main() {
	randSource := rand.New(rand.NewSource(time.Now().UnixNano()))
	rand.New(randSource)

	c := make(chan struct{}, 0)

	js.Global().Set("executeCommand", js.FuncOf(executeCommand))

	<-c
}

func executeCommand(this js.Value, p []js.Value) interface{} {
	command := p[0].String()
	var result string

	switch command {
	case "1":
		result = `Commands: 
  1 - About me 
  2 - My skills 
  3 - My projects
  4 - Work experience
  5 - Contact info
  6 - Clear terminal
  7 - Current date and time
  8 - System information
  9 - Random joke`
	case "2":
		result = "I'm Kavindu Harshana, a full stack developer. I work with Go lang in backend, React and React Native in front end."
	case "3":
		result = `Skills:
 - Backend: GO, Express,
 - Frontend: React (Next.js, Vite, Astro), React Native (Expo, Metro bundler).
 - Database: MySQL, SQLite, MongoDB, Postgres
 - Other: Git, RESTful APIs`
	case "4":
		result = `
Work Experience:
1. Senior Software Developer
   Extra Co Group (Sep 2023 - Present)

2. Software Engineer
   Dhabione (Feb 2023 - Sep 2023)

3. Associate Software Engineer
   Lead Lanka International (Jul 2022 - Feb 2023)

4. Wordpress Developer
   Cyber Concepts Sri Lanka (Aug 2021 - Jun 2022)

5. Trainee Web Developer
   SATASME HOLDINGS PVT LTD (Jan 2020 - Dec 2021)`
	case "5":
		result = `
Contact:
  Email: kythonlk@gmail.com
  GitHub: github.com/kythonlk`
	case "6":
		result = ""
	case "7":
		result = fmt.Sprintf("Current date and time: %s", time.Now().Format("2006-01-02 15:04:05"))
	case "8":
		result = getSystemInfo()
	case "9":
		result = getRandomJoke()
	default:
		result = fmt.Sprintf("Unknown command: %s. Type '1' for commands.", command)
	}
	fmt.Println("Returning result:", result)
	return js.ValueOf(result)
}

func getSystemInfo() string {
	return `

 _   __      _   _                 _ _                
| | / /     | | | |               | | |               
| |/ / _   _| |_| |__   ___  _ __ | | | __   ___  ___ 
|    \| | | | __| '_ \ / _ \| '_ \| | |/ /  / _ \/ __|
| |\  \ |_| | |_| | | | (_) | | | | |   <  | (_) \__ \
\_| \_/\__, |\__|_| |_|\___/|_| |_|_|_|\_\  \___/|___/
        __/ |                                         
       |___/                                          
                                                                                          
 
OS:       Kythonlk OS
Kernel:   Kythonlk posix
Architecture: WASM
Host:     kythonlkOS
`
}

func getRandomJoke() string {
	jokes := []string{
		"Why don't scientists trust atoms? Because they make up everything!",
		"I told my wife she was drawing her eyebrows too high. She looked surprised.",
		"Why don't programmers like nature? It has too many bugs.",
		"What do you call fake spaghetti? An impasta!",
		"Why do cows have hooves instead of feet? Because they lactose.",
		"Why did the scarecrow win an award? Because he was outstanding in his field!",
		"I would tell you a joke about an elevator, but it's an uplifting experience.",
		"Why did the math book look sad? Because it had too many problems.",
	}

	return jokes[rand.Intn(len(jokes))]
}
