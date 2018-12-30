/*
 *
 * Copyright 2015 gRPC authors.
 *
 * Licensed under the Apache License, Version 2.0 (the "License");
 * you may not use this file except in compliance with the License.
 * You may obtain a copy of the License at
 *
 *     http://www.apache.org/licenses/LICENSE-2.0
 *
 * Unless required by applicable law or agreed to in writing, software
 * distributed under the License is distributed on an "AS IS" BASIS,
 * WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
 * See the License for the specific language governing permissions and
 * limitations under the License.
 *
 */

package main

import (
	"context"
	"log"
	"net"
	"fmt"
	"net/http"
	"io/ioutil"
	"os"
	"encoding/json"

	"google.golang.org/grpc"
	pb ".."
	"google.golang.org/grpc/reflection"
)

const (
	port = ":50051"
)

//Struct
type IssLocation struct {
	Message     string `json:"message"`
	Timestamp   int64    `json:"timestamp"`
	IssPosition struct {
		Longitude string `json:"longitude"`
		Latitude  string `json:"latitude"`
	} `json:"iss_position"`
}

// sample
// 20181230152434
// http://api.open-notify.org/iss-now.json

// {
//   "message": "success",
//   "timestamp": 1546151073,
//   "iss_position": {
//     "longitude": "-167.9689",
//     "latitude": "-51.5217"
//   }
// }

// server is used to implement helloworld.GreeterServer.
type server struct{}

// SayHello implements helloworld.GreeterServer
func (s *server) GetIssLocation(ctx context.Context, in *pb.IssRequest) (*pb.IssResponse, error) {
	log.Printf("request has been recieved")

	// get ISS Location with calling an API
	response, err := http.Get("http://api.open-notify.org/iss-now.json")

	if err != nil {
		fmt.Print(err.Error())
		os.Exit(1)
	}

	data, err := ioutil.ReadAll(response.Body)
	if err != nil {
		log.Fatal(err)
	}
	json_byte := ([]byte)(data)
	iss_location := new(IssLocation)

	json.Unmarshal(json_byte, &iss_location)
	
	// for debug
	// fmt.Println(time.Unix(iss_location.Timestamp, 0))
	// fmt.Println(iss_location.IssPosition)

	return &pb.IssResponse{Longitude: iss_location.IssPosition.Longitude, Latitude: iss_location.IssPosition.Latitude}, nil
}

func main() {
	lis, err := net.Listen("tcp", port)
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	s := grpc.NewServer()
	pb.RegisterIssServer(s, &server{})
	// Register reflection service on gRPC server.
	reflection.Register(s)
	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
