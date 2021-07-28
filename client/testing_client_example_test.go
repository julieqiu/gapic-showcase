// Copyright 2021 Google LLC
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     https://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

// Code generated by protoc-gen-go_gapic. DO NOT EDIT.

package client_test

import (
	"context"

	client "github.com/googleapis/gapic-showcase/client"
	genprotopb "github.com/googleapis/gapic-showcase/server/genproto"
	"google.golang.org/api/iterator"
	locationpb "google.golang.org/genproto/googleapis/cloud/location"
	iampb "google.golang.org/genproto/googleapis/iam/v1"
	longrunningpb "google.golang.org/genproto/googleapis/longrunning"
)

func ExampleNewTestingClient() {
	ctx := context.Background()
	c, err := client.NewTestingClient(ctx)
	if err != nil {
		// TODO: Handle error.
	}
	defer c.Close()

	// TODO: Use client.
	_ = c
}

func ExampleNewTestingRESTClient() {
	ctx := context.Background()
	c, err := client.NewTestingRESTClient(ctx)
	if err != nil {
		// TODO: Handle error.
	}
	defer c.Close()

	// TODO: Use client.
	_ = c
}

func ExampleTestingClient_CreateSession() {
	ctx := context.Background()
	c, err := client.NewTestingClient(ctx)
	if err != nil {
		// TODO: Handle error.
	}
	defer c.Close()

	req := &genprotopb.CreateSessionRequest{
		// TODO: Fill request struct fields.
	}
	resp, err := c.CreateSession(ctx, req)
	if err != nil {
		// TODO: Handle error.
	}
	// TODO: Use resp.
	_ = resp
}

func ExampleTestingClient_GetSession() {
	ctx := context.Background()
	c, err := client.NewTestingClient(ctx)
	if err != nil {
		// TODO: Handle error.
	}
	defer c.Close()

	req := &genprotopb.GetSessionRequest{
		// TODO: Fill request struct fields.
	}
	resp, err := c.GetSession(ctx, req)
	if err != nil {
		// TODO: Handle error.
	}
	// TODO: Use resp.
	_ = resp
}

func ExampleTestingClient_ListSessions() {
	ctx := context.Background()
	c, err := client.NewTestingClient(ctx)
	if err != nil {
		// TODO: Handle error.
	}
	defer c.Close()

	req := &genprotopb.ListSessionsRequest{
		// TODO: Fill request struct fields.
	}
	it := c.ListSessions(ctx, req)
	for {
		resp, err := it.Next()
		if err == iterator.Done {
			break
		}
		if err != nil {
			// TODO: Handle error.
		}
		// TODO: Use resp.
		_ = resp
	}
}

func ExampleTestingClient_DeleteSession() {
	ctx := context.Background()
	c, err := client.NewTestingClient(ctx)
	if err != nil {
		// TODO: Handle error.
	}
	defer c.Close()

	req := &genprotopb.DeleteSessionRequest{
		// TODO: Fill request struct fields.
	}
	err = c.DeleteSession(ctx, req)
	if err != nil {
		// TODO: Handle error.
	}
}

func ExampleTestingClient_ReportSession() {
	ctx := context.Background()
	c, err := client.NewTestingClient(ctx)
	if err != nil {
		// TODO: Handle error.
	}
	defer c.Close()

	req := &genprotopb.ReportSessionRequest{
		// TODO: Fill request struct fields.
	}
	resp, err := c.ReportSession(ctx, req)
	if err != nil {
		// TODO: Handle error.
	}
	// TODO: Use resp.
	_ = resp
}

func ExampleTestingClient_ListTests() {
	ctx := context.Background()
	c, err := client.NewTestingClient(ctx)
	if err != nil {
		// TODO: Handle error.
	}
	defer c.Close()

	req := &genprotopb.ListTestsRequest{
		// TODO: Fill request struct fields.
	}
	it := c.ListTests(ctx, req)
	for {
		resp, err := it.Next()
		if err == iterator.Done {
			break
		}
		if err != nil {
			// TODO: Handle error.
		}
		// TODO: Use resp.
		_ = resp
	}
}

func ExampleTestingClient_DeleteTest() {
	ctx := context.Background()
	c, err := client.NewTestingClient(ctx)
	if err != nil {
		// TODO: Handle error.
	}
	defer c.Close()

	req := &genprotopb.DeleteTestRequest{
		// TODO: Fill request struct fields.
	}
	err = c.DeleteTest(ctx, req)
	if err != nil {
		// TODO: Handle error.
	}
}

