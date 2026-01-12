$env:PATH = "C:\msys64\mingw64\bin;" + $env:PATH
$env:CGO_ENABLED = "1"
Set-Location "c:\Users\98317\Desktop\DMXAPI共享文件夹\网站文件\tokens_hs"
go build -ldflags "-H windowsgui" -o tokens_hs.exe
Write-Host "Build completed with exit code: $LASTEXITCODE"
