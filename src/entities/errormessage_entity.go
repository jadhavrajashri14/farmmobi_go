package entities

import (
  "fmt"
)

type Error struct {
  Message			  string 			`json:"message"`
}


func (error Error) ToString() string {

  return fmt.Sprintf(`Error:%s\n`, error.Message)
}
