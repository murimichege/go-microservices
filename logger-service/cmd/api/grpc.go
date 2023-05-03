package main

import (
	"context"
	"log-service/data"
	"log-service/logs"
)

// This is the grpc server. Its going to receive requests from client
type LogServer struct {
	logs.UnimplementedLogServiceServer
	Models data.Models
}

func (l *LogServer) WriteLog(ctx context.Context, req *logs.LogRequest) (*logs.LogResponse, error) {
	input := req.GetLogEntry()
	// write the log
	logEntry := data.LogEntry{
		Name: input.Name,
		Data: input.Data,
	}
	err := l.Models.LogEntry.Insert(logEntry)
	if err != nil {
		res := &logs.LogResponse{Result: "failed"}
		return res, err

	}
	res := &logs.LogResponse{Result: "logged"}
	return res, err
}
