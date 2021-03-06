package erratum

func opener(o func() (Resource, error)) (Resource, error) {
	resource, err := o()
	if err != nil {
		if _, ok := err.(TransientError); ok {
			return opener(o)
		}
		return nil, err
	}
	return resource, nil
}

func Use(o func() (Resource, error), input string) (err error) {
	resource, err := opener(o)
	if err != nil {
		return err
	}
	defer resource.Close()

	defer func() {
		if r := recover(); r != nil {
			if val, ok := r.(FrobError); ok {
				resource.Defrob(val.defrobTag)
			}
			err = r.(error)
		}
	}()

	resource.Frob(input)

	return nil
}
