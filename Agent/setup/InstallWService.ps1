#Requires -RunAsAdministrator
cmd.exe /c "WinSW.exe install ..\SEMTA.xml"
Write-Output "Starting Service"
cmd.exe /c "sc.exe start SEMTA"
Write-Output "Done"
Pause