package main

import "fmt"

//func copy(path string) {
//	rel, _ := filepath.Rel(from, path)
//	toPath := filepath.Join(to, rel)
//	bytes, _ := ioutil.ReadFile(path)
//	fmt.Printf("Writing %s to %s\n", path, toPath)
//	ioutil.WriteFile(toPath, bytes, os.ModePerm)
//}
//
//func walkFunc(path string, info os.FileInfo, err error) error {
//	if !strings.Contains(path, ".wdmc") && !strings.Contains(path, ".picasa.ini") && !info.IsDir() {
//		copy(path)
//	}
//
//	return nil
//}

func main() {
//	filepath.Walk(from, filepath.WalkFunc(walkFunc))

	fmt.Printf("Hi from app\n")
}

