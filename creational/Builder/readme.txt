Builder design pattern – reusing an algorithm to create many implementations of an interface

 construct complex objects without directly instantiating their struct

 Wrapping up the Builder design pattern
The Builder design pattern helps us maintain an unpredictable number of products by using a common construction algorithm that is used by the director. The construction
process is always abstracted from the user of the product. At the same time, having a defined construction pattern helps when a newcomer to our
source code needs to add a new product to the pipeline. The BuildProcess interface
specifies what he must comply to be part of the possible builders.
However, try to avoid the Builder pattern when you are not completely sure that the
algorithm is going to be more or less stable because any small change in this interface will
affect all your builders and it could be awkward if you add a new method that some of your
builders need and others Builders do not.