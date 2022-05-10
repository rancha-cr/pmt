package pics

import (
    "fmt"
    "image"
    "os"
    _ "image/jpeg"
    _ "image/png"
)

func FetchImages(folder string) ([]int, error) {
    if err := os.Open(filename)
    if err != nil {
        return "", err
    }
    defer f.Close()    // f.Close will close the file when finished.

    var result []byte
    buf := make([]byte, 100)
    for {
        n, err := f.Read(buf[0:])
        result = append(result, buf[0:n]...)
        if err != nil {
            break
        }
        return "", err  // f will be closed if we return here
    }
    return string(result), nil // f will be closed if we return here.
}

func GetImageDimension(imagePath string) (int, int) {
    file, err := os.Open(imagePath)
    if err != nil {
        log.Println(err)
    }

    image, _, err := image.DecodeConfig(file)
    if err != nil {
        fmt.Fprintf(os.Stderr, "%s:%v\n", imagePath, err)
    }
    return image.Width, image.Height
}
