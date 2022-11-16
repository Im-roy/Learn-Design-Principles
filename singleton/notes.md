# Singleton design pattern

for some component it makes sense to have only one instance present
- Database object
- logger object

Object construction is expensive.
- We only do it once
- We give every one the same object

We also need to make sure our system doesn't allow to create
additional copy of the object.
And the object instantiation should be lazy (which means once required).