package main

import (
    "fmt"
    "context"
    "time"
)

func main() {

    // 空の Context を生成　TODO()でも可
    ctx := context.Background()
    go parentRoutine(ctx)

    time.Sleep(20 * time.Second)
    fmt.Println("main func finish")

}

func parentRoutine(ctx context.Context) {

    // cancel 可能な状態でコンテキストを生成
    subCtx, cancel := context.WithCancel(ctx)
    subsubCtx, cancel2 := context.WithCancel(subCtx)

    go childRoutine(subCtx, "sub context")
    go childRoutine(subsubCtx, "sub sub context")

    time.Sleep(5 * time.Second)

    // subsubCtx をキャンセルする
    cancel2()

    time.Sleep(5 * time.Second)

    // subCtx をキャンセルする
    cancel()

    // timeout を設定してコンテキストを生成する
    tCtx, _ := context.WithTimeout(ctx, time.Second * 5)
    go childRoutine(tCtx, "context with timeout")

}

func childRoutine(ctx context.Context, prefix string) {
    for i := 0; ; i++{

        select {

        // cancel 時の挙動
        case <-ctx.Done():
            fmt.Printf("routine %s cancel.\n", prefix)
            return

         case <-time.After(1 * time.Second):
             fmt.Printf("routine %s has value %d \n", prefix, i)

        }
    }
}