# Solid State Drives

Up: [External Memory](external_memory)
Brother(s): [Hard Disk Drives](hard_disk_drives)
TARGET DECK
Leaves of the Tree of Ohara::Branch I::Semester I::1203 CompSys::Computer Systems


SSDs are a form of external storage that uses electronic circuitry built with semiconductors. These either used NAND logic or Flash memory. They are more expensive per byte compared to HDDs, but that price difference is slowly decreasing.

They have much faster access times in comparison to HDDs, and much less latency.
SSD read speed ranges from 500-7000MiB/s, with an Input/output Operation per Second rate of 90000.
The largest are currently around 18TiB.

### Key Benefits 

They are much more shock resistant, more power efficient and silent in comparison to HDDs.


#### Disadvantages

SSDs have a limit on the amount that they can be written to (Millions of times)

So, how is this mitigated?
	> **Wear levelling** - writes to different areas on the drive and not just one, to reduce the load on any one area
	> **Overprovision** - spare storage used by SSD controller, just in case

The more full an SSD is, the quicker it takes to wear out the remaining space, since less space reduces the effectiveness of wear levelling.

How do SSDs mitigate the write limit problem #flashcard 
**Wear levelling** - writing to different areas on the disk to reduce the load on any area
**Overprovision** - spare storage used by [Solid State Drives|SSD](solid_state_drives|ssd) controller, just in case
<!--ID: 1701369224364-->


SSDs must constantly update their *firmware* and drivers.

![Pasted image 20231129134422.png](pasted_image_20231129134422.png)

Data is written/read in blocks on SSDs.
The controller handles interfacing, wear levelling, addressing, error correction/detection, etc.

[PCIe as a SSD interface|PCIe can be used as a SSD interface](pcie_as_a_ssd_interface|pcie_can_be_used_as_a_ssd_interface)
































#### Why:
#### How:









