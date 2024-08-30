# Go Stock Data API

## Packages Used

1. gorilla/mux
2. gorm.io/gorm
3. gorm.io/driver/postgres
4. google/uuid
5. net/http
6. log
7. encoding/json

## Routes

- `/stocks` - GetAllStocks() handler to fetch all the available stocks
- `/stocks/{symbol}` - GetStockBySymbol() handler to fetch a single stock using the `stockId`
- `/stocks` - CreateStock() handler to create a new stock with unique uuid
- `/stocks/{symbol}` - DeleteStock() handler to delete an instance of stock from the db
- `/stocks/{symbol}` - UpdateStock() handler to update an existing stock record in db

## Database

- PostgreSQL

### Folder Structure

<img src="./folder-structure.svg" alt="Folder Structure" style="width:100%;"/>

### Architecture Diagram

<img src="./architecture-diagram.svg" alt="Architecture Diagram" style="width:100%;"/>
