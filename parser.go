package agilis

func parse(dataStr string) (data Data) {
	i := -1
	data = make(Data)
	var tempName, tempValue string
	var b byte

	f := func(str *string) bool {
		if str != nil {
			*str += string(b)
		}
		i++
		if i >= len(dataStr) {
			return false
		}
		b = dataStr[i]
		return true
	}

	for i < len(dataStr) {
		for b != ':' {
			ok := f(&tempName)
			if !ok {
				break
			}
		}
		f(nil)
		for b != '=' {
			ok := f(nil)
			if !ok {
				break
			}
		}
		f(nil)
		for b != ',' {
			ok := f(&tempValue)
			if !ok {
				break
			}
		}
		f(nil)
		f(nil)
		tempValue += "}"
		data[tempName] = tempValue
		tempName = ""
		tempValue = ""
	}

	return
}
