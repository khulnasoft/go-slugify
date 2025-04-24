// Package slugify implements Make Pretty Slug
//
//	slugifier := (&slugify.Slugifier{}).ToLower(false).InvalidChar("-").WordSeparator("-")
//	s := "北京kožušček,abc"
//	fmt.Println(slugifier.Slugify(s))
//	// Output: bei-jing-kozuscek-abc
package slugify
