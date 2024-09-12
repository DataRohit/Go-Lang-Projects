# Go Rate Limiting Techniques

## Packages Used

1. github.com/didip/tollbooth/v7
2. golang.org/x/time

## Toll Booth Rate Limiting

Tollbooth is a Go package for rate-limiting HTTP requests. It limits the number of requests a client can make to an API within a specific time window. By setting limits on the request rate, it helps prevent abuse and ensures fair usage of resources. Tollbooth can be customized to define limits based on factors like IP address or other client-specific identifiers and supports different HTTP frameworks.

## Token Bucket Rate Limiting

Token bucket rate limiting is an algorithm that controls the flow of requests by distributing tokens into a bucket at a fixed rate. Each request requires a token to proceed. If the bucket has available tokens, the request is allowed; if not, it's denied or delayed. The bucket has a maximum capacity, so if tokens aren’t used, they accumulate up to the limit. This method allows for bursts of traffic up to the bucket's capacity, while ensuring an overall average request rate.

## Client Rate Limiting

Client rate limiting is a technique used to control the number of requests a specific client (such as a user or IP address) can make to a server or API over a certain period of time. The goal is to prevent individual clients from overloading the system with too many requests, ensuring fair usage and protecting resources from abuse.

Different clients can have their own limits based on factors like IP addresses or user IDs, with common rate-limiting algorithms like token bucket or leaky bucket applied to enforce these limits.
