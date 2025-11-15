@echo off
chcp 65001 >nul
title Fyne Application Builder

:menu
cls
echo ====================================
echo    Fyne Application Builder
echo ====================================
echo.
echo 1. Build GUI Version (No Console)
echo 2. Build Console Version
echo 3. Run GUI Version
echo 4. Run Console Version
echo 5. Clean Build Files
echo 6. Exit
echo.
set /p choice=Please select an option (1-6): 

if "%choice%"=="1" goto build_gui
if "%choice%"=="2" goto build_console
if "%choice%"=="3" goto run_gui
if "%choice%"=="4" goto run_console
if "%choice%"=="5" goto clean
if "%choice%"=="6" goto exit
goto menu

:build_gui
echo.
echo Building GUI version (without console window)...
echo.

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
    echo.
    pause
) else (
    echo.
    echo Build failed, please check error messages!
    echo.
    pause
)
goto menu

:build_console
echo.
echo Building console version...
echo.

REM Set environment variables
set CGO_ENABLED=1
set GOOS=windows
set GOARCH=amd64

REM Compile with console
go build -o fyne-app.exe main.go

if %ERRORLEVEL% EQU 0 (
    echo.
    echo Build successful!
    echo Generated executable: fyne-app.exe
    echo.
    pause
) else (
    echo.
    echo Build failed, please check error messages!
    echo.
    pause
)
goto menu

:run_gui
echo.
echo Running GUI version...
if exist fyne-app-gui.exe (
    start fyne-app-gui.exe
    echo Application started.
) else (
    echo GUI version not found. Please build it first.
    pause
)
goto menu

:run_console
echo.
echo Running console version...
if exist fyne-app.exe (
    start fyne-app.exe
    echo Application started.
) else (
    echo Console version not found. Please build it first.
    pause
)
goto menu

:clean
echo.
echo Cleaning build files...
if exist fyne-app.exe del fyne-app.exe
if exist fyne-app-gui.exe del fyne-app-gui.exe
echo Clean completed.
echo.
pause
goto menu

:exit
echo.
echo Goodbye!
exit /b 0