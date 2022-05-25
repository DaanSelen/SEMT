#Requires -RunAsAdministrator
cmd.exe /c "WinSW.exe install ..\SEMTS.xml"
Write-Output "Starting Service"
cmd.exe /c "sc.exe start SEMTS"
Write-Output "Done"
Pause