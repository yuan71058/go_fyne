@echo off
chcp 65001 >nul
echo Building Fyne GUI application...

REM Set environment variables
set CGO_ENABLED=1
set GOOS=windows
set GOARCH=amd64

REM Compile GUI program without command line window
go build -ldflags="-H=windowsgui" -o fyne-app-gui.exe main.go

if %ERRORLEVEL% EQU 0 (
    echo.
    echo Build successful!
    echo Generated executable: fyne-app-gui.exe
) else (
    echo.
    echo Build failed, please check error messages!
)