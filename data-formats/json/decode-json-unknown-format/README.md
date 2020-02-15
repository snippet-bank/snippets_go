Although Go strongly encourages you to be specific with types, you may decide to work with arbitrary JSON objects.

In these cases, you can use empty interfaces (`interface{}`) and type-switching to interact with the data. 

This example is mainly here for you to see that this is unpleasant to do. 

It is _much better_ to define specific structs that match the incoming data. This avoids / catches errors.
