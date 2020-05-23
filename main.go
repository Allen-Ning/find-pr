package main

import (
	"fmt"
	"log"
	"os/exec"
	"runtime"
	"strings"

	"github.com/go-git/go-git/v5"
	"github.com/go-git/go-git/v5/plumbing"
	"github.com/go-git/go-git/v5/plumbing/object"
)

// test := false
func main() {
	fmt.Println("looks cool")
	r, _ := git.PlainOpen("/Users/allen/Desktop/api.activepipe.com")
	ref, _ := r.Head()
	// ... retrieves the commit history
	cIter, _ := r.Log(&git.LogOptions{From: ref.Hash()})

	value := "da2f5cd7d4cf5658182f28688ea5b07552b45390"
	// test := false
	// ... just iterates over the commits, printing it
	_ = cIter.ForEach(func(c *object.Commit) error {
		if c.NumParents() > 1 {
			if hasValue(c.ParentHashes, value) {
				fmt.Println("____________________")
				fmt.Println(c.Message)
				fmt.Println(c.Hash)
				words := strings.Fields(c.Message)
				pr_id := words[3][1:]
				fmt.Println(pr_id)
				open_browser("https://github.com/ActivePipe/api.activepipe.com/pull/" + pr_id)
			}
		}
		return nil
	})
}

func hasValue(s []plumbing.Hash, hash string) bool {
	for _, value := range s {
		if value.String() == hash {
			return true
		}
	}
	return false
}

func open_browser(url string) {
	var err error

	switch runtime.GOOS {
	case "linux":
		err = exec.Command("xdg-open", url).Start()
	case "windows":
		err = exec.Command("rundll32", "url.dll,FileProtocolHandler", url).Start()
	case "darwin":
		err = exec.Command("open", url).Start()
	default:
		err = fmt.Errorf("unsupported platform")
	}
	if err != nil {
		log.Fatal(err)
	}

}
