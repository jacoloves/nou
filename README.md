# nou
nou(no-option-usage) is a tool that outputs Usage even when there are no options in Go. 

## Usage
1. If nothing is passed as an argument   

smp.go
```
package main

import "github.com/jacoloves/nou"

func main() {
	nou.DispUsage("")
}
```

result
```
$go run smp.go
Usage: smp [OPTION] ...
Try 'smp -help' for more information.
```

2. If you want to write your own Usage

smp.go
```
package main

import "github.com/jacoloves/nou"

func main() {
	usage_str := `Usage: smp [OPTION] ...
	-t: Output test text
	-s: Output sample text
	-h: Output help
This is sample Usage.`
	nou.DispUsage(usage_str)
}
```

result
```
$go run smp.go
Usage: smp [OPTION] ...
        -t: Output test text
        -s: Output sample text
        -h: Output help
This is sample Usage.
```

## Installation
```
go install github.com/jacoloves/nou
```

or

```
go get github.com/jacoloves/nou
```

## License
Distibuted under MIT License. See LICENSE.
