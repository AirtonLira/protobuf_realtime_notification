package main

import (
	"context"
	"log"
	"time"

	pb "github.com/AirtonLira/protobuf_realtime_notification/internal/proto/notification"
	"github.com/AirtonLira/protobuf_realtime_notification/pkg/duckdb"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

func main() {
	conn, err := grpc.NewClient(":50051", grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}
	defer conn.Close()

	client := pb.NewNotificationServiceClient(conn)

	// Test GetNotifications
	testGetNotifications(client)

	// Test StreamNotifications
	// testStreamNotifications(client)
}

func testGetNotifications(client pb.NotificationServiceClient) {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()

	db, err := duckdb.InitDB()

	if err != nil {
		log.Fatalf("failed to load duckdb: %v ", err)
	}

	defer db.Close()

	req := &pb.NotificationRequest{UserId: "user123"}
	res, err := client.GetNotifications(ctx, req)
	if err != nil {
		log.Fatalf("could not get notifications: %v", err)
	}
	log.Printf("GetNotifications Response: %v", res.Notifications)

}

// func testStreamNotifications(client pb.NotificationServiceClient) {
// 	ctx, cancel := context.WithTimeout(context.Background(), time.Minute)
// 	defer cancel()

// 	req := &pb.NotificationRequest{UserId: "user123"}
// 	stream, err := client.StreamNotifications(ctx, req)
// 	if err != nil {
// 		log.Fatalf("could not stream notifications: %v", err)
// 	}

// 	for {
// 		notification, err := stream.Recv()
// 		if err != nil {
// 			log.Fatalf("error receiving notification: %v", err)
// 		}
// 		log.Printf("StreamNotifications Response: %v", notification)
// 	}
// }
