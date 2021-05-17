# git-clone-open-code
Git Clone Open in VSCode


Replaced with Simple Powershell Function

function gcode {
  $url = $args[0]
  $folder = Split-Path $url -Leaf
  if(Test-Path $folder){
	Write-Output "Opening existing folder"
	code $folder
  }else{
	git clone $url && code $folder
  }
}