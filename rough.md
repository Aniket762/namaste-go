## Notes for futureself!

1.  If-else  
    '''
    loginUser :=23
    var result string
    // { } positing are imp!
    if loginUser<10 {
    result = "regular"
    } else if loginUser>10 {
    result = "watchout"
    }else{
    result = "10"
    }

        // assigning and then checking
        if num:=3;num<10{
            result="hehe"
        }else{
            result="huihui"
        }

    '''

2.  Switch case
    '''
    switch diceNumber{
    case 1:
    fmt.Println("....")
    case 2:
    ...
    default:
    ...
    }

'''

3. If we write defer before any function it will be
   excecuted at the very last!
   func main()
   {
   defer fmt.Println("Hi")
   fmt.Println("Hello")
   }

o/p for this will be:
Hello
Hi

ie due to use of defer it will be excecuted at the last

now understanding LIFO

func main()
{
defer fmt.Println("World")
defer fmt.Println("One")
defer fmt.Println("Two")
fmt.Println("Hello")
}

o/p:
Hello
Two
One
World

ie when more than one defer is mentioned
it follows LIFO ie the last defer gives output first

4. Go routines are threads running via concurrency controlled by go runtimes
