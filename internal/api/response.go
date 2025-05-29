package api

import (
	"math/rand"
	models "tidybeaver/pkg/models"
	"time"

	"github.com/google/uuid"
)

func CreateRandomResponse(path string) models.APIResponse {
	rand.Seed(time.Now().UnixNano())
	resp := responses[rand.Intn(len(responses))]
	msg := resp.Messages[rand.Intn(len(resp.Messages))]

	return models.APIResponse{
		StatusCode: resp.StatusCode,
		Status:     resp.Status,
		Message:    msg,
		Timestamp:  time.Now(),
		Path:       path,
		RequestID:  uuid.New().String(),
	}
}

var responses = []models.APIResponseVariant{
	// Info
	{StatusCode: 100, Status: "Continue", Messages: []string{
		"Request received, continuing processing.",
		"Headers received, awaiting body.",
		"Initial part of the request accepted.",
		"Please continue sending the request.",
	}},
	{StatusCode: 101, Status: "Switching Protocols", Messages: []string{
		"Protocol switch initiated.",
		"Server is switching protocols.",
		"Handshake complete for new protocol.",
		"Client requested a protocol change.",
	}},
	// Success
	{StatusCode: 200, Status: "OK", Messages: []string{
		"Request completed successfully.",
		"Everything is working as expected.",
		"Data fetched successfully.",
		"Operation executed correctly.",
	}},
	{StatusCode: 201, Status: "Created", Messages: []string{
		"New resource was created successfully.",
		"The user was successfully registered.",
		"Resource added to the database.",
		"Creation successful and acknowledged.",
	}},
	// Client Errors
	{StatusCode: 400, Status: "Bad Request", Messages: []string{
		"Malformed request syntax.",
		"Invalid parameters sent.",
		"Check your input format.",
		"Request could not be understood.",
	}},
	{StatusCode: 404, Status: "Not Found", Messages: []string{
		"The resource does not exist.",
		"Endpoint not found.",
		"Nothing was found at this URL.",
		"No matching route.",
	}},
	// Server Errors
	{StatusCode: 500, Status: "Internal Server Error", Messages: []string{
		"Something went wrong on our end.",
		"Unexpected condition encountered.",
		"We're fixing an internal issue.",
		"Oops! Server crashed temporarily.",
	}},
	{StatusCode: 503, Status: "Service Unavailable", Messages: []string{
		"Service is temporarily overloaded.",
		"Try again later, server busy.",
		"Service down for maintenance.",
		"System is currently unavailable.",
	}},
}
