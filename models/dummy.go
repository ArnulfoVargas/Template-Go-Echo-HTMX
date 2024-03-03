package models

type dummy struct{
  Name string
}

func NewDummy() *dummy {
  return &dummy{
    Name: "Test",
  }
}
