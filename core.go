package main
import (
    "strings"
    "sort"
)
// Analyze summarizes a stream of input without modifying user files.
func Analyze(input string) map[string]interface{} {

groups:=map[string][]string{"features":{},"fixes":{},"other":{}}
for _,line:=range strings.Split(input,"\n"){subject:=strings.TrimSpace(line);if subject==""{continue}
key:="other";if strings.HasPrefix(subject,"feat:"){key="features"}else if strings.HasPrefix(subject,"fix:"){key="fixes"}
groups[key]=append(groups[key],subject)}
keys:=[]string{};for key:=range groups{if len(groups[key])>0{keys=append(keys,key)}};sort.Strings(keys)
return map[string]interface{}{"groups":groups,"sections":keys}
}
