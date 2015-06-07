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

Make an HTTP request

This program makes an HTTP request to the Reddit API
and copies its response to standard output.
Put this in a file named main.go inside your reddit directory.

Screencast: https://www.youtube.com/watch?v=2KmHtgtEZ1s
Slides:     https://talks.golang.org/2012/tutorial.slide#1
*/

// Package reddit implements a basic client for the Reddit API.
package reddit

// import statement
import (
    "encoding/json"
    "errors"
    "fmt"
    "net/http"
)

type response struct {
    Data struct {
        Children []struct {
            Data Item
        }
    }
}

// Item describes a Reddit item.
type Item struct {
	Title string
	URL string
	Comments int `json:"num_comments"`
}

func (i Item) String() string {
    com := ""
    switch i.Comments {
    case 0:
        // nothing
    case 1:
        com = " (1 comment)"
    default:
        com = fmt.Sprintf(" (%d comments)", i.Comments)
    }
    return fmt.Sprintf("%s\n%s\t%s\n", i.Title, com, i.URL)
}

// Get fetches the most recent Items posted to the specified subreddit.
func Get(reddit string) ([]Item, error) {

	// construct the URL
	url := fmt.Sprintf("http://reddit.com/r/%s.json", reddit)

	// make the request
    resp, err := http.Get(url)                               
    if err != nil {
    	// nil slice and a non-nil error value
        return nil, err
    }

    // defer clean-up work
    defer resp.Body.Close()              
    // ^-- Defer a call to the response body's Close method,
    // to guarantee that we clean up after the HTTP request.
    // The call will be executed after the function returns.
    
	// make an error
    if resp.StatusCode != http.StatusOK {
    	// nil slice and a non-nil error value
        return nil, errors.New(resp.Status)
    }

    r := new(response)                        
    err = json.NewDecoder(resp.Body).Decode(r)
    if err != nil {
    	// nil slice and a non-nil error value
        return nil, err
    }

    // prepare the response
    items := make([]Item, len(r.Data.Children))
    // ^-- Use the make function to allocate an Item slice
    // big enough to store the response data

    // convert the response
    for i, child := range r.Data.Children {    
        items[i] = child.Data
    }

    // non-nil slice and a nil error value
    return items, nil
}
