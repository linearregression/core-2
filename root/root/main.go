// Copyright (c) 2016 Pani Networks
// All Rights Reserved.
//
// Licensed under the Apache License, Version 2.0 (the "License"); you may
// not use this file except in compliance with the License. You may obtain
// a copy of the License at
//
//   http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS, WITHOUT
// WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied. See the
// License for the specific language governing permissions and limitations
// under the License.

// Command to launch root service
package main

import (
	"flag"
	"fmt"
	"github.com/romana/core/common"
	"github.com/romana/core/root"
	"log"
)

// Main entry point for the root microservice
func main() {
	configFileName := flag.String("c", "", "Configuration file")
	version := flag.Bool("version", false, "Build Information.")
	flag.Parse()
	if *version {
		fmt.Println(common.BuildInfo())
		return
	}
	svcInfo, err := root.Run(*configFileName)
	if err != nil {
		panic(err)
	}
	for {
		msg := <-svcInfo.Channel
		log.Println(msg)
	}
}
