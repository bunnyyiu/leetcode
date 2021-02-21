func subdomainVisits(cpdomains []string) []string {
    subdomains := make(map[string]int)
    for _, cpDomain := range cpdomains {
        cpDomainParts := strings.Split(cpDomain, " ")
        count, _ := strconv.Atoi(cpDomainParts[0])
        domain := cpDomainParts[1]
        
        parts := strings.Split(domain, ".")
        
        for i:= len(parts) - 1; i >= 0; i-- {
            newPart := parts[i : len(parts)]
            subdomain := strings.Join(newPart, ".")
            
            if val, ok := subdomains[subdomain]; ok {
                subdomains[subdomain] = val + count
            } else {
                subdomains[subdomain] = count
            }
        }
    }
    
    
    results := make([]string, len(subdomains))
    i := 0
    for k, v := range subdomains {
        results[i] = fmt.Sprintf("%d %s", v, k)
        i++
    }
    return results
}