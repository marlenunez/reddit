/*
Copyright 2013 Google Inc.
Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at
	http://www.apache.org/licenses/LICENSE-2.0
Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

package main

// import statement
import (
	"fmt"
	"log"
	"os"
	"github.com/marlenunez/reddit"
)

// function declaration
func main() {

	if len(os.Args) != 1 {
		fmt.Printf("Usage : %s subreddit\nDefaulting to golang...", os.Args[0])
		sub = "golang"
	} else {
		sub := os.Args[1]
	}

	// fetch the sub reddit
	items, err := reddit.Get(sub)

	// error handing
	if err != nil {
		log.Fatal(err)
	}

	for _, item := range items {
		// print the title of the items
		fmt.Println(item)
	}
}
