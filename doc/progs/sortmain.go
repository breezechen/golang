// Copyright 2009 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package main

import Sort "sort"

func ints() {
	data := []int{74, 59, 238, -784, 9845, 959, 905, 0, 0, 42, 7586, -5467984, 7586};
	a := Sort.IntArray{&data};
	Sort.Sort(&a);
	if !Sort.IsSorted(&a) {
		panic()
	}
}

func strings() {
	data := []string{"monday", "tuesday", "wednesday", "thursday", "friday", "saturday", "sunday"};
	a := Sort.StringArray{&data};
	Sort.Sort(&a);
	if !Sort.IsSorted(&a) {
		panic()
	}
}

type Day struct {
	num        int;
	short_name string;
	long_name  string;
}

type DayArray struct {
	data *[]*Day;
}

func (p *DayArray) len() int            { return len(p.data); }
func (p *DayArray) less(i, j int) bool  { return p.data[i].num < p.data[j].num; }
func (p *DayArray) swap(i, j int)       { p.data[i], p.data[j] = p.data[j], p.data[i]; }

func days() {
	Sunday :=    Day{ 0, "SUN", "Sunday" };
	Monday :=    Day{ 1, "MON", "Monday" };
	Tuesday :=   Day{ 2, "TUE", "Tuesday" };
	Wednesday := Day{ 3, "WED", "Wednesday" };
	Thursday :=  Day{ 4, "THU", "Thursday" };
	Friday :=    Day{ 5, "FRI", "Friday" };
	Saturday :=  Day{ 6, "SAT", "Saturday" };
	data := []*Day{&Tuesday, &Thursday, &Sunday, &Monday, &Friday};
	a := DayArray{&data};
	Sort.Sort(&a);
	if !Sort.IsSorted(&a) {
		panic()
	}
	for i := 0; i < len(data); i++ {
		print(data[i].long_name, " ")
	}
	print("\n")
}


func main() {
	ints();
	strings();
	days();
}
