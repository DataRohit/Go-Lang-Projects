@echo off

if "%1" == "build" (
    echo Building the project...
    go build -o bin\ecom.exe cmd\main.go
    echo Build successful!
    exit /b
)

if "%1" == "test" (
    echo Running tests...
    go test -v ./...
    echo Tests successful!
    exit /b
)

if "%1" == "run" (
    echo Building the project...
    go build -o bin\ecom.exe cmd\main.go
    echo Build successful!
    echo Running the application...
    bin\ecom.exe
    exit /b
)

if "%1" == "migration" (
    echo Creating migrations...
    shift
    if "%~1"=="" (
        echo Please provide a migration name.
        exit /b 1
    )
    migrate create -ext sql -dir cmd/migrate/migrations %*
    echo Migrations created successfully!
    exit /b
)

if "%1" == "migrate-up" (
    echo Applying migrations...
    go run cmd/migrate/main.go up
    echo Migrations applied successfully!
    exit /b
)

if "%1" == "migrate-down" (
    echo Reverting migrations...
    go run cmd/migrate/main.go down
    echo Migrations reverted successfully!
    exit /b
)

echo Invalid option. Use:
echo   build - to build the project
echo   test  - to run tests
echo   run   - to run the project
echo   migration [name] - to create a migration
echo   migrate-up - to apply migrations
echo   migrate-down - to revert migrations
exit /b 1
