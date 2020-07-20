/* example copied from https://www.linode.com/docs/development/go/go-context/ to demostrate below points

This time we create a context using context.TODO() instead of context.Background(). Although both functions return a non-nil, empty Context, their purposes differ. You should never pass a nil context –– use the context.TODO() function to create a suitable context. Use the context.TODO() function when you are not sure about the Context that you want to use.

The context.TODO() function signifies that we intend to use an operation context, without being sure about it yet. The good thing is that TODO() is recognized by static analysis tools, which allows them to determine whether a context.Context variable is propagated correctly in a program or not.

The context.WithValue() function that is used in main() offers a way to associate a value with a Context`.

The searchKey() function retrieves a value from a Context variable and checks whether that value exists or not.
*/
//another good example is at https://golangbyexample.com/using-context-in-golang-complete-guide/
package main

import (
	"context"
	"fmt"
)

type aKey string

func searchKey(ctx context.Context, k aKey) {
	v := ctx.Value(k)
	if v != nil {
		fmt.Println("found value:", v)
		return
	} else {
		fmt.Println("key not found:", k)
	}
}

func main() {

	myKey := aKey("myValue")
	ctx := context.WithValue(context.Background(), myKey, "myValue")
	searchKey(ctx, myKey)

	emptyCtx := context.TODO()
	searchKey(emptyCtx, aKey("notThere"))
}
