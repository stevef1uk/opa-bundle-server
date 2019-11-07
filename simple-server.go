/*
Copyright 2018 The Knative Authors
Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at
    https://www.apache.org/licenses/LICENSE-2.0
Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

package main

import (
	"log"
        "io/ioutil"
        "bytes"
        "time"
	"net/http"
)

func handler(w http.ResponseWriter, r *http.Request) {
    log.Print("Hello world received a request.")
    data, err := ioutil.ReadFile("/tmp/bundle.tar.gz")
    if(err != nil){
        log.Fatal(err)
    }
    w.Header().Set("Content-Type", "application/octet-stream")
    w.Header().Set("Content-Disposition", "attachment; filename=" + "bundle.tar.gz")
    w.Header().Set("Content-Transfer-Encoding", "binary")
    w.Header().Set("Expires", "0")
    http.ServeContent(w, r, "Fred", time.Now(), bytes.NewReader(data))

}

func main() {
	log.Print("Hello world sample started.")

	http.HandleFunc("/bundles/istio/authz", handler)
	http.ListenAndServe(":8080", nil)
}

