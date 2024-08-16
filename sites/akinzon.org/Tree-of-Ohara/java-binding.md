# JavaFX Binding

Up: [JavaFX](JavaFX)

In JavaFX, binding is a way of synchronising the values of two components/variables.

The motivation for it is to have interface elements ([nodes](JavaFX-nodes))

keep up-to-date with underlying data.

JavaFX uses [Properties](JavaFX-Properties) to implement binding.

Binding can be one-way, where a property `B` updates when property `A` updates.

Binding can also be two-way, where property `B` can affect property `A` as well.
