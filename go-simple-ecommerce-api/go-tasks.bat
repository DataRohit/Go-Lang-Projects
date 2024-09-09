@echo off

if "%1" == "build" (
    echo Building the project...
    go build -o bin\ecom.exe cmd\main.go
    if %errorlevel% neq 0 (
        echo Build failed!
        exit /b %errorlevel%
    )
    echo Build successfully!
    exit /b
)

if "%1" == "test" (
    echo Running tests...
    go test -v ./...
    if %errorlevel% neq 0 (
        echo Tests failed!
        exit /b %errorlevel%
    )
    echo Tests successful!
    exit /b
)

if "%1" == "run" (
    echo Running the application...
    bin\ecom
    exit /b
)

echo Invalid option. Use:
echo   build - to build the project
echo   test  - to run tests
echo   run   - to run the project
exit /b 1
