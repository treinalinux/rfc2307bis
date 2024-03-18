package main

import (
  "flag"
  "fmt"
  "log"
  "os/exec"
  "strings"
)

func showAttributes(attribute string)  {
  attribute_valids := []string{"uid", "uidNumber", "gidNumber", "loginShell", "unixHomeDirectory"}

  for _, v := range attribute_valids {
    search_value := v + ": "
    attr := strings.HasPrefix(attribute, search_value)
    if attr {
      fmt.Printf("%v\n", attribute)
    }
  }
}

func searchAccountAttribute(name string) string {
  account := strings.Join([]string{"(&(objectClass=user)(cn=", name, "))"}, "")
  cmd, err := exec.Command("ldapsearch", "-Q", "-LLL", "-o", "ldif-wrap=no", account, "dn", "uid", "uidNumber", "gidNumber", "unixHomeDirectory","loginShell").Output()

  if err != nil {
      log.Printf("Erros ao consultar a chave: %v", err)
  }
  result := string(cmd)
  return result
}

func main()  {
  user := flag.String("user", "", "User Name in lowecase")
  flag.Parse()

  rldap := searchAccountAttribute(*user)
  attributes := strings.Split(rldap, "\n")
  for e := 0; e < len(attributes); e++ {
    showAttributes(attributes[e])
  }
}
