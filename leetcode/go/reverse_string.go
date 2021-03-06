func reverseString(s []byte)  {
    j := 0
    for i := len(s)-1; i >= len(s)/2; i-- {
        s[i], s[j] = s[j], s[i]
        j++
    }
}

func reverseString2(s []byte)  {
    for i:=0;i < len(s)/2; i++ {
        s[i], s[len(s)-i-1] = s[len(s)-i-1], s[i]
    }   
}