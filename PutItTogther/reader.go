package errortest 

import (
	"fmt"
	"os"
	"encoding/json"
)

func Readconfig(pathFile string , v interface{}) error{
	data , err := os.ReadFile(pathFile)
	if err != nil{
		return fmt.Errorf("reader file not working %s %w ",pathFile,err)
	}

	if err := json.Unmarshal(data,v); err != nil{
		return fmt.Errorf("this is not json data %s %w ",pathFile,err)
	}
	return nil
}