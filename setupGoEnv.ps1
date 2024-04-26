if (-NOT ([Security.Principal.WindowsPrincipal][Security.Principal.WindowsIdentity]::GetCurrent()).IsInRole([Security.Principal.WindowsBuiltInRole] "Administrator"))   
{   
    $arguments = "& '" + $myinvocation.mycommand.definition + "'"
    Start-Process powershell -Verb runAs -ArgumentList $arguments
    Break 
}

# Pobierz ścieżki do skryptu
$SCRIPT_PATH = Split-Path -Parent $MyInvocation.MyCommand.Definition
$NEIGHBOR_PATH = Join-Path -Path $SCRIPT_PATH -ChildPath "imports"
$BIN_PATH = Join-Path -Path $SCRIPT_PATH -ChildPath "builds"

Write-Host "Eksportowanie GOPATH do $NEIGHBOR_PATH"
[Environment]::SetEnvironmentVariable
    ($NEIGHBOR_PATH, $env:GOPATH, [System.EnvironmentVariableTarget]::User)

Write-Host "Eksportowanie GOBIN do $BIN_PATH"
[Environment]::SetEnvironmentVariable
    ($BIN_PATH, $env:GOBIN, [System.EnvironmentVariableTarget]::User)

Write-Host "Jeżeli skrypt nie zadziałał, musisz dodać ręcznie ściezki do patha."
Write-Host "Kliknij cokolwiek, by zako�czy� skrypt"
Read-Host