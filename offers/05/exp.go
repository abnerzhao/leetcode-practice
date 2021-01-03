// 替换空格
// https://leetcode-cn.com/problems/ti-huan-kong-ge-lcof/


// replaceSpace 空格切分，"%20"组合
func replaceSpace(s string) string {
    strList := strings.Split(s, " ")
    return strings.Join(strList, "%20")
}

func replaceSpaceV2(s string) string {
	return strings.Replace(s, " ", "%20")
}

func replaceSpaceV3(s string) string{
	var str string
    for i := range s {
        if s[i] == ' '{
            str = str + "%20"
        } else {
            str = str + string(s[i])
        }
    }
    return str
}

func replaceSpaceV4(s string) string{
	var res strings.Builder
    for i:=range s{
    	if s[i]==' '{
			res.WriteString("%20")
		}else {
			res.WriteByte(s[i])
		}
	}
	return res.String()
}