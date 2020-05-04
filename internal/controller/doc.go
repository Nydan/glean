package controller

// Controller Layer
// This layer will act as a presenter for deliver your API response.
// In this package we can define API support for HTTP REST, GRPC and/or HTML.
//
// Responsibility
// This layer is going to accept input from client based on what kind of protocol being used by the client.
// The validation of client request happen in this layer such as required parameters, formatting, etc.
// Controller layer will call usecase layer for processing the client request and response to client.
