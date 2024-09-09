@echo off

if "%1" == "build" (
    echo Building the project...
    go build -o bin\ecom.exe cmd\main.go
    echo Build successfully!
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
    echo Build successfully!
    echo Running the application...
    bin\ecom.exe
    exit /b
)

echo Invalid option. Use:
echo   build - to build the project
echo   test  - to run tests
echo   run   - to run the project
exit /b 1
