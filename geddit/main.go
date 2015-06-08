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

Copyright 2015 Marleny Nunez
- Command line arguments addition
- Nicer output formatting
*/

package main

import (
	"fmt"
	"log"
	"os"
	"github.com/marlenunez/reddit"
)

func main() {

	var subreddit string
	if len(os.Args) < 2 {
		fmt.Println("Usage: geddit <subreddit>\nNo subreddit provided, defaulting to golang...")
		subreddit = "golang"
	} else {
		subreddit = os.Args[1]
	}

	// fetch the subreddit
	items, err := reddit.Get(subreddit)
	if err != nil {
		log.Fatal(err)
	}

	for _, item := range items {
		fmt.Println(item)
	}
}
