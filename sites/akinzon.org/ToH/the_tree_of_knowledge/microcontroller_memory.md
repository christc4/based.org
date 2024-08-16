# Microcontroller memory

Up: [Microcontrollers](microcontrollers)
Brother(s):
TARGET DECK
Leaves of the Tree of Ohara::Branch I::Semester I::1203 CompSys::Basic Concepts and History of the Computer

Microcontrollers do have [Registers|registers](registers|registers)

Microcontrollers use [SRAM](sram) which can be as low as 1kB

Microcontrollers also have Flash memory/EEP[ROM](rom), ranging from 16-64kB

Microcontrollers usually do not have cache

Microcontrollers also usually do not have a [Memory management by the OS|memory management](memory_management_by_the_os|memory_management) unit (MMU)

**Problems of no MMU** - cannot hide physical memory from processes like normal computers do - no capacity for [Paging](paging), [Segmentation](segmentation), etc.

Therefore, microcontrollers cannot run your typical [The Operating system|operating systems](the_operating_system|operating_systems)
(While a MMU can be emulated in software, that would mean having a really large overhead)



































#### Why:
#### How:









