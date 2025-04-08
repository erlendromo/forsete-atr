@echo off
setlocal enabledelayedexpansion

:: Install Git LFS
echo Installing Git LFS...
git lfs install

:: Install Make (via Chocolatey)
echo Installing Make...
choco install make

:: Install Go 1.23.4
echo Installing Go 1.23.4...
choco install golang --version=1.23.4

:: Install Docker (via Chocolatey)
echo Installing Docker...
choco install docker-desktop

:: Install Docker Compose Plugin
echo Installing Docker Compose Plugin...
choco install docker-compose

:: Install Swagger (Swag) for Go
echo Installing Swagger for Go...
go install github.com/swaggo/http-swagger/v2@latest

:: Create model directories
echo Creating model directories...
mkdir models\linesegmentation
mkdir models\regionsegmentation
mkdir models\textrecognition

:: Clone Hugging Face models
echo Cloning Hugging Face models...
git lfs clone https://huggingface.co/Riksarkivet/yolov9-lines-within-regions-1 models\linesegmentation\yolov9-lines-within-regions-1
git lfs clone https://huggingface.co/Sprakbanken/TrOCR-norhand-v3 models\textrecognition\TrOCR-norhand-v3

:: Create .env file
echo Creating .env file...
move example.env .env

echo Setup complete!
pause
