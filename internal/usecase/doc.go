package usecase

// Use case Layer
// This layer will act as a logic processor for controller layer and repository layer.
//
// Responsibility
// In this package we will do any kind of business logic, from calling func that query to database, do the caching and
// communicate with different domain.
// Communicate with different domain also happen in this layer by injecting other domain usecase to
// the designated domain usecase.