func ExampleTestingClient_VerifyTest() {
	ctx := context.Background()
	c, err := client.NewTestingClient(ctx)
	if err != nil {
		// TODO: Handle error.
	}
	defer c.Close()

	req := &genprotopb.VerifyTestRequest{
		// TODO: Fill request struct fields.
	}
	resp, err := c.VerifyTest(ctx, req)
	if err != nil {
		// TODO: Handle error.
	}
	// TODO: Use resp.
	_ = resp
}

func ExampleTestingClient_ListLocations() {
	ctx := context.Background()
	c, err := client.NewTestingClient(ctx)
	if err != nil {
		// TODO: Handle error.
	}
	defer c.Close()

	req := &locationpb.ListLocationsRequest{
		// TODO: Fill request struct fields.
	}
	it := c.ListLocations(ctx, req)
	for {
		resp, err := it.Next()
		if err == iterator.Done {
			break
		}
		if err != nil {
			// TODO: Handle error.
		}
		// TODO: Use resp.
		_ = resp
	}
}

func ExampleTestingClient_GetLocation() {
	ctx := context.Background()
	c, err := client.NewTestingClient(ctx)
	if err != nil {
		// TODO: Handle error.
	}
	defer c.Close()

	req := &locationpb.GetLocationRequest{
		// TODO: Fill request struct fields.
	}
	resp, err := c.GetLocation(ctx, req)
	if err != nil {
		// TODO: Handle error.
	}
	// TODO: Use resp.
	_ = resp
}

func ExampleTestingClient_SetIamPolicy() {
	ctx := context.Background()
	c, err := client.NewTestingClient(ctx)
	if err != nil {
		// TODO: Handle error.
	}
	defer c.Close()

	req := &iampb.SetIamPolicyRequest{
		// TODO: Fill request struct fields.
	}
	resp, err := c.SetIamPolicy(ctx, req)
	if err != nil {
		// TODO: Handle error.
	}
	// TODO: Use resp.
	_ = resp
}

func ExampleTestingClient_GetIamPolicy() {
	ctx := context.Background()
	c, err := client.NewTestingClient(ctx)
	if err != nil {
		// TODO: Handle error.
	}
	defer c.Close()

	req := &iampb.GetIamPolicyRequest{
		// TODO: Fill request struct fields.
	}
	resp, err := c.GetIamPolicy(ctx, req)
	if err != nil {
		// TODO: Handle error.
	}
	// TODO: Use resp.
	_ = resp
}

func ExampleTestingClient_TestIamPermissions() {
	ctx := context.Background()
	c, err := client.NewTestingClient(ctx)
	if err != nil {
		// TODO: Handle error.
	}
	defer c.Close()

	req := &iampb.TestIamPermissionsRequest{
		// TODO: Fill request struct fields.
	}
	resp, err := c.TestIamPermissions(ctx, req)
	if err != nil {
		// TODO: Handle error.
	}
	// TODO: Use resp.
	_ = resp
}

func ExampleTestingClient_ListOperations() {
	ctx := context.Background()
	c, err := client.NewTestingClient(ctx)
	if err != nil {
		// TODO: Handle error.
	}
	defer c.Close()

	req := &longrunningpb.ListOperationsRequest{
		// TODO: Fill request struct fields.
	}
	it := c.ListOperations(ctx, req)
	for {
		resp, err := it.Next()
		if err == iterator.Done {
			break
		}
		if err != nil {
			// TODO: Handle error.
		}
		// TODO: Use resp.
		_ = resp
	}
}

func ExampleTestingClient_GetOperation() {
	ctx := context.Background()
	c, err := client.NewTestingClient(ctx)
	if err != nil {
		// TODO: Handle error.
	}
	defer c.Close()

	req := &longrunningpb.GetOperationRequest{
		// TODO: Fill request struct fields.
	}
	resp, err := c.GetOperation(ctx, req)
	if err != nil {
		// TODO: Handle error.
	}
	// TODO: Use resp.
	_ = resp
}

func ExampleTestingClient_DeleteOperation() {
	ctx := context.Background()
	c, err := client.NewTestingClient(ctx)
	if err != nil {
		// TODO: Handle error.
	}
	defer c.Close()

	req := &longrunningpb.DeleteOperationRequest{
		// TODO: Fill request struct fields.
	}
	err = c.DeleteOperation(ctx, req)
	if err != nil {
		// TODO: Handle error.
	}
}

func ExampleTestingClient_CancelOperation() {
	ctx := context.Background()
	c, err := client.NewTestingClient(ctx)
	if err != nil {
		// TODO: Handle error.
	}
	defer c.Close()

	req := &longrunningpb.CancelOperationRequest{
		// TODO: Fill request struct fields.
	}
	err = c.CancelOperation(ctx, req)
	if err != nil {
		// TODO: Handle error.
	}
}
