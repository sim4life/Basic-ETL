package main

type NameSorter Hotels

func (ns NameSorter) Len() int {
	return len(ns.Hotels)
}
func (ns NameSorter) Less(i, j int) bool {
	return ns.Hotels[i].Name < ns.Hotels[j].Name
}
func (ns NameSorter) Swap(i, j int) {
	ns.Hotels[i], ns.Hotels[j] = ns.Hotels[j], ns.Hotels[i]
}
