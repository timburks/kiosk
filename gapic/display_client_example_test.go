// Copyright 2022 Google LLC
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

package gapic_test

import (
	"context"

	gapic "github.com/googleapis/kiosk/gapic"
	kioskpb "github.com/googleapis/kiosk/rpc"
	emptypb "google.golang.org/protobuf/types/known/emptypb"
)

func ExampleNewDisplayClient() {
	ctx := context.Background()
	// This snippet has been automatically generated and should be regarded as a code template only.
	// It will require modifications to work:
	// - It may require correct/in-range values for request initialization.
	// - It may require specifying regional endpoints when creating the service client as shown in:
	//   https://pkg.go.dev/cloud.google.com/go#hdr-Client_Options
	c, err := gapic.NewDisplayClient(ctx)
	if err != nil {
		// TODO: Handle error.
	}
	defer c.Close()

	// TODO: Use client.
	_ = c
}

func ExampleDisplayClient_CreateKiosk() {
	ctx := context.Background()
	// This snippet has been automatically generated and should be regarded as a code template only.
	// It will require modifications to work:
	// - It may require correct/in-range values for request initialization.
	// - It may require specifying regional endpoints when creating the service client as shown in:
	//   https://pkg.go.dev/cloud.google.com/go#hdr-Client_Options
	c, err := gapic.NewDisplayClient(ctx)
	if err != nil {
		// TODO: Handle error.
	}
	defer c.Close()

	req := &kioskpb.Kiosk{
		// TODO: Fill request struct fields.
		// See https://pkg.go.dev/github.com/googleapis/kiosk/rpc#Kiosk.
	}
	resp, err := c.CreateKiosk(ctx, req)
	if err != nil {
		// TODO: Handle error.
	}
	// TODO: Use resp.
	_ = resp
}

func ExampleDisplayClient_ListKiosks() {
	ctx := context.Background()
	// This snippet has been automatically generated and should be regarded as a code template only.
	// It will require modifications to work:
	// - It may require correct/in-range values for request initialization.
	// - It may require specifying regional endpoints when creating the service client as shown in:
	//   https://pkg.go.dev/cloud.google.com/go#hdr-Client_Options
	c, err := gapic.NewDisplayClient(ctx)
	if err != nil {
		// TODO: Handle error.
	}
	defer c.Close()

	req := &emptypb.Empty{
		// TODO: Fill request struct fields.
		// See https://pkg.go.dev/google.golang.org/protobuf/types/known/emptypb#Empty.
	}
	resp, err := c.ListKiosks(ctx, req)
	if err != nil {
		// TODO: Handle error.
	}
	// TODO: Use resp.
	_ = resp
}

func ExampleDisplayClient_GetKiosk() {
	ctx := context.Background()
	// This snippet has been automatically generated and should be regarded as a code template only.
	// It will require modifications to work:
	// - It may require correct/in-range values for request initialization.
	// - It may require specifying regional endpoints when creating the service client as shown in:
	//   https://pkg.go.dev/cloud.google.com/go#hdr-Client_Options
	c, err := gapic.NewDisplayClient(ctx)
	if err != nil {
		// TODO: Handle error.
	}
	defer c.Close()

	req := &kioskpb.GetKioskRequest{
		// TODO: Fill request struct fields.
		// See https://pkg.go.dev/github.com/googleapis/kiosk/rpc#GetKioskRequest.
	}
	resp, err := c.GetKiosk(ctx, req)
	if err != nil {
		// TODO: Handle error.
	}
	// TODO: Use resp.
	_ = resp
}

func ExampleDisplayClient_DeleteKiosk() {
	ctx := context.Background()
	// This snippet has been automatically generated and should be regarded as a code template only.
	// It will require modifications to work:
	// - It may require correct/in-range values for request initialization.
	// - It may require specifying regional endpoints when creating the service client as shown in:
	//   https://pkg.go.dev/cloud.google.com/go#hdr-Client_Options
	c, err := gapic.NewDisplayClient(ctx)
	if err != nil {
		// TODO: Handle error.
	}
	defer c.Close()

	req := &kioskpb.DeleteKioskRequest{
		// TODO: Fill request struct fields.
		// See https://pkg.go.dev/github.com/googleapis/kiosk/rpc#DeleteKioskRequest.
	}
	err = c.DeleteKiosk(ctx, req)
	if err != nil {
		// TODO: Handle error.
	}
}

func ExampleDisplayClient_CreateSign() {
	ctx := context.Background()
	// This snippet has been automatically generated and should be regarded as a code template only.
	// It will require modifications to work:
	// - It may require correct/in-range values for request initialization.
	// - It may require specifying regional endpoints when creating the service client as shown in:
	//   https://pkg.go.dev/cloud.google.com/go#hdr-Client_Options
	c, err := gapic.NewDisplayClient(ctx)
	if err != nil {
		// TODO: Handle error.
	}
	defer c.Close()

	req := &kioskpb.Sign{
		// TODO: Fill request struct fields.
		// See https://pkg.go.dev/github.com/googleapis/kiosk/rpc#Sign.
	}
	resp, err := c.CreateSign(ctx, req)
	if err != nil {
		// TODO: Handle error.
	}
	// TODO: Use resp.
	_ = resp
}

