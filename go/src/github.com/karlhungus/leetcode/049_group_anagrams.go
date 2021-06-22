package main

func groupAnagrams(strs []string) [][]string {
	anagrams := map[[26]int][]string{}
	//letters are a-z only so we know we have 26 char
	//create an array of 26 ints, store the count of each int at it's position
	//use this as the hash for the string
	//any string that hashes to that value is an anagram
	for _, str := range strs {
		var hash [26]int
		for _, r := range str {
			hash[int(r)-'a']++
		}
		anagrams[hash] = append(anagrams[hash], str)
	}

	var res [][]string
	for _, v := range anagrams {
		res = append(res, v)
	}
	return res
}
