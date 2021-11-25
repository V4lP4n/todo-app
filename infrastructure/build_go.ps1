Set-Location backend 
go build -o ../srv.exe
Set-Location ..
echo "Builded"
.\srv.exe