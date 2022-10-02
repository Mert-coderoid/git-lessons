# GO

## **INTERFACE**
```
	f = map[string]interface{}{
		"Name": "Wednesday",
		"Age":  6,
		"Parents": []interface{}{
			"Gomez",
			"Morticia",
		},
		"Children": 2,
		"bool":     true,
	}

	m := f.(map[string]interface{})

	for k, v := range m {
		switch vv := v.(type) {
		case string:
			fmt.Println(k, "is string", vv)
		case float64:
			fmt.Println(k, "is float64", vv)
		case []interface{}:
			fmt.Println(k, "is an array:")
			for i, u := range vv {
				fmt.Println(i, u)
			}
		// if its int
		case int:
			fmt.Println(k, "is int", vv)

		default:
			fmt.Println(k, "is of a type I don't know how to handle")
		}
	}

```

## **RANGE**
```
    // range
    for i, v := range []int{1, 2, 3, 4, 5} {
        fmt.Println(i, v)
    }

```
