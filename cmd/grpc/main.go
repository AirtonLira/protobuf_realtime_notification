package main

import (
	"context"
	"log"
	"net"
	"time"

	pb "github.com/AirtonLira/protobuf_realtime_notification/internal/proto/notification"

	"google.golang.org/grpc"
)

type server struct {
	pb.UnimplementedNotificationServiceServer
}

func (s *server) GetNotifications(ctx context.Context, req *pb.NotificationRequest) (*pb.NotificationResponse, error) {
	notifications := []*pb.Notification{
		{Id: "1", Title: "Welcome!", Message: "Your account has been created successfully.", UserId: req.UserId, Timestamp: time.Now().Unix()},
	}
	return &pb.NotificationResponse{Notifications: notifications}, nil
}

func (s *server) StreamNotifications(req *pb.NotificationRequest, stream pb.NotificationService_StreamNotificationsServer) error {
	for {
		notification := &pb.Notification{Id: "2", Title: "New Event!", Message: "You have a new event.", UserId: req.UserId, Timestamp: time.Now().Unix()}
		if err := stream.Send(notification); err != nil {
			return err
		}
		time.Sleep(10 * time.Second) // Simulate new notifications arriving
	}
}

func main() {
	lis, err := net.Listen("tcp", ":50051")
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	s := grpc.NewServer()
	pb.RegisterNotificationServiceServer(s, &server{})
	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
