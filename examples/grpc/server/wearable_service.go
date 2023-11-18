package main

import (
	"math/rand"
	"time"

	wearablepb "github.com/norbybaru/go-proto-buf/gen/go/wearable/v1"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

type WearableServer struct {
	wearablepb.UnimplementedWearableServiceServer
}

func (w *WearableServer) BeatsPerMinute(
	req *wearablepb.BeatsPerMinuteRequest,
	stream wearablepb.WearableService_BeatsPerMinuteServer,
) error {

	for {
		select {
		case <-stream.Context().Done():
			return status.Error(codes.Canceled, "Stream has ended")
		case <-time.After(time.Second):
			value := 30 + rand.Int31n(80)

			err := stream.SendMsg(&wearablepb.BeatsPerMinuteResponse{
				Value:  uint32(value),
				Minute: uint32(time.Now().Second()),
			})

			if err != nil {
				errMsg := "Stream has ended with err: " + err.Error()
				return status.Error(codes.Canceled, errMsg)
			}
		}
	}
}
