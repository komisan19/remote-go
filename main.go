package main

import (
    "github.com/moby/moby/client"
    "github.com/docker/docker/api/types"
    "context"
    "fmt"
)

func main() {
    ctx := context.Background()
    cl, _ := client.NewEnvClient()

    list, _ := cl.ImageList(ctx, types.ImageListOptions{All: true})


    fmt.Println(list)
}
