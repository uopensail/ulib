# source
This is a library that reads the item features from a file and then generates a Source. Mainly using `sample.ImmutableFeatures` to store features data. At the same time, `Collection` and `Condition` are supported here, as follows:

## Collection
Collection is a collection of a series of item IDs, calculated according to a certain condition (only using item features).

## Condition
A condition is a condition that is calculated according to a condition (the condition not only uses item features, but also relies on other external features).