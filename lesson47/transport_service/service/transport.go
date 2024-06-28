package service

import (
	"context"
	pb "transport/generated/transport"
	"transport/storage/postgres"
)

type TransportService struct {
	pb.UnimplementedTransportServiceServer

	Transport *postgres.TransportRepo
}

func (ts *TransportService) GetBusSchedule(ctx context.Context, in *pb.GetBusScheduleRequest) (*pb.GetBusScheduleResponse, error) {
	resp, err := ts.Transport.BusSchedule(in)

	return resp, err
}

func (tc *TransportService) TrackBusLocation(ctx context.Context, in *pb.TrackBusLocationRequest) (*pb.TrackBusLocationResponse, error) {
	resp, err := tc.Transport.TrackBusLocation(in)

	return resp, err
}

func (tc *TransportService) ReportTrafficJam(ctx context.Context, in *pb.ReportTrafficJamRequest) (*pb.ReportTrafficJamResponse, error) {
	resp, err := tc.Transport.ReportTrafficJam(in)

	return resp, err
}