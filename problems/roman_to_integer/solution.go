func romanToInt(s string) int {
    result := 0
    
    vAppear := false
    xAppear := false
    lAppear := false
    cAppear := false
    dAppear := false
    mAppear := false
    
    for i := len(s) - 1; i >= 0; i-- {
        c := s[i]
        switch c {
            case 'I':
                if vAppear || xAppear {
                    result -= 1
                } else {
                    result +=1 
                }
            case 'V':
                vAppear = true
                result += 5
            case 'X':
                xAppear = true
                if lAppear || cAppear {
                    result -= 10
                } else {
                    result +=10 
                }
            case 'L':
                lAppear = true
                result += 50
            case 'C':
                cAppear = true
                if dAppear || mAppear {
                    result -= 100
                } else {
                    result += 100
                }
            case 'D':
                dAppear = true
                result += 500
            case 'M':
                mAppear = true
                result += 1000
        }
    }
    return result
}