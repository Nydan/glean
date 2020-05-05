package repository

// Repository layer
// This layer act as data layer for the project. It handle database queries, redis cache and HTTP call
// to the other microservice. It should not do any business logic since all the value already handled by use case
// layer. Simply CRUD to DB or http call, so the dependencies to this layer would be implementation of DB connection,
// redis connection, or http client instance.
