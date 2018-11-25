package main

var shapes = map[string]string{
	"heart": `:x::x::x::x::x::x::x::x::x::x::x:
:x::x::x:            :x:           :x::x::x:
:x::x:                                        :x::x:
:x::x::x:                             :x::x::x:
:x::x::x::x:                  :x::x::x::x:
:x::x::x::x::x:      :x::x::x::x::x:
:x::x::x::x::x::x::x::x::x::x::x:`,
	"like": `                                    :x::x:
                                    :x::x:
                              :x::x::x:
:x::x:           :x::x::x::x::x::x::x:
:x::x::x::x::x::x::x::x::x::x::x:
:x::x::x::x::x::x::x::x::x::x::x:
:x::x::x::x::x::x::x::x::x::x::x:
:x::x:           :x::x::x::x::x::x::x:`,
}

func getShapeKeys() []string {
	var keys []string
	for key := range shapes {
		keys = append(keys, key)
	}
	return keys
}

func getShape(shape string) string {
	return shapes[shape]
}
