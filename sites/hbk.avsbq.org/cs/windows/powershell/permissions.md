# Change file permissions

`takeown /F E:\Users /R /D Y`

replace `E` with the drive letter

The /F switch specifies the file or directory, /R makes the command recursive, and /D Y assumes "Yes" for any confirmation prompts

`icacls E:\Users /grant %username%:F /T`

The /grant switch specifies the permissions to grant, %username% is your current user, :F indicates full control, and /T makes the command recursive
