@echo off
set PATH=C:\msys64\mingw64\bin;%PATH%
set CGO_ENABLED=1
go build -ldflags "-H windowsgui" -o tokens_hs.exe
echo Build complete!
pause
