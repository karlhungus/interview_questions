package main

isHappy(n int) bool {
    vals := make(map[int]bool)
    cur := n
    for {
        cur = sumSquares(cur)
        if cur == 1{
            return true
        }
        if _,ok := vals[cur];ok {
            return false
        } else {
            vals[cur] = true
        }
    } 
}

func sumSquares(n int) int {
    slice := strconv.Itoa(n)   
    sum := 0
    for _,x := range slice {
        b, _ := strconv.Atoi(string(x))
        sum = sum + (b*b)
        
    }
    return sum
}
