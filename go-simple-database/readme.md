# Go-Simple-Database

**_A lightweight, thread-safe database built with Go._** This database uses **mutex locks** to prevent race conditions and **JSON files** for persistent storage of records. Ideal for simple data storage needs.

## Features

-   `Write:` Add a new record to the database.
-   `Read:` Retrieve a specific record by its key.
-   `Read All:` Fetch all records from the database.
-   `Delete:` Remove a record by its key.

## Packages Used

-   `encoding/json:` To handle JSON encoding and decoding.
-   `path/filepath:` For working with file paths.
-   `sync:` For handling mutex locks.
-   `os:` To interact with the file system.

## Architecture Diagram

<img src="./architecture-diagram.svg" alt="Architecture Diagram" style="width:100%;"/>