func ExampleDisplayClient_ListSigns() {
	ctx := context.Background()
	// This snippet has been automatically generated and should be regarded as a code template only.
	// It will require modifications to work:
	// - It may require correct/in-range values for request initialization.
	// - It may require specifying regional endpoints when creating the service client as shown in:
	//   https://pkg.go.dev/cloud.google.com/go#hdr-Client_Options
	c, err := gapic.NewDisplayClient(ctx)
	if err != nil {
		// TODO: Handle error.
	}
	defer c.Close()

	req := &emptypb.Empty{
		// TODO: Fill request struct fields.
		// See https://pkg.go.dev/google.golang.org/protobuf/types/known/emptypb#Empty.
	}
	resp, err := c.ListSigns(ctx, req)
	if err != nil {
		// TODO: Handle error.
	}
	// TODO: Use resp.
	_ = resp
}

func ExampleDisplayClient_GetSign() {
	ctx := context.Background()
	// This snippet has been automatically generated and should be regarded as a code template only.
	// It will require modifications to work:
	// - It may require correct/in-range values for request initialization.
	// - It may require specifying regional endpoints when creating the service client as shown in:
	//   https://pkg.go.dev/cloud.google.com/go#hdr-Client_Options
	c, err := gapic.NewDisplayClient(ctx)
	if err != nil {
		// TODO: Handle error.
	}
	defer c.Close()

	req := &kioskpb.GetSignRequest{
		// TODO: Fill request struct fields.
		// See https://pkg.go.dev/github.com/googleapis/kiosk/rpc#GetSignRequest.
	}
	resp, err := c.GetSign(ctx, req)
	if err != nil {
		// TODO: Handle error.
	}
	// TODO: Use resp.
	_ = resp
}

func ExampleDisplayClient_DeleteSign() {
	ctx := context.Background()
	// This snippet has been automatically generated and should be regarded as a code template only.
	// It will require modifications to work:
	// - It may require correct/in-range values for request initialization.
	// - It may require specifying regional endpoints when creating the service client as shown in:
	//   https://pkg.go.dev/cloud.google.com/go#hdr-Client_Options
	c, err := gapic.NewDisplayClient(ctx)
	if err != nil {
		// TODO: Handle error.
	}
	defer c.Close()

	req := &kioskpb.DeleteSignRequest{
		// TODO: Fill request struct fields.
		// See https://pkg.go.dev/github.com/googleapis/kiosk/rpc#DeleteSignRequest.
	}
	err = c.DeleteSign(ctx, req)
	if err != nil {
		// TODO: Handle error.
	}
}

func ExampleDisplayClient_SetSignIdForKioskIds() {
	ctx := context.Background()
	// This snippet has been automatically generated and should be regarded as a code template only.
	// It will require modifications to work:
	// - It may require correct/in-range values for request initialization.
	// - It may require specifying regional endpoints when creating the service client as shown in:
	//   https://pkg.go.dev/cloud.google.com/go#hdr-Client_Options
	c, err := gapic.NewDisplayClient(ctx)
	if err != nil {
		// TODO: Handle error.
	}
	defer c.Close()

	req := &kioskpb.SetSignIdForKioskIdsRequest{
		// TODO: Fill request struct fields.
		// See https://pkg.go.dev/github.com/googleapis/kiosk/rpc#SetSignIdForKioskIdsRequest.
	}
	err = c.SetSignIdForKioskIds(ctx, req)
	if err != nil {
		// TODO: Handle error.
	}
}

func ExampleDisplayClient_GetSignIdForKioskId() {
	ctx := context.Background()
	// This snippet has been automatically generated and should be regarded as a code template only.
	// It will require modifications to work:
	// - It may require correct/in-range values for request initialization.
	// - It may require specifying regional endpoints when creating the service client as shown in:
	//   https://pkg.go.dev/cloud.google.com/go#hdr-Client_Options
	c, err := gapic.NewDisplayClient(ctx)
	if err != nil {
		// TODO: Handle error.
	}
	defer c.Close()

	req := &kioskpb.GetSignIdForKioskIdRequest{
		// TODO: Fill request struct fields.
		// See https://pkg.go.dev/github.com/googleapis/kiosk/rpc#GetSignIdForKioskIdRequest.
	}
	resp, err := c.GetSignIdForKioskId(ctx, req)
	if err != nil {
		// TODO: Handle error.
	}
	// TODO: Use resp.
	_ = resp
}
