# Maps 

Map is a data collection that store values, but instead of having a zero based index like array, we are actually going to provide our own key, our own index type.

- Maps are dynamically sized as same as slice 
- Maps work same as slice, when you copy the map and update the value in any of the map the result will reflect in both.
- if want to clone the map so that values not change after ateration use clone function ``maps.Clone``



```go

// declare A map

var m map[string]int // inside the bracket we put key type and after that value type
fmt.Println(m) // map[] (nil)

// Map literal
m = map[string]int{"foo":1, "bar":2}
fmt.Println(m) // map[foo:1 bar:2]

fmt.Println(m["foo"]) // lookup value in map 
m["bar"] = 99 // update value in map

delete(m, "foo") // remove entry from map
m["baz"] = 418 // add value to map

fmt.Println(m) // map[bar:99 baz:418]

// now the key foo is removeed what happen if we query the foo again?

fmt.Println(m["foo"]) // 0   ---> queries always return result even if there is no key associated and the default value of value data-type.

// if we don't want this behavior we can use comma ok syntax that go provide us 
v, ok := m["foo"] // comma okey syntax verifies presence

// here v variable store the value and ok is a variable which store if the key was present of not in the map as boolean. 

fmt.Println(v, ok) // 0, false

// copy the map
m := map[string]int{
    "foo":1,
    "bar":2,
    "baz":3
}

m2 := m 

// changing the values 
m["foo"], m2["bar"] = 99, 42
fmt.Println(m) // map[foo:99 bar:42 baz:3]
fmt.Println(m2) // map[foo:99 bar:42 baz:3]

// data is shared if mapp is copied
```


