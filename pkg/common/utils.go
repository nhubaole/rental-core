package common

func IfNullStr(requestVal, templateVal *string) string {
    if requestVal != nil && *requestVal != "" {
        return *requestVal
    }
    if templateVal != nil {
        return *templateVal
    }
    return ""
}


func IfNullFloat64(requestVal, templateVal *float64) float64 {
    if requestVal != nil {
        return *requestVal
    }
    if templateVal != nil {
        return *templateVal
    }
    return 0
}