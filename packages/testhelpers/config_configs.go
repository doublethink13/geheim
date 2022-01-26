package testhelpers

func GetConfig1() string {
	return `---
secretkey: 'test1'
`
}

func GetConfig2() string {
	return `---
secretkey: 123456789
`
}

func GetConfig3() string {
	return `---
secretkey: ''
`
}

func GetConfig4() string {
	return `---
secretkey: 'test'
files: []
`
}

func GetConfig5() string {
	return `---
secretkey: 'test'
files:
  - testfile1
  - testfile2
`
}

func GetConfig6() string {
	return `---
secretkey: 'test'
files: 'thisiswrong'
`
}
