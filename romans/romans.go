package romans

var romans string

func Romans(input int) string {
	for input > 0 {
		var num int
		if input >= 1000 {
			num = input / 1000
			romans += (MultiRom("M", num))
			input %= 1000
		} else if input >= 500 {
			if input >= 900 {
				romans += "CM"
				input -= 900
			} else {
				num -= 500
				if num >= 100 {
					if num >= 200 {
						if num >= 300 {
							romans += "DCCC"
							input -= 300
						} else {
							romans += "DCC"
							input -= 200
						}
					} else {
						romans += "DC"
						input -= 100
					}
				} else {
					romans += "D"
					input -= 500
				}

			}
		} else if input >= 100 {
			if input >= 200 {
				if input >= 300 {
					if input >= 400 {
						romans += "CD"
						input -= 400
					} else {
						romans += "CCC"
						input -= 300
					}
				} else {
					romans += "CC"
					input -= 200
				}
			} else {
				romans += "C"
				input -= 100
			}
		} else if input >= 50 {
			if input >= 90 {
				romans += "XC"
				input -= 90
			} else {
				if input >= 60 {
					if input >= 70 {
						if input >= 80 {
							romans += "LXXX"
							input -= 80
						} else {
							romans += "LXX"
							input -= 70
						}
					} else {
						romans += "LX"
						input -= 60
					}
				} else {
					romans += "L"
					input -= 50
				}
			}
		} else {
			if input >= 40 {
				romans += "XL"
				input -= 40
			} else {
				if input >= 10 {
					if input >= 20 {
						if input >= 30 {
							romans += "XXX"
							input -= 30
						} else {
							romans += "XX"
							input -= 20
						}
					} else {
						romans += "X"
						input -= 10
					}
				} else {
					if input >= 5 {
						if input == 9 {
							romans += "IX"
							input -= 9
						} else {
							romans += "V" + MultiRom("I", (input-5))
							input -= 5
						}
					} else {
						if input == 4 {
							romans += "IV"
							input -= 4
						} else {
							romans += MultiRom("I", (input))
							input = 0
						}
					}
				}
			}
		}
	}
	return romans
}

func MultiRom(s string, num int) (fin string) {
	for i := 1; i <= num; i++ {
		fin += s
	}
	return
}
