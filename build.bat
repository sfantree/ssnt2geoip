@echo off

rem GOARCH=arm GOOS=linux GOARM=6 

REM SET CGO_ENABLED=0
REM SET GOOS=linux
REM SET GOARCH=arm
REM SET GOARM=6
REM G:\Go\bin\go.exe build -v -o go_build_ssnt2geoip_go_armv6 ssnt2geoip.go

SET CGO_ENABLED=0
SET GOOS=linux
SET GOARCH=amd64

G:\Go\bin\go.exe build -v -o go_build_ssnt2geoip_go_amd64 ssnt2geoip.go

pause

