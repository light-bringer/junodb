//
//  Copyright 2023 PayPal Inc.
//
//  Licensed to the Apache Software Foundation (ASF) under one or more
//  contributor license agreements.  See the NOTICE file distributed with
//  this work for additional information regarding copyright ownership.
//  The ASF licenses this file to You under the Apache License, Version 2.0
//  (the "License"); you may not use this file except in compliance with
//  the License.  You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
//  Unless required by applicable law or agreed to in writing, software
//  distributed under the License is distributed on an "AS IS" BASIS,
//  WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
//  See the License for the specific language governing permissions and
//  limitations under the License.
//

package cli


// For TLS serevr, need to provide server.crt and server.pem.
// See main_test.go: cert, err := tls.LoadX509KeyPair("./server.crt", "./server.pem")
var ( // Change ip:port to that of JunoDB proxy service
	serverAddr    = "127.0.0.1:8080" // TCP 
	serverTls     = "127.0.0.1:5080" // TLS
)
