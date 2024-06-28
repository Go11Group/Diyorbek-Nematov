package postgres

import (
	"database/sql"
	pb "transport/generated/transport"
)

type TransportRepo struct {
	DB *sql.DB
}

func NewTransportRepo(db *sql.DB) *TransportRepo {
	return &TransportRepo{DB: db}
}

func (t *TransportRepo) BusSchedule(s *pb.GetBusScheduleRequest) (*pb.GetBusScheduleResponse, error) {
	schedules := []*pb.BusSchedule{}
	rows, err := t.DB.Query(`
		SELECT
			stop,
			arrival_time
		FROM
			schedule
		WHERE 
			bus_number = $1
	`, s.BusNumber)

	if err != nil {
		return nil, err
	}

	for rows.Next() {
		schedule := &pb.BusSchedule{}

		if err = rows.Scan(&schedule.Stop, &schedule.ArrivalTime); err != nil {
			return nil, err
		}

		schedules = append(schedules, schedule)
	}

	return &pb.GetBusScheduleResponse{Schedules: schedules}, nil
}

func (t *TransportRepo) TrackBusLocation(s *pb.TrackBusLocationRequest) (*pb.TrackBusLocationResponse, error) {
	resp := pb.TrackBusLocationResponse{}

	err := t.DB.QueryRow(`
		SELECT
			bus_number,
			location
		FROM
			transport_info
		WHERE 
			bus_number = $1
	`, s.BusNumber).Scan(&resp.BusNumber, &resp.Location)

	return &resp, err
}

func (t *TransportRepo) ReportTrafficJam(r *pb.ReportTrafficJamRequest) (*pb.ReportTrafficJamResponse, error) {
	var status pb.ReportTrafficJamResponse

	err := t.DB.QueryRow(`
		SELECT
			status
		FROM
			transport_info
		WHERE
			bus_number = $1 and location = $2
	`, r.BusNumber, r.Location).Scan(&status.Status)

	return &status, err
}