# Define the target directory
$targetDirectory = "../testFiles"

# Ensure the target directory exists
if (-not (Test-Path -Path $targetDirectory)) {
    New-Item -ItemType Directory -Path $targetDirectory | Out-Null
}

# Generate 11 test files with 100 bytes of random data
for ($i = 1; $i -le 11; $i++) {
    $fileName = "testFile$i.txt"
    $filePath = Join-Path -Path $targetDirectory -ChildPath $fileName

    # Generate 100 bytes of random data
    $randomData = -join ((65..90) + (97..122) | Get-Random -Count 100 | ForEach-Object {[char]$_})
    $randomData | Out-File -FilePath $filePath -Encoding ASCII
}

Write-Host "11 test files with 100 bytes of random data have been created in $targetDirectory"