package main

import (
  "fmt"
  "io/ioutil"
  "os"
  "os/exec"
  "os/user"
  "strconv"
  "strings"
  "bufio"

  "gopkg.in/yaml.v2"
)

func execute(command string) string {
  out, err := exec.Command("bash", "-c", command).Output()

  if err != nil {
    fmt.Println("%s", err)
  }

  output := string(out[:])
  return output
}


func main() {
  if len(os.Args) < 2 {
    fmt.Println("ERROR: A file need to be specified to build")
    os.Exit(1)
  }
  user, err := user.Current()
  if err != nil {
    panic(err)
  }
  configPath := "/home/" + user.Username + "/.buildit.yaml"

  file := os.Args[1]

  args := " "

  for i := 2; i < len(os.Args); i++{
    args += os.Args[i] + " "
  }

  fileSplited := strings.Split(file, ".")
  if len(fileSplited) < 2 {
    fmt.Println("ERROR: The file need to have a name, and a extension")
    os.Exit(1)
  }

  source, err := ioutil.ReadFile(configPath)
  if err != nil {
    panic(err)
  }
  var builditYAML map[string][]interface{}
  err = yaml.Unmarshal(source, &builditYAML)
  if err != nil {
    panic(err)
  }
  for lang := range builditYAML {
    if lang == fileSplited[1]{
      if len(builditYAML[lang]) == 1{
        command := fmt.Sprintf("%v", builditYAML[lang][0])
        fmt.Println(command + " " + file)
        output := execute(command + " " + file + args)
        fmt.Println(output)
        os.Exit(0)
      }else{
        fmt.Println("What command do you want to use? \n")
        for i := 0; i < len(builditYAML[lang]); i++{
          fmt.Println("[" + strconv.Itoa(i) + "]", builditYAML[lang][i])
        }
        reader := bufio.NewReader(os.Stdin)
        fmt.Print("Enter text: ")
        commandSelected, _ :=  reader.ReadString('\n')
        commandSelectedInt, _ := strconv.Atoi(commandSelected)
        command := fmt.Sprintf("%v", builditYAML[lang][commandSelectedInt])

        fmt.Println(command + " " + file + args)
        output := execute(command + " " + file + args)
        fmt.Println(output)
        os.Exit(0)
      }
    }
  }
  fmt.Println("ERROR: This file doesn't is defined in ~/.buildit.yaml")

}
