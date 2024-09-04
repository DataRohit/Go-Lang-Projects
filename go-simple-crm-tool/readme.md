# Go Stock Data API

## Packages Used

1. gorilla/mux
2. gorm.io/gorm
3. gorm.io/driver/postgres
4. google/uuid
5. net/http
7. encoding/json
8. github.com/sirupsen/logrus

## Routes

- `/leads` - POST - handlers.CreateMultipleLeadsHandler to create multiple leads together
- `/leads` - GET - handlers.GetAllLeadsHandler to get all the stored leads
- `/leads` - DELETE - handlers.DeleteMultipleLeadsHandler to delete multiple leads together using the ids
- `/leads` - PUT - handlers.UpdateMultipleLeadHandler to update multiple leads together

- `/lead` - POST - handlers.CreateSingleLeadHandler to create a single lead
- `/lead` - GET - handlers.GetRandomLeadHandler to get a random lead from database
- `/lead/{id}` - GET - handlers.GetLeadByIDHandler to get a specific lead by id
- `/lead/{id}` - DELETE - handlers.DeleteLeadByIDHandler to delete a specific lead by id
- `/lead/{id}` - PUT - handlers.UpdateLeadByIDHandler to update a specific lead by id

## Database

- PostgreSQL (Docker Postgres Latest Image)

### Architecture Diagram

<img src="./architecture-diagram.svg" alt="Architecture Diagram" style="width:100%;"/>
